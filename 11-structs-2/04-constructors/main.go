package main

import (
	"errors"
	"fmt"
	"time"
)

type DateRange struct {
	start time.Time
	end   time.Time
}

func NewDateRange(s, e time.Time) (DateRange, error) {
	if s.IsZero() || e.IsZero() {
		return DateRange{}, errors.New("Date should not be empty.")
	}

	if e.Before(s) {
		return DateRange{}, errors.New("End date should not be earlier than start date.")
	}

	return DateRange{
		start: s,
		end:   e,
	}, nil
}

func (d DateRange) Hours() float64 {
	return d.end.Sub(d.start).Hours()
}

func main() {
	lifetime, _ := NewDateRange(
		time.Date(1815, 12, 10, 0, 0, 0, 0, time.UTC),
		time.Date(1852, 11, 27, 0, 0, 0, 0, time.UTC),
	)

	fmt.Println(lifetime.Hours())

	travelInTime, _ := NewDateRange(
		time.Date(1852, 11, 27, 0, 0, 0, 0, time.UTC),
		time.Date(1815, 12, 10, 0, 0, 0, 0, time.UTC),
	)

	fmt.Println(travelInTime.Hours())
}
