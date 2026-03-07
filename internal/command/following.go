package command

import (
	"context"
	"fmt"

	"github.com/vohrr/blog_aggregator/internal/database"
)

func FollowingHandler(s *State, cmd Command, user database.User) error {

	following, err := s.Db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	for _, feed := range following {
		fmt.Println(feed.FeedName)
	}

	return nil
}
