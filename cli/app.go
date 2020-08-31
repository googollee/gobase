// Package cli provides a easy way to organize a cli program.
// For advance usage, check https://pkg.go.dev/github.com/urfave/cli/v2.
// Check the example in ./example.
package cli

import (
	"fmt"
	"runtime"

	"github.com/urfave/cli/v2"

	"github.com/googollee/gobase"
)

func init() {
	cli.AppHelpTemplate += fmt.Sprintf(`
COMPILED WITH:
   go version: %q
   code hash: %q
`, runtime.Version(), gobase.CodeHash)
}

// Apppplication is the main structure of a cli application.
type Application struct {
	cli.App
}

// App is the default application.
var App Application

// Command adds a sub command to the application.
func (a *Application) Command(name, usage string, f CommandHandler) *Command {
	ret := &Command{
		Command: &cli.Command{
			Name:  name,
			Usage: usage,
			Action: func(ctx *cli.Context) error {
				return f(&Context{
					Context: ctx,
				})
			},
		},
	}
	a.Commands = append(a.Commands, ret.Command)
	return ret
}

// CommandHandler is a function to handle something.
type CommandHandler func(ctx *Context) error

// Before hooks a handler to execute before any sub-subcommands are run, but after the context is ready.
// If a non-nil error is returned, no sub-subcommands are run.
func (a *Application) Before(f CommandHandler) {
	a.App.Before = func(ctx *cli.Context) error {
		return f(&Context{
			Context: ctx,
		})
	}
}

// After hooks a handler to execute after any subcommands are run, but after the subcommand has finished.
// It is run even if Action() panics.
func (a *Application) After(f CommandHandler) {
	a.App.After = func(ctx *cli.Context) error {
		return f(&Context{
			Context: ctx,
		})
	}
}

// Actions sets a handler running when no subcommands giving.
func (a *Application) Action(f CommandHandler) {
	a.App.Action = func(ctx *cli.Context) error {
		return f(&Context{
			Context: ctx,
		})
	}
}
