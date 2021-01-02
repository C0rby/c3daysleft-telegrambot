package telegram

import (
	"github.com/c0rby/c3daysleft-telegrambot/pkg/telegram/command"
)

// Option defines a single option function.
type Option func(o *Options)

// Options defines the available options for this package.
type Options struct {
	APIToken string
	Commands []command.Command
}

func newOptions(opts ...Option) Options {
	opt := Options{}
	for _, o := range opts {
		o(&opt)
	}
	return opt
}

// APIToken provides a function to set the ApiToken option.
func APIToken(token string) Option {
	return func(o *Options) {
		o.APIToken = token
	}
}

// Commands provides a function to add Commands.
func Commands(cmd ...command.Command) Option {
	return func(o *Options) {
		o.Commands = append(o.Commands, cmd...)
	}
}
