package command

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/vohrr/blog_aggregator/internal/database"
)

func FollowHandler(s *State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("Invalid command arguments, expecting: follow <url>")
	}
	url := cmd.Args[0]
	feed, err := s.Db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("Feed for %s does not exist, use <addFeed> command to create it", url)
		} else {
			return err
		}
	}
	feedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    cmd.UserID,
		FeedID:    feed.ID,
	}
	feedFollowResult, err := s.Db.CreateFeedFollow(context.Background(), feedFollowParams)
	if err != nil {
		return err
	}
	fmt.Println(feedFollowResult.UserName)
	fmt.Println(feedFollowResult.FeedName)
	return nil
}
