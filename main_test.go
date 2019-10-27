package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDaysBetween(t *testing.T) {
	now := time.Now()
	assert.Equal(t, 0, daysBetween(now, now), "Days between same times was incorrect.")

	day1 := time.Date(2019, time.February, 1, 0, 0, 0, 0, time.UTC)
	day2 := time.Date(2019, time.March, 1, 0, 0, 0, 0, time.UTC)

	assert.Equal(t, 28, daysBetween(day1, day2), "Days between without leap day was incorrect.")
	assert.Equal(t, 28, daysBetween(day2, day1), "Days between without leap day was incorrect.")

	day1LeapYear := time.Date(2020, time.February, 1, 0, 0, 0, 0, time.UTC)
	day2LeapYear := time.Date(2020, time.March, 1, 0, 0, 0, 0, time.UTC)

	assert.Equal(t, 29, daysBetween(day1LeapYear, day2LeapYear), "Days between with leap day was incorrect.")
	assert.Equal(t, 29, daysBetween(day2LeapYear, day1LeapYear), "Days between with leap day was incorrect.")
}

func TestDaysTilCongressFrom(t *testing.T) {
	from := time.Date(2019, time.December, 26, 0, 0, 0, 0, time.UTC)
	assert.Equal(t, 1, daysTilCongressFrom(from), "Days till congress form day bevore was incorrect.")

	fromSameDay := time.Date(2019, time.December, 27, 0, 0, 0, 0, time.UTC)
	assert.Equal(t, 0, daysTilCongressFrom(fromSameDay), "Days till congress from same day was incorrect.")

	fromAfterStart := time.Date(2019, time.December, 28, 0, 0, 0, 0, time.UTC)
	assert.Equal(t, 365, daysTilCongressFrom(fromAfterStart), "Days till congress after in year was incorrect.")
}
