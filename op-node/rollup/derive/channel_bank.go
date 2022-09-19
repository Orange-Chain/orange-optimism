package derive

import (
	"context"
	"io"

	"github.com/ethereum-optimism/optimism/op-node/eth"
	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
)

// ChannelBank is a stateful stage that does the following:
// 1. Unmarshalls frames from L1 transaction data
// 2. Applies those frames to a channel
// 3. Attempts to read from the channel when it is ready
// 4. Prunes channels (not frames) when the channel bank is too large.
//
// Note: we prune before we ingest data.
// As we switch between ingesting data & reading, the prune step occurs at an odd point
// Specifically, the channel bank is not allowed to become too large between successive calls
// to `IngestData`. This means that we can do an ingest and then do a read while becoming too large.

// ChannelBank buffers channel frames, and emits full channel data
type ChannelBank struct {
	log log.Logger
	cfg *rollup.Config

	channels     map[ChannelID]*Channel // channels by ID
	channelQueue []ChannelID            // channels in FIFO order

	prev    *L1Retrieval
	fetcher L1Fetcher
}

var _ ResetableStage = (*ChannelBank)(nil)

// NewChannelBank creates a ChannelBank, which should be Reset(origin) before use.
func NewChannelBank(log log.Logger, cfg *rollup.Config, prev *L1Retrieval, fetcher L1Fetcher) *ChannelBank {
	return &ChannelBank{
		log:          log,
		cfg:          cfg,
		channels:     make(map[ChannelID]*Channel),
		channelQueue: make([]ChannelID, 0, 10),
		prev:         prev,
		fetcher:      fetcher,
	}
}

func (ib *ChannelBank) Origin() eth.L1BlockRef {
	return ib.prev.Origin()
}

func (ib *ChannelBank) prune() {
	// check total size
	totalSize := uint64(0)
	for _, ch := range ib.channels {
		totalSize += ch.size
	}
	// prune until it is reasonable again. The high-priority channel failed to be read, so we start pruning there.
	for totalSize > MaxChannelBankSize {
		id := ib.channelQueue[0]
		ch := ib.channels[id]
		ib.channelQueue = ib.channelQueue[1:]
		delete(ib.channels, id)
		totalSize -= ch.size
	}
}

// IngestData adds new L1 data to the channel bank.
// Read() should be called repeatedly first, until everything has been read, before adding new data.\
func (ib *ChannelBank) IngestData(data []byte) {
	origin := ib.Origin()
	ib.log.Debug("channel bank got new data", "origin", origin, "data_len", len(data))

	// TODO: Why is the prune here?
	ib.prune()

	frames, err := ParseFrames(data)
	if err != nil {
		ib.log.Warn("malformed frame", "err", err)
		return
	}

	// Process each frame
	for _, f := range frames {
		currentCh, ok := ib.channels[f.ID]
		if !ok {
			// create new channel if it doesn't exist yet
			currentCh = NewChannel(f.ID, origin)
			ib.channels[f.ID] = currentCh
			ib.channelQueue = append(ib.channelQueue, f.ID)
		}

		// check if the channel is not timed out
		if currentCh.OpenBlockNumber()+ib.cfg.ChannelTimeout < origin.Number {
			ib.log.Warn("channel is timed out, ignore frame", "channel", f.ID, "frame", f.FrameNumber)
			continue
		}

		ib.log.Trace("ingesting frame", "channel", f.ID, "frame_number", f.FrameNumber, "length", len(f.Data))
		if err := currentCh.AddFrame(f, origin); err != nil {
			ib.log.Warn("failed to ingest frame into channel", "channel", f.ID, "frame_number", f.FrameNumber, "err", err)
			continue
		}
	}
}

// Read the raw data of the first channel, if it's timed-out or closed.
// Read returns io.EOF if there is nothing new to read.
func (ib *ChannelBank) Read() (data []byte, err error) {
	if len(ib.channelQueue) == 0 {
		return nil, io.EOF
	}
	first := ib.channelQueue[0]
	ch := ib.channels[first]
	timedOut := ch.OpenBlockNumber()+ib.cfg.ChannelTimeout < ib.Origin().Number
	if timedOut {
		ib.log.Debug("channel timed out", "channel", first, "frames", len(ch.inputs))
		delete(ib.channels, first)
		ib.channelQueue = ib.channelQueue[1:]
		return nil, io.EOF
	}
	if !ch.IsReady() {
		return nil, io.EOF
	}

	delete(ib.channels, first)
	ib.channelQueue = ib.channelQueue[1:]
	r := ch.Reader()
	// Suprress error here. io.ReadAll does return nil instead of io.EOF though.
	data, _ = io.ReadAll(r)
	return data, nil
}

// NextData pulls the next piece of data from the channel bank.
// Note that it attempts to pull data out of the channel bank prior to
// loading data in (unlike most other stages). This is to ensure maintain
// consistency around channel bank pruning which depends upon the order
// of operations.
func (ib *ChannelBank) NextData(ctx context.Context) ([]byte, error) {

	// Do the read from the channel bank first
	data, err := ib.Read()
	if err == io.EOF {
		// continue - We will attempt to load data into the channel bank
	} else if err != nil {
		return nil, err
	} else {
		return data, nil
	}

	// Then load data into the channel bank
	if data, err := ib.prev.NextData(ctx); err == io.EOF {
		return nil, io.EOF
	} else if err != nil {
		return nil, err
	} else {
		ib.IngestData(data)
		return nil, NotEnoughData
	}
}

func (ib *ChannelBank) Reset(ctx context.Context, base eth.L1BlockRef) error {
	ib.channels = make(map[ChannelID]*Channel)
	ib.channelQueue = make([]ChannelID, 0, 10)
	return io.EOF
}

type L1BlockRefByHashFetcher interface {
	L1BlockRefByHash(context.Context, common.Hash) (eth.L1BlockRef, error)
}
