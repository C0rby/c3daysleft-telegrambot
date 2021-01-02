package telegram

import (
	"log"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

// NewBot returns an instance of a telegram bot
func NewBot(opts ...Option) *tb.Bot {
	options := newOptions(opts...)

	b, err := tb.NewBot(tb.Settings{
		Token:  options.APIToken,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatalf("Could not create the telegram bot err: %s", err.Error())
	}

	for _, cmd := range options.Commands {
		b.Handle(cmd.Name(), cmd.Handler(b))
	}

	return b
}
