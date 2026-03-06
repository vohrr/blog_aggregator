package command

import (
	"context"
	"fmt"

	"github.com/vohrr/blog_aggregator/internal/rss"
)

func AggHandler(s *State, cmd Command) error {
	// only fetching 1 url as proof of concept right now

	rss_feed, err := rss.FetchFeed(context.Background(), "https://wagslane.dev/index.xml")
	if err != nil {
		return err
	}

	fmt.Println(rss_feed)
	return nil
}
