package command

import (
	"bytes"

	"github.com/c0rby/c3daysleft-telegrambot/pkg/event"
	tb "gopkg.in/tucnak/telebot.v2"
)

// NewEvents returns an instance of the days command
func NewEvents(events []*event.Event) Command {
	return Events{events: events}
}

// Events implements the days command.
type Events struct {
	events []*event.Event
}

// Name returns the name of the commmand.
func (c Events) Name() string {
	return "/events"
}

// Handler returns the telegram handler for the days command.
func (c Events) Handler(bot *tb.Bot) func(*tb.Message) {
	return func(msg *tb.Message) {
		var b bytes.Buffer
		for _, e := range c.events {
			b.WriteString(e.String())
			b.WriteRune('\n')
		}
		bot.Send(msg.Chat, b.String())
	}
}
