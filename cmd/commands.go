package cmd

import (
	"github.com/mukeshmahato17/miniflux-sync/config"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

// Commands returns the commands for the CLI.
func Commands(cfg *config.GlobalFlags) []*cli.Command {
	dumpFlags := &config.DumpFlags{}
	syncFlags := &config.SyncFlags{}

	return []*cli.Command{
		{
			Name:    "sync",
			Aliases: []string{"s"},
			Usage:   "Update Miniflux using a local YAML file.",
			Flags:   syncFlags.Flags(),
			Action: func(ctx *cli.Context) error {
				if err := sync(cfg, syncFlags); err != nil {
					return errors.Wrap(err, "running sync command")
				}

				return nil
			},
		},
		{
			Name:    "dump",
			Aliases: []string{"d"},
			Flags:   dumpFlags.Flags(),
			Usage:   "Dump the current remote Miniflux state to your machine.",
			Action: func(ctx *cli.Context) error {
				if err := dump(cfg, dumpFlags); err != nil {
					return errors.Wrap(err, "running dump command")
				}

				return nil
			},
		},
	}
}
