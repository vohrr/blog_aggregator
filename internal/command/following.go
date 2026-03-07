package command

import (
	"context"
	"fmt"
)

func FollowingHandler(s *State, cmd Command) error {

	following, err := s.Db.GetFeedFollowsForUser(context.Background(), cmd.UserID)
	if err != nil {
		return err
	}

	for _, feed := range following {
		fmt.Println(feed.FeedName)
	}

	return nil
}
