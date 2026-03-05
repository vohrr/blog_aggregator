package command

import (
	"fmt"

	"github.com/vohrr/blog_aggregator/internal/config"
)

type Command struct {
	Name string
	Args []string
}

type State struct {
	Cfg *config.Config
}

type CommandHandler func(s *State, cmd Command) error

type Commands struct {
	Commands map[string]CommandHandler
}

func Initialize(cfg *config.Config) (*State, Commands) {
	state := State{
		Cfg: cfg,
	}
	cmds := Commands{
		Commands: make(map[string]CommandHandler),
	}
	cmds.Register("login", LoginHandler)

	return &state, cmds
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
