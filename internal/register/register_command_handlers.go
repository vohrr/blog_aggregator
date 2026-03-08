package register

import (
	"database/sql"

	"github.com/vohrr/blog_aggregator/internal/command"
	"github.com/vohrr/blog_aggregator/internal/config"
	"github.com/vohrr/blog_aggregator/internal/database"
	"github.com/vohrr/blog_aggregator/internal/middleware"
)

func Initialize(cfg *config.Config) (*command.State, command.Commands, error) {
	db, err := sql.Open("postgres", cfg.DbUrl)
	if err != nil {
		return &command.State{}, command.Commands{}, err
	}

	state := command.State{
		Cfg: cfg,
		Db:  database.New(db),
	}

	cmds := command.Commands{
		Commands: make(map[string]command.CommandHandler),
	}

	RegisterCommandHandlers(&cmds)
	return &state, cmds, nil
}

func RegisterCommandHandlers(cmds *command.Commands) {
	cmds.Register("login", command.LoginHandler)
	cmds.Register("register", command.RegisterHandler)
	cmds.Register("reset", command.ResetHandler)
	cmds.Register("users", command.UsersHandler)
	cmds.Register("agg", command.AggHandler)
	cmds.Register("feeds", command.FeedsHandler)
	cmds.Register("addfeed", middleware.LoggedIn(command.AddFeedHandler))
	cmds.Register("follow", middleware.LoggedIn(command.FollowHandler))
	cmds.Register("following", middleware.LoggedIn(command.FollowingHandler))
	cmds.Register("unfollow", middleware.LoggedIn(command.UnfollowHandler))
	cmds.Register("browse", middleware.LoggedIn(command.BrowseHandler))
}
