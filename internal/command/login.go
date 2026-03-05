package command

import "fmt"

func LoginHandler(s *State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("Invalid Arguments, expecting <username>")
	}
	username := cmd.Args[0]
	fmt.Printf("Logging in as %s...\n", username)
	err := s.Cfg.SetUser(username)
	if err != nil {
		return err
	}
	fmt.Printf("Username set to %s\n", username)
	return nil
}
