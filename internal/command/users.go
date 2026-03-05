package command

import (
	"context"
	"fmt"
)

func UsersHandler(s *State, cmd Command) error {
	users, err := s.Db.GetUsers(context.Background())
	if err != nil {
		return err
	}

	for _, user := range users {
		if s.Cfg.CurrentUserName == user.Name {
			fmt.Printf("%s (current)\n", user.Name)
		}
		fmt.Printf("%s\n", user.Name)
	}
	return nil
}
