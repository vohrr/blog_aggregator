package command

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/vohrr/blog_aggregator/internal/config"
	"github.com/vohrr/blog_aggregator/internal/database"
)

type Command struct {
	Name string
	Args []string
}

type State struct {
	Cfg *config.Config
	Db  *database.Queries
}

type CommandHandler func(s *State, cmd Command) error

type Commands struct {
	Commands map[string]CommandHandler
}

func Initialize(cfg *config.Config) (*State, Commands, error) {
	db, err := sql.Open("postgres", cfg.DbUrl)
	if err != nil {
		return &State{}, Commands{}, err
	}

	state := State{
		Cfg: cfg,
		Db:  database.New(db),
	}

	cmds := Commands{
		Commands: make(map[string]CommandHandler),
	}
	cmds.Register("login", LoginHandler)
	cmds.Register("register", RegisterHandler)

	return &state, cmds, nil
}

func Parse(args []string) (Command, error) {
	commandName := args[0]
	if len(commandName) == 0 {
		return Command{}, fmt.Errorf("Invalid command")
	}
	otherArgs := args[1:]
	cmd := Command{
		Name: commandName,
		Args: otherArgs,
	}
	return cmd, nil
}

func (c *Commands) Run(s *State, cmd Command) error {
	if handler, ok := c.Commands[cmd.Name]; ok {
		err := handler(s, cmd)
		if err != nil {
			return err
		}
		return nil
	} else {
		return fmt.Errorf("Command not registered")
	}
}

func (c *Commands) Register(name string, f CommandHandler) error {
	if _, ok := c.Commands[name]; ok {
		return fmt.Errorf("Command already registered")
	} else {
		c.Commands[name] = f
		return nil
	}
}
