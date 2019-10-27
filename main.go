package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

var formats []string = []string{"%d", "%#x", "%#o", "%b"}

func daysBetween(a, b time.Time) int {
	if a.After(b) {
		a, b = b, a
	}

	days := -a.YearDay()
	for year := a.Year(); year < b.Year(); year++ {
		days += time.Date(year, time.December, 31, 0, 0, 0, 0, time.UTC).YearDay()
	}
	days += b.YearDay()

	return days
}

func daysTilCongress() int {
	now := time.Now()
	day1 := time.Date(now.Year(), time.December, 27, 0, 0, 0, 0, time.UTC)
	return daysBetween(now, day1)
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
		s := rand.NewSource(time.Now().UnixNano())
		r := rand.New(s)
		format := formats[r.Intn(len(formats))]
		b.Send(m.Chat, fmt.Sprintf("There are "+format+" days remaining until CCCongress", days))
	})

	b.Handle("/nights", func(m *tb.Message) {
		days := daysTilCongress() - 1
		s := rand.NewSource(time.Now().UnixNano())
		r := rand.New(s)
		format := formats[r.Intn(len(formats))]
		b.Send(m.Chat, fmt.Sprintf("You have to sleep "+format+" times until CCCongress", days))
	})

	b.Start()
}
