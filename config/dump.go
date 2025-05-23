package config

import (
	"context"

	"github.com/mukeshmahato17/miniflux-sync/kitchensink"
	"github.com/urfave/cli/v2"
)

// DumpFlags holds the flags for the dump command.
type DumpFlags struct {
	Path string
}

// Flags returns the flags for the dump command.
func (d *DumpFlags) Flags(ctx context.Context) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "path",
			Usage:       "Path to file for exported data. (optional)",
			EnvVars:     []string{"MINIFLUX_SYNC_PATH"},
			Destination: &d.Path,
			Aliases:     []string{"p"},
			Action: func(_ *cli.Context, s string) error {
				return kitchensink.ValidateFileExtension(ctx, s, []string{".yaml", ".yml"})
			},
		},
	}
}
