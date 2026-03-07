package command

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/vohrr/blog_aggregator/internal/database"
)

func UnfollowHandler(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("Invalid arguments: expecting %s <url>", cmd.Name)
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

	unfollowParams := database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}

	err = s.Db.DeleteFeedFollow(context.Background(), unfollowParams)
	if err != nil {
		return err
	}
	return nil
}
