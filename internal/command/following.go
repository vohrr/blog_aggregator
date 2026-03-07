package command

import (
	"context"
	"fmt"
)

func FollowingHandler(s *State, cmd Command) error {

	user, err := s.Db.GetByName(context.Background(), s.Cfg.CurrentUserName)
	if err != nil {
		return err
	}

	following, err := s.Db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	for _, feed := range following {
		fmt.Println(feed.FeedName)
	}

	return nil
}
