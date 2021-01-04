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
		),
	)
	b.Start()
}
