// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package types

import (
	"encoding/json"
	"errors"
	"math/big"

	"github.com/MetisProtocol/l2geth/common"
	"github.com/MetisProtocol/l2geth/common/hexutil"
)

var _ = (*txdataMarshaling)(nil)

// TransactionMarshalJSON marshals as JSON.
func (t txdata) TransactionMarshalJSON() ([]byte, error) {
	type txdata struct {
		AccountNonce hexutil.Uint64  `json:"nonce"    gencodec:"required"`
		Price        *hexutil.Big    `json:"gasPrice" gencodec:"required"`
		GasLimit     hexutil.Uint64  `json:"gas"      gencodec:"required"`
		Recipient    *common.Address `json:"to"       rlp:"nil"`
		Amount       *hexutil.Big    `json:"value"    gencodec:"required"`
		Payload      hexutil.Bytes   `json:"input"    gencodec:"required"`
		V            *hexutil.Big    `json:"v" gencodec:"required"`
		R            *hexutil.Big    `json:"r" gencodec:"required"`
		S            *hexutil.Big    `json:"s" gencodec:"required"`
		Hash         *common.Hash    `json:"hash" rlp:"-"`
		// NOTE 20210724
		// L1Info
		// L1BlockNumber     *hexutil.Big          `json:"l1BlockNumber"`
		// L1Timestamp       hexutil.Uint64            `json:"l1Timestamp"`
		// L1MessageSender   *common.Address   `json:"L1MessageSender" gencodec:"required"`
		// SignatureHashType SignatureHashType `json:"signatureHashType" gencodec:"required"`
		// QueueOrigin       *hexutil.Big          `json:"queueOrigin" gencodec:"required"`
		// // The canonical transaction chain index
		// Index hexutil.Uint64 `json:"index" gencodec:"required"`
		// // The queue index, nil for queue origin sequencer transactions
		// QueueIndex hexutil.Uint64 `json:"queueIndex" gencodec:"required"`
	}
	var enc txdata
	enc.AccountNonce = hexutil.Uint64(t.AccountNonce)
	enc.Price = (*hexutil.Big)(t.Price)
	enc.GasLimit = hexutil.Uint64(t.GasLimit)
	enc.Recipient = t.Recipient
	enc.Amount = (*hexutil.Big)(t.Amount)
	enc.Payload = t.Payload
	enc.V = (*hexutil.Big)(t.V)
	enc.R = (*hexutil.Big)(t.R)
	enc.S = (*hexutil.Big)(t.S)
	enc.Hash = t.Hash
	// NOTE 20210724
	// enc.L1BlockNumber = (*hexutil.Big)(t.L1BlockNumber)
	// enc.L1Timestamp = hexutil.Uint64(t.L1Timestamp)
	// enc.L1MessageSender = t.L1MessageSender
	// enc.SignatureHashType = t.SignatureHashType
	// enc.QueueOrigin = (*hexutil.Big)(t.QueueOrigin)
	// enc.Index = hexutil.Uint64(*t.Index)
	// enc.QueueIndex = hexutil.Uint64(*t.QueueIndex)
	return json.Marshal(&enc)
}

// TransactionUnmarshalJSON unmarshals from JSON.
func (t *txdata) TransactionUnmarshalJSON(input []byte) error {
	type txdata struct {
		AccountNonce *hexutil.Uint64 `json:"nonce"    gencodec:"required"`
		Price        *hexutil.Big    `json:"gasPrice" gencodec:"required"`
		GasLimit     *hexutil.Uint64 `json:"gas"      gencodec:"required"`
		Recipient    *common.Address `json:"to"       rlp:"nil"`
		Amount       *hexutil.Big    `json:"value"    gencodec:"required"`
		Payload      *hexutil.Bytes  `json:"input"    gencodec:"required"`
		V            *hexutil.Big    `json:"v" gencodec:"required"`
		R            *hexutil.Big    `json:"r" gencodec:"required"`
		S            *hexutil.Big    `json:"s" gencodec:"required"`
		Hash         *common.Hash    `json:"hash" rlp:"-"`
		// NOTE 20210724
		// L1Info
		// L1BlockNumber     *hexutil.Big          `json:"l1BlockNumber"`
		// L1Timestamp       *hexutil.Uint64            `json:"l1Timestamp"`
		// L1MessageSender   *common.Address   `json:"L1MessageSender" gencodec:"required"`
		// SignatureHashType SignatureHashType `json:"signatureHashType" gencodec:"required"`
		// QueueOrigin       *hexutil.Big          `json:"queueOrigin" gencodec:"required"`
		// // The canonical transaction chain index
		// Index *hexutil.Uint64 `json:"index" gencodec:"required"`
		// // The queue index, nil for queue origin sequencer transactions
		// QueueIndex *hexutil.Uint64 `json:"queueIndex" gencodec:"required"`
	}
	var dec txdata
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.AccountNonce == nil {
		return errors.New("missing required field 'nonce' for txdata")
	}
	t.AccountNonce = uint64(*dec.AccountNonce)
	if dec.Price == nil {
		return errors.New("missing required field 'gasPrice' for txdata")
	}
	t.Price = (*big.Int)(dec.Price)
	if dec.GasLimit == nil {
		return errors.New("missing required field 'gas' for txdata")
	}
	t.GasLimit = uint64(*dec.GasLimit)
	if dec.Recipient != nil {
		t.Recipient = dec.Recipient
	}
	if dec.Amount == nil {
		return errors.New("missing required field 'value' for txdata")
	}
	t.Amount = (*big.Int)(dec.Amount)
	if dec.Payload == nil {
		return errors.New("missing required field 'input' for txdata")
	}
	t.Payload = *dec.Payload
	if dec.V == nil {
		return errors.New("missing required field 'v' for txdata")
	}
	t.V = (*big.Int)(dec.V)
	if dec.R == nil {
		return errors.New("missing required field 'r' for txdata")
	}
	t.R = (*big.Int)(dec.R)
	if dec.S == nil {
		return errors.New("missing required field 's' for txdata")
	}
	t.S = (*big.Int)(dec.S)
	if dec.Hash != nil {
		t.Hash = dec.Hash
	}
	// NOTE 20210724
	// if dec.L1BlockNumber != nil {
	// 	t.L1BlockNumber = (*big.Int)(dec.L1BlockNumber)
	// }
	// if dec.L1Timestamp != nil {
	// 	t.L1Timestamp = uint64(*dec.L1Timestamp)
	// }
	// if dec.L1MessageSender != nil {
	// 	t.L1MessageSender = dec.L1MessageSender
	// }	
	// t.SignatureHashType = dec.SignatureHashType
	// if dec.QueueOrigin != nil {
	// 	t.QueueOrigin = (*big.Int)(dec.QueueOrigin)
	// }
	// if dec.Index != nil {
	// 	index1 := uint64(*dec.Index)
	// 	t.Index = &index1
	// }
	// if dec.QueueIndex != nil {
	// 	queueIndex1 := uint64(*dec.QueueIndex)
	// 	t.QueueIndex = &queueIndex1
	// }
	return nil
}
