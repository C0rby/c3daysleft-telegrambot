package answer

import (
	"fmt"
	"math/rand"
	"strings"
)

// Unit defines an enumeration to distinguish between value units like days or seconds.
type Unit string

const (
	// Days Unit
	Days Unit = "days"
	// Seconds Unit
	Seconds = "seconds"
)

var _formats = [...]string{"%d", "%#x (%d)", "%#o (%d)", "%#b (%d)"}

var _genericAnswers = [...]string{
	"There are {value} {unit} remaining until {event}.",
	"Only {value} {unit} left to hack something great for {event}.",
	"In {value} {unit} you will drink more Mate than usual. Cheers!"}

var _daysAnswers = [...]string{
	"You have to sleep {value} times until {event}.",
	"You have {value} nights to sleep well to be fit for {event}."}

// TODO: Event specific answers need to be added to the events.json
// Like: "In %d %s you'll have a good excuse for leaving christmas with your family!"}

// RandomAnswer returns a prepared answer ready to be send.
func RandomAnswer(value int64, u Unit, event string) string {
	v := fmt.Sprintf(randomFormat(), value, value)
	r := strings.NewReplacer("{value}", v, "{unit}", string(u), "{event}", event)
	return r.Replace(randomAnswer(u))
}

func randomFormat() string {
	return _formats[rand.Intn(len(_formats))]
}

func randomAnswer(u Unit) string {
	a := _genericAnswers[:]
	if u == Days {
		a = append(a, _daysAnswers[:]...)
	}
	return a[rand.Intn(len(a))]
}
