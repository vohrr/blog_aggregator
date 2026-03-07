package command

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/vohrr/blog_aggregator/internal/database"
)

func AddFeedHandler(s *State, cmd Command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("Invalid arguments, expecting: addfeed <name> <url>")
	}
	user, err := s.Db.GetByName(context.Background(), s.Cfg.CurrentUserName)
	if err != nil {
		return err
	}
	name := cmd.Args[0]
	url := cmd.Args[1]
	//validate url?
	feedParams := database.AddFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	}
	feed, err := s.Db.AddFeed(context.Background(), feedParams)
	if err != nil {
		return err
	}
	ffParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}
	_, err = s.Db.CreateFeedFollow(context.Background(), ffParams)
	if err != nil {
		return err
	}
	fmt.Println(feed)
	return nil
}
