package op_e2e

import (
	"context"
	"fmt"
	"math/rand"
	"net"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/stretchr/testify/require"

	bss "github.com/ethereum-optimism/optimism/op-batcher/batcher"
	"github.com/ethereum-optimism/optimism/op-batcher/compressor"
	batcherFlags "github.com/ethereum-optimism/optimism/op-batcher/flags"
	con "github.com/ethereum-optimism/optimism/op-conductor/conductor"
	conrpc "github.com/ethereum-optimism/optimism/op-conductor/rpc"
	rollupNode "github.com/ethereum-optimism/optimism/op-node/node"
	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum-optimism/optimism/op-node/rollup/derive"
	"github.com/ethereum-optimism/optimism/op-node/rollup/driver"
	"github.com/ethereum-optimism/optimism/op-node/rollup/sync"
	oplog "github.com/ethereum-optimism/optimism/op-service/log"
	oprpc "github.com/ethereum-optimism/optimism/op-service/rpc"
	"github.com/ethereum-optimism/optimism/op-service/sources"
	"github.com/ethereum-optimism/optimism/op-service/testlog"
)

const (
	Sequencer1Name = "sequencer1"
	Sequencer2Name = "sequencer2"
	Sequencer3Name = "sequencer3"
	VerifierName   = "verifier"

	localhost = "127.0.0.1"
)

type conductor struct {
	service       *con.OpConductor
	client        conrpc.API
	consensusPort int
}

func (c *conductor) ConsensusEndpoint() string {
	return fmt.Sprintf("%s:%d", localhost, c.consensusPort)
}

func setupSequencerFailoverTest(t *testing.T) (*System, map[string]*conductor) {
	InitParallel(t)
	ctx := context.Background()

	conductorRpcPorts := map[string]int{
		Sequencer1Name: findAvailablePort(t),
		Sequencer2Name: findAvailablePort(t),
		Sequencer3Name: findAvailablePort(t),
	}

	// 3 stopped sequencers, 1 verifier
	cfg := sequencerFailoverSystemConfig(t, conductorRpcPorts)
	sys, err := cfg.Start(t)
	require.NoError(t, err)

	// 3 conductors that connects to 1 sequencer each.
	conductors := make(map[string]*conductor)

	// initialize all conductors in paused mode
	conductorCfgs := []struct {
		name      string
		port      int
		bootstrap bool
	}{
		{Sequencer1Name, conductorRpcPorts[Sequencer1Name], true}, // one in bootstrap mode so that we can form a cluster.
		{Sequencer2Name, conductorRpcPorts[Sequencer2Name], false},
		{Sequencer3Name, conductorRpcPorts[Sequencer3Name], false},
	}
	for _, cfg := range conductorCfgs {
		cfg := cfg
		nodePRC := sys.RollupNodes[cfg.name].HTTPEndpoint()
		engineRPC := sys.EthInstances[cfg.name].HTTPEndpoint()
		conductors[cfg.name] = setupConductor(t, cfg.name, t.TempDir(), nodePRC, engineRPC, cfg.port, cfg.bootstrap, *sys.RollupConfig)
	}

	// form a cluster
	c1 := conductors[Sequencer1Name]
	c2 := conductors[Sequencer2Name]
	c3 := conductors[Sequencer3Name]
	require.NoError(t, waitForLeadershipChange(t, c1, true))
	require.NoError(t, c1.client.AddServerAsVoter(ctx, Sequencer2Name, c2.ConsensusEndpoint()))
	require.NoError(t, c1.client.AddServerAsVoter(ctx, Sequencer3Name, c3.ConsensusEndpoint()))
	require.True(t, leader(t, ctx, c1))
	require.False(t, leader(t, ctx, c2))
	require.False(t, leader(t, ctx, c3))

	// start sequencing on leader
	lid, _ := findLeader(t, conductors)
	unsafeHead, err := sys.Clients[lid].BlockByNumber(ctx, nil)
	require.NoError(t, err)
	require.NoError(t, sys.RollupClient(lid).StartSequencer(ctx, unsafeHead.Hash()))

	// 1 batcher that listens to all 3 sequencers, in started mode.
	setupBatcher(t, sys)

	require.Eventually(t, func() bool {
		return healthy(t, ctx, c1) &&
			healthy(t, ctx, c2) &&
			healthy(t, ctx, c3)
	}, 30*time.Second, 500*time.Millisecond, "Expected sequencer 1 to become healthy")

	// unpause all conductors
	require.NoError(t, c1.client.Resume(ctx))
	require.NoError(t, c2.client.Resume(ctx))
	require.NoError(t, c3.client.Resume(ctx))

	// final check, make sure everything is in the right place
	require.True(t, conductorActive(t, ctx, c1))
	require.True(t, conductorActive(t, ctx, c2))
	require.True(t, conductorActive(t, ctx, c3))

	require.True(t, sequencerActive(t, ctx, sys.RollupClient(Sequencer1Name)))
	require.False(t, sequencerActive(t, ctx, sys.RollupClient(Sequencer2Name)))
	require.False(t, sequencerActive(t, ctx, sys.RollupClient(Sequencer3Name)))

	require.True(t, healthy(t, ctx, c1))
	require.True(t, healthy(t, ctx, c2))
	require.True(t, healthy(t, ctx, c3))

	return sys, conductors
}

func setupConductor(
	t *testing.T,
	serverID, dir, nodePRC, engineRPC string,
	rpcPort int,
	bootstrap bool,
	rollupCfg rollup.Config,
) *conductor {
	// it's unfortunate that it is not possible to pass 0 as consensus port and get back the actual assigned port from raft implementation.
	// So we find an available port and pass it in to avoid test flakiness (avoid port already in use error).
	consensusPort := findAvailablePort(t)
	cfg := con.Config{
		ConsensusAddr:  localhost,
		ConsensusPort:  consensusPort,
		RaftServerID:   serverID,
		RaftStorageDir: dir,
		RaftBootstrap:  bootstrap,
		NodeRPC:        nodePRC,
		ExecutionRPC:   engineRPC,
		Paused:         true,
		HealthCheck: con.HealthCheckConfig{
			Interval:       1, // per test setup, l2 block time is 1s.
			UnsafeInterval: 3, // 1s block time + 2s buffer
			SafeInterval:   4,
			MinPeerCount:   2, // per test setup, each sequencer has 2 peers
		},
		RollupCfg: rollupCfg,
		LogConfig: oplog.CLIConfig{
			Level: log.LvlInfo,
			Color: false,
		},
		RPC: oprpc.CLIConfig{
			ListenAddr: localhost,
			ListenPort: rpcPort,
		},
	}

	ctx := context.Background()
	service, err := con.New(ctx, &cfg, testlog.Logger(t, log.LvlInfo), "0.0.1")
	require.NoError(t, err)
	err = service.Start(ctx)
	require.NoError(t, err)

	rawClient, err := rpc.DialContext(ctx, service.HTTPEndpoint())
	require.NoError(t, err)
	client := conrpc.NewAPIClient(rawClient)

	return &conductor{
		service:       service,
		client:        client,
		consensusPort: consensusPort,
	}
}

func setupBatcher(t *testing.T, sys *System) {
	batcherMaxL1TxSizeBytes := sys.Cfg.BatcherMaxL1TxSizeBytes
	if batcherMaxL1TxSizeBytes == 0 {
		batcherMaxL1TxSizeBytes = 240_000
	}

	// enable active sequencer follow mode.
	l2EthRpc := strings.Join([]string{
		sys.EthInstances[Sequencer1Name].WSEndpoint(),
		sys.EthInstances[Sequencer2Name].WSEndpoint(),
		sys.EthInstances[Sequencer3Name].WSEndpoint(),
	}, ",")
	rollupRpc := strings.Join([]string{
		sys.RollupNodes[Sequencer1Name].HTTPEndpoint(),
		sys.RollupNodes[Sequencer2Name].HTTPEndpoint(),
		sys.RollupNodes[Sequencer3Name].HTTPEndpoint(),
	}, ",")
	batcherCLIConfig := &bss.CLIConfig{
		L1EthRpc:               sys.EthInstances["l1"].WSEndpoint(),
		L2EthRpc:               l2EthRpc,
		RollupRpc:              rollupRpc,
		MaxPendingTransactions: 0,
		MaxChannelDuration:     1,
		MaxL1TxSize:            batcherMaxL1TxSizeBytes,
		CompressorConfig: compressor.CLIConfig{
			TargetL1TxSizeBytes: sys.Cfg.BatcherTargetL1TxSizeBytes,
			TargetNumFrames:     1,
			ApproxComprRatio:    0.4,
		},
		SubSafetyMargin: 0,
		PollInterval:    50 * time.Millisecond,
		TxMgrConfig:     newTxMgrConfig(sys.EthInstances["l1"].WSEndpoint(), sys.Cfg.Secrets.Batcher),
		LogConfig: oplog.CLIConfig{
			Level:  log.LvlInfo,
			Format: oplog.FormatText,
		},
		Stopped:              false,
		BatchType:            derive.SingularBatchType,
		DataAvailabilityType: batcherFlags.CalldataType,
	}

	batcher, err := bss.BatcherServiceFromCLIConfig(context.Background(), "0.0.1", batcherCLIConfig, sys.Cfg.Loggers["batcher"])
	require.NoError(t, err)
	err = batcher.Start(context.Background())
	require.NoError(t, err)
	sys.BatchSubmitter = batcher
}

func sequencerFailoverSystemConfig(t *testing.T, ports map[string]int) SystemConfig {
	cfg := DefaultSystemConfig(t)
	delete(cfg.Nodes, "sequencer")
	cfg.Nodes[Sequencer1Name] = sequencerCfg(ports[Sequencer1Name])
	cfg.Nodes[Sequencer2Name] = sequencerCfg(ports[Sequencer1Name])
	cfg.Nodes[Sequencer3Name] = sequencerCfg(ports[Sequencer1Name])

	delete(cfg.Loggers, "sequencer")
	cfg.Loggers[Sequencer1Name] = testlog.Logger(t, log.LvlInfo).New("role", Sequencer1Name)
	cfg.Loggers[Sequencer2Name] = testlog.Logger(t, log.LvlInfo).New("role", Sequencer2Name)
	cfg.Loggers[Sequencer3Name] = testlog.Logger(t, log.LvlInfo).New("role", Sequencer3Name)

	cfg.P2PTopology = map[string][]string{
		Sequencer1Name: {Sequencer2Name, Sequencer3Name},
		Sequencer2Name: {Sequencer3Name, VerifierName},
		Sequencer3Name: {VerifierName, Sequencer1Name},
		VerifierName:   {Sequencer1Name, Sequencer2Name},
	}

	return cfg
}

func sequencerCfg(port int) *rollupNode.Config {
	return &rollupNode.Config{
		Driver: driver.Config{
			VerifierConfDepth:  0,
			SequencerConfDepth: 0,
			SequencerEnabled:   true,
			SequencerStopped:   true,
		},
		// Submitter PrivKey is set in system start for rollup nodes where sequencer = true
		RPC: rollupNode.RPCConfig{
			ListenAddr:  localhost,
			ListenPort:  0,
			EnableAdmin: true,
		},
		L1EpochPollInterval:         time.Second * 2,
		RuntimeConfigReloadInterval: time.Minute * 10,
		ConfigPersistence:           &rollupNode.DisabledConfigPersistence{},
		Sync:                        sync.Config{SyncMode: sync.CLSync},
		ConductorEnabled:            true,
		ConductorAddr:               localhost,
		ConductorPort:               port,
		ConductorRpcTimeout:         5 * time.Second,
	}
}

func waitForLeadershipChange(t *testing.T, c *conductor, leader bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			isLeader, err := c.client.Leader(ctx)
			if err != nil {
				return err
			}
			if isLeader == leader {
				return nil
			}
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func leader(t *testing.T, ctx context.Context, con *conductor) bool {
	leader, err := con.client.Leader(ctx)
	require.NoError(t, err)
	return leader
}

func healthy(t *testing.T, ctx context.Context, con *conductor) bool {
	healthy, err := con.client.SequencerHealthy(ctx)
	require.NoError(t, err)
	return healthy
}

func conductorActive(t *testing.T, ctx context.Context, con *conductor) bool {
	active, err := con.client.Active(ctx)
	require.NoError(t, err)
	return active
}

func sequencerActive(t *testing.T, ctx context.Context, rollupClient *sources.RollupClient) bool {
	active, err := rollupClient.SequencerActive(ctx)
	require.NoError(t, err)
	return active
}

func findAvailablePort(t *testing.T) int {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	for {
		select {
		case <-ctx.Done():
			t.Error("Failed to find available port")
		default:
			port := rand.Intn(65535-1024) + 1024 // Random port in the range 1024-65535
			addr := fmt.Sprintf("127.0.0.1:%d", port)
			l, err := net.Listen("tcp", addr)
			if err == nil {
				l.Close() // Close the listener and return the port if it's available
				return port
			}
		}
	}
}

func findLeader(t *testing.T, conductors map[string]*conductor) (string, *conductor) {
	for id, con := range conductors {
		if leader(t, context.Background(), con) {
			return id, con
		}
	}
	return "", nil
}
