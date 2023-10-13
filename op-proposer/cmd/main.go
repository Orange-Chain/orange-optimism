package main

import (
	"context"
	"fmt"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/ethereum-optimism/optimism/op-proposer/cmd/doc"
	"github.com/ethereum-optimism/optimism/op-proposer/flags"
	"github.com/ethereum-optimism/optimism/op-proposer/proposer"
	"github.com/ethereum-optimism/optimism/op-service/cliapp"
	oplog "github.com/ethereum-optimism/optimism/op-service/log"
	"github.com/ethereum-optimism/optimism/op-service/opio"
	"github.com/ethereum/go-ethereum/log"
)

var (
	Version   = "v0.10.14"
	GitCommit = ""
	GitDate   = ""
)

func main() {
	oplog.SetupDefaults()

	// Invoke cancel when an interrupt is received.
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		opio.BlockOnInterrupts()
		cancel()
	}()

	app := cli.NewApp()
	app.Flags = cliapp.ProtectFlags(flags.Flags)
	app.Version = fmt.Sprintf("%s-%s-%s", Version, GitCommit, GitDate)
	app.Name = "op-proposer"
	app.Usage = "L2Output Submitter"
	app.Description = "Service for generating and submitting L2 Output checkpoints to the L2OutputOracle contract"
	app.Action = curryMain(Version)
	app.Commands = []*cli.Command{
		{
			Name:        "doc",
			Subcommands: doc.Subcommands,
		},
	}

	err := app.RunContext(ctx, os.Args)
	if err != nil {
		log.Crit("Application failed", "message", err)
	}
}

// curryMain transforms the proposer.Main function into an app.Action
// This is done to capture the Version of the proposer.
func curryMain(version string) func(ctx *cli.Context) error {
	return func(ctx *cli.Context) error {
		return proposer.Main(version, ctx)
	}
}
