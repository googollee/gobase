package cli

import (
	"context"

	"github.com/urfave/cli/v2"
)

type Context struct {
	*cli.Context
}

func (c *Context) SetFlag(name, value string) error {
	return c.Context.Set(name, value)
}

func (c *Context) StoreValue(key string, value interface{}) {
	meta := &c.Context.App.Metadata
	if *meta == nil {
		*meta = make(map[string]interface{})
	}
	(*meta)[key] = value
}

func (c *Context) Value(key interface{}) interface{} {
	str, ok := key.(string)
	if !ok {
		return c.Context.Context.Value(key)
	}

	meta := &c.Context.App.Metadata
	if *meta == nil {
		return c.Context.Context.Value(key)
	}

	ret, ok := (*meta)[str]
	if !ok {
		return c.Context.Context.Value(key)
	}

	return ret
}

// Check &Context{} could be compatible with context.Context.
var _ context.Context = &Context{}
