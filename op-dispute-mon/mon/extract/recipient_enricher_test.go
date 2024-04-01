package extract

import (
	"context"
	"testing"

	"github.com/ethereum-optimism/optimism/op-service/sources/batching/rpcblock"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestRecipientEnricher(t *testing.T) {
	game, recipients := makeTestGame()
	game.Recipients = make(map[common.Address]bool)
	enricher := NewRecipientEnricher()
	caller := &mockGameCaller{}
	ctx := context.Background()
	err := enricher.Enrich(ctx, rpcblock.Latest, caller, game)
	require.NoError(t, err)
	for _, recipient := range recipients {
		require.Contains(t, game.Recipients, recipient)
	}
}
