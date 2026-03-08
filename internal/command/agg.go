package command

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/vohrr/blog_aggregator/internal/database"
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
		err = savePost(s, item, feed)
		if err != nil {
			//properly handle the error without breaking the feedloop
			fmt.Printf("Error saving post - Title: %s, Error: %s", item.Title, err.Error())
		}
	}
	return nil
}

func savePost(s *State, item rss.RSSItem, feed database.Feed) error {
	var pubDate sql.NullTime
	pubtime, err := time.Parse(time.DateTime, item.PubDate)
	if err != nil {
		pubDate = sql.NullTime{Valid: false}
	} else {
		pubDate = sql.NullTime{Time: pubtime, Valid: true}
	}
	postParams := database.CreatePostParams{
		ID:          uuid.New(),
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
		Title:       item.Title,
		Url:         item.Link,
		Description: sql.NullString{String: item.Description},
		PublishedAt: pubDate,
		FeedID:      feed.ID,
	}

	_, err = s.Db.CreatePost(context.Background(), postParams)
	if err != nil {
		return err
	}
	return nil
}
