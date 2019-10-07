package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

func daysTilCongress() int {
	now := time.Now()
	year, _, _ := now.Date()
	day1 := time.Date(year, time.December, 27, 0, 0, 0, 0, time.UTC)
	// calculate total number of days
	duration := day1.Sub(now)
	return int(math.Ceil(duration.Hours() / 24))
}

func main() {
	var (
		apiToken = flag.String("apitoken", "", "Telegram API Token")
	)
	flag.Parse()

	if len(*apiToken) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	b, err := tb.NewBot(tb.Settings{
		Token: *apiToken,
		// You can also set custom API URL. If field is empty it equals to "https://api.telegram.org"
		// URL:    "",
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/days", func(m *tb.Message) {
		days := daysTilCongress()
		b.Send(m.Chat, fmt.Sprintf("There are %d days remaining until CCCongress", days))
	})

	b.Start()
}
