package bonds

import (
	"context"
	"fmt"
	"math/big"

	faultTypes "github.com/ethereum-optimism/optimism/op-challenger/game/fault/types"
	"github.com/ethereum-optimism/optimism/op-service/sources/batching"
	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
	"golang.org/x/exp/maps"
)

var noBond = new(uint256.Int).Sub(uint256.NewInt(0), uint256.NewInt(1)).ToBig()

type BondContract interface {
	GetCredits(ctx context.Context, block batching.Block, recipients ...common.Address) ([]*big.Int, error)
}

// CalculateRequiredCollateral determines the minimum balance required for a fault dispute game contract in order
// to pay the outstanding bonds and credits.
// It returns the sum of unpaid bonds from claims, plus the sum of allocated but unclaimed credits.
func CalculateRequiredCollateral(ctx context.Context, contract BondContract, blockHash common.Hash, claims []faultTypes.Claim) (*big.Int, error) {
	unpaidBonds := big.NewInt(0)
	recipients := make(map[common.Address]bool)
	for _, claim := range claims {
		if noBond.Cmp(claim.Bond) != 0 {
			unpaidBonds = new(big.Int).Add(unpaidBonds, claim.Bond)
		}
		recipients[claim.Claimant] = true
		if claim.CounteredBy != (common.Address{}) {
			recipients[claim.CounteredBy] = true
		}
	}

	credits, err := contract.GetCredits(ctx, batching.BlockByHash(blockHash), maps.Keys(recipients)...)
	if err != nil {
		return nil, fmt.Errorf("failed to load credits: %w", err)
	}
	for _, credit := range credits {
		unpaidBonds = new(big.Int).Add(unpaidBonds, credit)
	}
	return unpaidBonds, nil
}
