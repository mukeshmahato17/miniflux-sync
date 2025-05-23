package cmd

import (
	"context"
	"path/filepath"
	"strings"

	"github.com/mukeshmahato17/miniflux-sync/api"
	"github.com/mukeshmahato17/miniflux-sync/config"
	"github.com/mukeshmahato17/miniflux-sync/diff"
	"github.com/mukeshmahato17/miniflux-sync/log"
	"github.com/mukeshmahato17/miniflux-sync/parse"
	"github.com/pkg/errors"
	miniflux "miniflux.app/v2/client"
)

// Sync is the entry point for the sync command in the CLI.
func sync( //nolint:cyclop,funlen
	ctx context.Context, flags *config.SyncFlags, client *miniflux.Client,
) error {
	var localState *diff.State
	var err error

	switch filepath.Ext(flags.Path) {
	case ".yaml", ".yml":
		localState, err = parse.Parse(ctx, flags.Path)
		if err != nil {
			return errors.Wrap(err, "loading data from yaml file")
		}

		// TODO: Implement logic for YAML.

	default:
		return errors.New("invalid file extension") // Should never happen, as we validate flag before.

	}

	log.Info(ctx, "local feeds", log.Metadata{
		"count": len(localState.FeedURLs()),
	})
	log.Info(ctx, "local categories", log.Metadata{
		"count": len(localState.CategoryTitles()),
	})

	feeds, categories, err := api.FetchData(ctx, client)
	if err != nil {
		return errors.Wrap(err, "fetching data")
	}

	remoteState, err := api.GenerateDiffState(feeds)
	if err != nil {
		return errors.Wrap(err, "generating remote state")
	}

	log.Info(ctx, "remote feeds", log.Metadata{
		"count": len(remoteState.FeedURLs()),
	})
	log.Info(ctx, "remote categories", log.Metadata{
		"count": len(remoteState.CategoryTitles()),
	})

	actions, err := diff.CalculateDiff(localState, remoteState)
	if err != nil {
		return errors.Wrap(err, "calculating diff")
	}

	if len(actions) == 0 {
		log.Info(ctx, "no actions to perform")
		return nil
	}

	log.Info(ctx, "actions to perform", log.Metadata{
		"count": len(actions),
	})

	for _, action := range actions {
		log.Info(ctx, strings.ToLower(string(action.Type)), log.Metadata{
			"category_title": action.CategoryTitle,
			"feed_url":       action.FeedURL,
		})
	}

	if flags.DryRun {
		log.Info(ctx, "dry run complete")
		return nil
	}

	if err := api.Update(ctx, client, actions, feeds, categories); err != nil {
		return errors.Wrap(err, "performing actions")
	}

	return nil
}
