package command

import (
	"context"
	"fmt"
	"time"

	"github.com/vohrr/blog_aggregator/internal/rss"
)

func AggHandler(s *State, cmd Command) error {
	// only fetching 1 url as proof of concept right now
	if len(cmd.Args) != 1 {
		return fmt.Errorf("Invalid argument, expecting %s <time_between_requests>", cmd.Name)
	}
	time_between_requests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return err
	}
	fmt.Printf("Collecting feeds every %d\n", time_between_requests)

	schedule := time.Tick(time_between_requests)

	for _ = range schedule {
		scrapeFeeds(s)
	}
	return nil
}

func scrapeFeeds(s *State) error {
	feed, err := s.Db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}
	err = s.Db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		return err
	}
	rss_feed, err := rss.FetchFeed(context.Background(), feed.Url)
	if err != nil {
		return err
	}

	for _, item := range rss_feed.Channel.Item {
		fmt.Println(item.Title)
	}
	return nil
}
