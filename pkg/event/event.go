package event

import "time"

// Event combines the attributes of an event.
type Event struct {
	Name    string    `json:name`
	Website string    `json:website`
	Begin   time.Time `json:begin`
	End     time.Time `json:end`
}

func (e Event) String() string {
	return e.Name + " " + e.Website + " " + e.Begin.Format("2006-01-02") + " - " + e.End.Format("2006-01-02")
}

type ByBegin []*Event

func (a ByBegin) Len() int           { return len(a) }
func (a ByBegin) Less(i, j int) bool { return a[i].Begin.Before(a[j].Begin) }
func (a ByBegin) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
