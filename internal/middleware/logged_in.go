package middleware

import (
	"context"
	"github.com/vohrr/blog_aggregator/internal/command"
)

func LoggedIn(handler command.CommandHandler) command.CommandHandler {
	return func(s *command.State, cmd command.Command) error {
		user, err := s.Db.GetByName(context.Background(), s.Cfg.CurrentUserName)
		if err != nil {
			return err
		}
		cmd.UserID = user.ID
		handler(s, cmd)
		return nil
	}
}
