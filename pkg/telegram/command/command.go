package command

import tb "gopkg.in/tucnak/telebot.v2"

// Command defines the interface for a telegram command.
type Command interface {
	Name() string
	Handler(*tb.Bot) func(*tb.Message)
}
