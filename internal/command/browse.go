package command

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

	"github.com/vohrr/blog_aggregator/internal/database"
)

func BrowseHandler(s *State, cmd Command, user database.User) error {
	var limit int64
	var err error
	if len(cmd.Args) != 2 {
		limit = 2
	} else {
		limit, err = strconv.ParseInt(cmd.Args[1], 10, 32)
		if err != nil {
			limit = 2
		}
	}
	getPostsParams := database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	}

	posts, err := s.Db.GetPostsForUser(context.Background(), getPostsParams)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	for _, post := range posts {
		fmt.Printf("Title: %s\n", post.Title)
		fmt.Printf("Link: %s\n", post.Url)
		fmt.Printf("Description: %s\n", post.Description.String)
	}
	return nil
}
