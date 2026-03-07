package middleware

import (
	"context"
	"github.com/vohrr/blog_aggregator/internal/command"
)

func LoggedIn(handler command.AuthCommandHandler) command.CommandHandler {
	return func(s *command.State, cmd command.Command) error {
		user, err := s.Db.GetByName(context.Background(), s.Cfg.CurrentUserName)
		if err != nil {
			return err
		}
		handler(s, cmd, user)
		return nil
	}
}
