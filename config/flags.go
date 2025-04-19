package config

import "github.com/urfave/cli/v2"

// Config holds the configuration for the CLI.
type Config struct {
	Endpoint string
	Username string
	Password string
	Version  string
}

// New is a convenience function for creating a new Config.
func New(v string) *Config {
	return &Config{
		Version: v,
	}
}

// Flags returns the flags for the CLI.
func (c *Config) Flags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "endpoint",
			Usage:       "Subflux API endpoint.",
			EnvVars:     []string{"SUBFLUX_SYNC_ENDPOINT"},
			Destination: &c.Endpoint,
			Aliases:     []string{"e"},
			Required:    true,
		},
		&cli.StringFlag{
			Name:        "username",
			Usage:       "Subflux API username.",
			EnvVars:     []string{"SUBFLUX_SYNC_USERNAME"},
			Destination: &c.Username,
			Aliases:     []string{"u"},
			Required:    true,
		},
		&cli.StringFlag{
			Name:        "password",
			Usage:       "Subflux API password.",
			EnvVars:     []string{"SUBFLUX_SYNC_PASSWORD"},
			Destination: &c.Password,
			Aliases:     []string{"p"},
			Required:    true,
		},
	}
}
