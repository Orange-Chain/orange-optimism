package fault

import (
	"context"
	"math/big"

	"github.com/ethereum-optimism/optimism/op-bindings/bindings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
)

type FaultDisputeGameCaller interface {
	Status(opts *bind.CallOpts) (uint8, error)
	ClaimDataLen(opts *bind.CallOpts) (*big.Int, error)
}

type FaultCaller struct {
	FaultDisputeGameCaller
	log     log.Logger
	fdgAddr common.Address
}

func NewFaultCaller(fdgAddr common.Address, caller FaultDisputeGameCaller, log log.Logger) *FaultCaller {
	return &FaultCaller{
		caller,
		log,
		fdgAddr,
	}
}

func NewFaultCallerFromBindings(fdgAddr common.Address, client *ethclient.Client, log log.Logger) (*FaultCaller, error) {
	caller, err := bindings.NewFaultDisputeGameCaller(fdgAddr, client)
	if err != nil {
		return nil, err
	}
	return &FaultCaller{
		caller,
		log,
		fdgAddr,
	}, nil
}

// LogGameInfo logs the game info.
func (fc *FaultCaller) LogGameInfo(ctx context.Context) {
	status, err := fc.GetGameStatus(ctx)
	if err != nil {
		fc.log.Error("failed to get game status", "err", err)
		return
	}
	claimLen, err := fc.GetClaimDataLength(ctx)
	if err != nil {
		fc.log.Error("failed to get claim count", "err", err)
		return
	}
	fc.log.Info("Game info", "addr", fc.fdgAddr, "claims", claimLen, "status", GameStatusString(status))
}

// GetGameStatus returns the current game status.
// 0: In Progress
// 1: Challenger Won
// 2: Defender Won
func (fc *FaultCaller) GetGameStatus(ctx context.Context) (GameStatus, error) {
	status, err := fc.Status(&bind.CallOpts{Context: ctx})
	return GameStatus(status), err
}

// GetClaimDataLength returns the number of claims in the game.
func (fc *FaultCaller) GetClaimDataLength(ctx context.Context) (*big.Int, error) {
	return fc.ClaimDataLen(&bind.CallOpts{Context: ctx})
}

func (fc *FaultCaller) LogClaimDataLength(ctx context.Context) {
	claimLen, err := fc.GetClaimDataLength(ctx)
	if err != nil {
		fc.log.Error("failed to get claim count", "err", err)
		return
	}
	fc.log.Info("Number of claims", "length", claimLen)
}

// GameStatusString returns the current game status as a string.
func GameStatusString(status GameStatus) string {
	switch status {
	case GameStatusInProgress:
		return "In Progress"
	case GameStatusChallengerWon:
		return "Challenger Won"
	case GameStatusDefenderWon:
		return "Defender Won"
	default:
		return "Unknown"
	}
}
