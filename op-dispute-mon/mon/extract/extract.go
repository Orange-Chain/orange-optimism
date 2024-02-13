package extract

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"

	gameTypes "github.com/ethereum-optimism/optimism/op-challenger/game/types"
	monTypes "github.com/ethereum-optimism/optimism/op-dispute-mon/mon/types"
)

type CreateGameCaller func(game gameTypes.GameMetadata) (GameCaller, error)
type FactoryGameFetcher func(ctx context.Context, blockHash common.Hash, earliestTimestamp uint64) ([]gameTypes.GameMetadata, error)

type Extractor struct {
	logger         log.Logger
	createContract CreateGameCaller
	fetchGames     FactoryGameFetcher
}

func NewExtractor(logger log.Logger, creator CreateGameCaller, fetchGames FactoryGameFetcher) *Extractor {
	return &Extractor{
		logger:         logger,
		createContract: creator,
		fetchGames:     fetchGames,
	}
}

func (e *Extractor) Extract(ctx context.Context, blockHash common.Hash, minTimestamp uint64) ([]monTypes.EnrichedGameData, error) {
	// Fetch games from the factory
	games, err := e.fetchGames(ctx, blockHash, minTimestamp)
	if err != nil {
		return nil, fmt.Errorf("failed to load games: %w", err)
	}

	// For each game, create a game caller
	var enrichedGames []monTypes.EnrichedGameData
	for _, game := range games {
		caller, err := e.createContract(game)
		if err != nil {
			return nil, fmt.Errorf("failed to create game caller: %w", err)
		}
		l2BlockNum, rootClaim, status, err := caller.GetGameMetadata(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch game metadata: %w", err)
		}
		// TODO: Fetch the expected root from the output rollup
		enrichedGames = append(enrichedGames, monTypes.EnrichedGameData{
			L2BlockNumber: l2BlockNum,
			RootClaim:     rootClaim,
			Status:        status,
		})
	}
	return enrichedGames, nil
}
