package command

import (
	"fmt"
	"strings"

	"github.com/c0rby/c3daysleft-telegrambot/pkg/answer"
	"github.com/c0rby/c3daysleft-telegrambot/pkg/event"
	"github.com/c0rby/c3daysleft-telegrambot/pkg/time"
	tb "gopkg.in/tucnak/telebot.v2"
)

// NewSeconds returns an instance of the days command
func NewSeconds(events []*event.Event) Command {
	return Seconds{events: events}
}

// Seconds implements the days command.
type Seconds struct {
	events []*event.Event
}

// Name returns the name of the commmand.
func (c Seconds) Name() string {
	return "/seconds"
}

// Handler returns the telegram handler for the days command.
func (c Seconds) Handler(bot *tb.Bot) func(*tb.Message) {
	return func(msg *tb.Message) {
		if strings.TrimSpace(msg.Text) == c.Name() {
			bot.Send(msg.Chat, "Please specify an event. Like:\n /seconds <event>")
			return
		}
		eventName := strings.TrimPrefix(msg.Text, c.Name()+" ")

		for _, e := range c.events {
			if e.Name == eventName {
				seconds := time.SecondsFromNow(e.Begin)
				answer := answer.RandomAnswer(seconds, answer.Seconds, e.Name)
				bot.Send(msg.Chat, answer)
				return
			}
		}
		bot.Send(msg.Chat, fmt.Sprintf("Couldn't find the event %s", eventName))
	}
}
