package main

import (
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/vohrr/blog_aggregator/internal/command"
	"github.com/vohrr/blog_aggregator/internal/config"
)

func main() {

	fmt.Println("Loading configuration data....")

	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	} else {
		fmt.Println("Configuration loaded")
		fmt.Printf("Logged in as %s\n", cfg.CurrentUserName)
	}
	state, cmds, err := command.Initialize(cfg)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	args := os.Args
	if len(args) < 2 {
		fmt.Println("Expected gator <command>")
		os.Exit(1)
	}
	args = args[1:]
	cmd, err := command.Parse(args)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Printf("Running %s...\n", cmd.Name)
	err = cmds.Run(state, cmd)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
