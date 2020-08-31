package main

import (
	"fmt"

	"github.com/googollee/gobase/cli"
)

func init() {
	cmd := cli.App.Command("cron", "run cron jobs", cron)
	cmd.Flags = []cli.Flag{
		&cli.FlagString{
			Name:  "rss",
			Usage: "rss url",
			Value: "",
		},
	}
}

func cron(ctx *cli.Context) error {
	cfg := ctx.Value("config")
	rss := ctx.String("rss")
	fmt.Println("cron", cfg, rss)
	return nil
}
