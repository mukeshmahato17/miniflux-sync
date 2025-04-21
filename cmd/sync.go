package cmd

import (
	"log"

	"github.com/mukeshmahato17/miniflux-sync/api"
	"github.com/mukeshmahato17/miniflux-sync/config"
	"github.com/pkg/errors"
)

// Sync is the entry point for the sync command in the CLI.
func sync(cfg *config.GlobalFlags, _ *config.SyncFlags) error {
	client, err := api.Client(cfg)
	if err != nil {
		return errors.Wrap(err, "creating miniflux client")
	}

	feedsByCategory, err := api.GetFeedsByCategories(client)
	if err != nil {
		return errors.Wrap(err, "getting feeds by category")
	}

	log.Printf("feeds: %d\n", len(feedsByCategory.Feeds()))
	log.Printf("categories: %d\n", len(feedsByCategory))

	return nil
}
