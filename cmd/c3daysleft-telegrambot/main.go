package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"sort"

	"github.com/c0rby/c3daysleft-telegrambot/pkg/event"
	"github.com/c0rby/c3daysleft-telegrambot/pkg/telegram"
	"github.com/c0rby/c3daysleft-telegrambot/pkg/telegram/command"
)

var formats = [...]string{"%d", "%#x", "%#o", "%#b"}
var daysMessages = [...]string{
	"There are %s days remaining until CCCongress",
	"You have to sleep %s times until CCCongress",
	"Only %s nights left to hack something great for CCCongress",
	"In %s days you will drink more Mate than on all other days of the year. Cheers!",
	"You have %s nights to sleep well to be fit for the CCCongress",
	"In %s days you'll have a good excuse for leaving christmas with your family!"}
var secondsMessages = [...]string{
	"There are %s seconds remaining until CCCongress",
	"Just count to %s and the CCCongress will start!"}

func main() {
	var (
		apiToken   = flag.String("apitoken", "", "Telegram API Token")
		eventsFile = flag.String("events", "", "Path to events.json")
	)
	flag.Parse()

	if len(*apiToken) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	// read file
	data, err := ioutil.ReadFile(*eventsFile)
	if err != nil {
		log.Fatalf("Could not read events file %s", err.Error())
	}
	var events []*event.Event
	err = json.Unmarshal(data, &events)
	if err != nil {
		log.Fatalf("Could not unmarshal events file %s", err.Error())
	}

	sort.Sort(event.ByBegin(events))

	b := telegram.NewBot(
		telegram.APIToken(*apiToken),
		telegram.Commands(
			command.NewEvents(events),
			command.NewDays(events),
			command.NewSeconds(events),
			command.NewSeconds(events),
		),
	)
	b.Start()
}
