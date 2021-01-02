package time

import (
	"testing"
	"time"
)

type data struct {
	a, b    time.Time
	days    int
	seconds int64
}

func TestDaysAndSecondsBetween(t *testing.T) {
	now := time.Now()
	table := []data{
		{
			now,
			now,
			0,
			0,
		},
		{
			time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC),
			time.Date(2017, time.January, 1, 0, 0, 0, 0, time.UTC),
			366,
			31622400,
		},
		{
			time.Date(2017, time.January, 1, 0, 0, 0, 0, time.UTC),
			time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC),
			366,
			31622400,
		},
	}

	for _, test := range table {
		days := DaysBetween(test.a, test.b)
		if days != test.days {
			t.Errorf("DaysBetween returned %d but expected was %d.", days, test.days)
		}
		seconds := SecondsBetween(test.a, test.b)
		if seconds != test.seconds {
			t.Errorf("SecondsBetween returned %d but expected was %d.", seconds, test.seconds)
		}
	}
}

func TestDaysAndSecondsFromNow(t *testing.T) {
	now := time.Now()
	then := time.Now().Add(time.Hour * 1337)
	daysExpected := DaysBetween(now, then)
	secondsExpected := SecondsBetween(now, then)

	_now = func() time.Time { return now }
	days := DaysFromNow(then)
	if days != daysExpected {
		t.Errorf("DaysFromNow returned %d but expected was %d.", days, daysExpected)
	}

	seconds := SecondsFromNow(then)
	if seconds != secondsExpected {
		t.Errorf("SecondsFromNow returned %d but expected was %d.", seconds, secondsExpected)
	}
}
