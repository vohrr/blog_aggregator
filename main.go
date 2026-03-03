package main

import (
	"fmt"

	"github.com/vohrr/blog_aggregator/internal/config"
)

func main() {

	fmt.Println("Loading configuration data....")

	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err.Error())
		return
	} else {
		fmt.Println("Configuration loaded")
	}

	err = cfg.SetUser("vohrr")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	cfg, err = config.Read()
	if err != nil {
		fmt.Println(err.Error())
		return
	} else {
		fmt.Println(cfg.DbUrl)
		fmt.Println(cfg.CurrentUserName)
	}
}
