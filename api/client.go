package api

import (
	"context"

	"github.com/mukeshmahato17/miniflux-sync/config"
	"github.com/mukeshmahato17/miniflux-sync/log"
	"github.com/pkg/errors"
	miniflux "miniflux.app/v2/client"
)

// Client creates a new Miniflux API client, whilst checking the health of the Miniflux instance.
func Client(ctx context.Context, cfg *config.GlobalFlags) (*miniflux.Client, error) {
	log.Info(ctx, "connecting to miniflux instance")
	client := miniflux.New(cfg.Endpoint, cfg.APIKey)

	log.Info(ctx, "checking health of miniflux instance")
	if err := client.Healthcheck(); err != nil {
		return nil, errors.Wrap(err, "checking health of miniflux instance")
	}

	return client, nil
}
