package main

import (
	"fmt"
	"os"

	"github.com/googollee/gobase/cli"
)

func main() {
	cli.App.Name = "cli_test"
	cli.App.Usage = "An example of how to use cli."

	var config string
	cli.App.Flags = []cli.Flag{
		&cli.FlagString{
			Name:        "config",
			Usage:       "the path of configuration file",
			Value:       "./config",
			Destination: &config,
		},
	}

	cli.App.Before(func(ctx *cli.Context) error {
		fmt.Println("before")
		ctx.StoreValue("config", config)
		return nil
	})

	cli.App.Run(os.Args)
}
