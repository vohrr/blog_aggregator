package command

import (
	"context"
	"fmt"
)

func LoginHandler(s *State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("Invalid Arguments, expecting login <username>")
	}
	username := cmd.Args[0]
	_, err := s.Db.GetByName(context.Background(), username)
	if err != nil {
		return fmt.Errorf("User not found")
	}
	fmt.Printf("Logging in as %s...\n", username)
	err = s.Cfg.SetUser(username)
	if err != nil {
		return err
	}
	fmt.Printf("Username set to %s\n", username)
	return nil
}

