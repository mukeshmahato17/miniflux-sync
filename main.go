package main

import (
	"log"
	"os"

	"github.com/mukeshmahato17/subflux/config"
	"github.com/mukeshmahato17/subflux/sync"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

func main() {
	cfg := config.New()

	app := &cli.App{
		Name:  "subflux-sync",
		Usage: "Manage and sync your Subflux feeds with YAML. ",
		Flags: cfg.Flags(),
		Action: func(ctx *cli.Context) error {
			if err := sync.Sync(); err != nil {
				return errors.Wrap(err, "syncing config")
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Printf("error: %+v\n", err)
		os.Exit(1)
	}
}
