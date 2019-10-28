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
var daysMessages []string = []string{
	"There are %s days remaining until CCCongress",
	"You have to sleep %s times until CCCongress",
	"Only %s nights left to hack something great for CCCongress",
	"In %s days you will drink more Mate than on all other days of the year. Cheers!",
	"You have %s nights to sleep well to be fit for the CCCongress",
	"In %s you'll have a good excuse for leaving christmas with your family!"}
var secondsMessages []string = []string{
	"There are %s seconds remaining until CCCongress",
	"Just count to %s and the CCCongress will start!"}

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
	return daysTilCongressFrom(now)
}

func daysTilCongressFrom(now time.Time) int {
	day1 := time.Date(now.Year(), time.December, 27, 0, 0, 0, 0, time.UTC)
	if now.After(day1) {
		day1 = time.Date(now.Year()+1, time.December, 27, 0, 0, 0, 0, time.UTC)
	}
	return daysBetween(now, day1)
}

func secondsTilCongress() int64 {
	now := time.Now()
	startDay1 := time.Date(now.Year(), time.December, 27, 10, 0, 0, 0, time.UTC)
	return startDay1.Unix() - now.Unix()
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
		formattedDays := fmt.Sprintf(format, days)

		message := daysMessages[r.Intn(len(daysMessages))]
		b.Send(m.Chat, fmt.Sprintf(message, formattedDays))
	})

	b.Handle("/seconds", func(m *tb.Message) {
		seconds := secondsTilCongress()
		s := rand.NewSource(time.Now().UnixNano())
		r := rand.New(s)

		format := formats[r.Intn(len(formats))]
		formattedSeconds := fmt.Sprintf(format, seconds)

		message := secondsMessages[r.Intn(len(secondsMessages))]
		b.Send(m.Chat, fmt.Sprintf(message, formattedSeconds))
	})

	b.Start()
}
