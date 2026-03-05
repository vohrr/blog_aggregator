package command

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/vohrr/blog_aggregator/internal/database"
)

func RegisterHandler(s *State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("Invalid Arguments, expecting register <username>")
	}

	_, err := s.Db.GetByName(context.Background(), cmd.Args[0])
	if err == nil {
		return fmt.Errorf("A user with that name already exists")
	}

	userParams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
	}

	user, err := s.Db.CreateUser(context.Background(), userParams)
	if err != nil {
		return err
	}
	fmt.Printf("Successfully registered new user %s.\n %s\n", user.Name, user)
	s.Cfg.SetUser(user.Name)
	return nil
}
