package cli

import "github.com/urfave/cli/v2"

type Command struct {
	*cli.Command
}

func (c *Command) Before(f CommandHandler) {
	c.Command.Before = func(ctx *cli.Context) error {
		return f(&Context{
			Context: ctx,
		})
	}
}

func (c *Command) After(f CommandHandler) {
	c.Command.After = func(ctx *cli.Context) error {
		return f(&Context{
			Context: ctx,
		})
	}
}

func (c *Command) Action(f CommandHandler) {
	c.Command.Action = func(ctx *cli.Context) error {
		return f(&Context{
			Context: ctx,
		})
	}
}
