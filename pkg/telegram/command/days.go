package command

import (
	"fmt"
	"strings"

	"github.com/c0rby/c3daysleft-telegrambot/pkg/answer"
	"github.com/c0rby/c3daysleft-telegrambot/pkg/event"
	"github.com/c0rby/c3daysleft-telegrambot/pkg/time"
	tb "gopkg.in/tucnak/telebot.v2"
)

// NewDays returns an instance of the days command
func NewDays(events []*event.Event) Command {
	return Days{events: events}
}

// Days implements the days command.
type Days struct {
	events []*event.Event
}

// Name returns the name of the commmand.
func (c Days) Name() string {
	return "/days"
}

// Handler returns the telegram handler for the days command.
func (c Days) Handler(bot *tb.Bot) func(*tb.Message) {
	return func(msg *tb.Message) {
		if strings.TrimSpace(msg.Text) == c.Name() {
			bot.Send(msg.Chat, "Please specify an event. Like:\n /days <event>")
			return
		}
		eventName := strings.TrimPrefix(msg.Text, c.Name()+" ")

		for _, e := range c.events {
			if e.Name == eventName {
				days := time.DaysFromNow(e.Begin)
				answer := answer.RandomAnswer(int64(days), answer.Days, e.Name)
				bot.Send(msg.Chat, answer)
				return
			}
		}
		bot.Send(msg.Chat, fmt.Sprintf("Couldn't find the event %s", eventName))
	}
}
