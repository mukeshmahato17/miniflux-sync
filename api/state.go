package api

import (
	"github.com/mukeshmahato17/miniflux-sync/diff"
	"github.com/pkg/errors"
	miniflux "miniflux.app/v2/client"
)

// GenerateDiffState generates a diff.State struct from a list of feeds.
func GenerateDiffState(feeds []*miniflux.Feed) (*diff.State, error) {
	state := diff.State{
		FeedURLsByCategoryTitle: map[string][]string{},
	}

	// Populate state with values, and create category set.
	for _, feed := range feeds {
		if feed.Category == nil {
			return nil, errors.New("feed has no category")
		}
		categoryTitle := feed.Category.Title

		state.FeedURLsByCategoryTitle[categoryTitle] = append(
			state.FeedURLsByCategoryTitle[categoryTitle], feed.FeedURL,
		)
	}

	return &state, nil
}
