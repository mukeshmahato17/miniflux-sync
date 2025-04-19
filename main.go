package main

import (
	"log"
	"os"

	"github.com/mukeshmahato17/subflux/cmd"
	"github.com/mukeshmahato17/subflux/config"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

//go:embed VERSION
var version string

func main() {
	cfg := config.New(version)

	app := &cli.App{
		Name:    "subflux-sync",
		Usage:   "Manage and sync your Subflux feeds with YAML. ",
		Version: cfg.Version,
		Flags:   cfg.Flags(),
		Action: func(ctx *cli.Context) error {
			if err := cmd.Sync(cfg); err != nil {
				return errors.Wrap(err, "running sync command")
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Printf("error: %+v\n", err)
		os.Exit(1)
	}
}
