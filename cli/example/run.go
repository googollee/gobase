package main

import (
	"fmt"

	"github.com/googollee/gobase/cli"
)

func init() {
	cmd := cli.App.Command("run", "run jobs", run)
	cmd.Flags = []cli.Flag{
		&cli.FlagBool{
			Name:  "dry_run",
			Usage: "won't submit if true",
			Value: false,
		},
	}
}

func run(ctx *cli.Context) error {
	cfg := ctx.Value("config")
	dryRun := ctx.Bool("dry_run")
	fmt.Println("run", cfg, dryRun)
	return nil
}
