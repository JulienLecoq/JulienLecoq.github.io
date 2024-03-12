package interval

import (
	"errors"
	"time"
)

var invalidInterval = errors.New("invalid interval")

const minDuration = 1

type interval struct {
	start time.Time
	end   time.Time
}

func New(start time.Time, end time.Time) (*interval, error) {
	if end.Before(start) {
		return nil, invalidInterval
	}

	return &interval{
		start: start,
		end:   end,
	}, nil
}

func FromDuration(start time.Time, duration time.Duration) (*interval, error) {
	if duration < minDuration {
		return nil, invalidInterval
	}

	end := start.Add(duration)

	return &interval{
		start: start,
		end:   end,
	}, nil
}

func (i *interval) SetStart(start time.Time) error {
	if start.After(i.end) {
		return invalidInterval
	}

	i.start = start
	return nil
}

func (i *interval) SetEnd(end time.Time) error {
	if end.Before(i.start) {
		return invalidInterval
	}

	i.end = end
	return nil
}

func (i *interval) SetDuration(duration time.Duration) error {
	if duration < minDuration {
		return invalidInterval
	}

	i.end = i.start.Add(duration)

	return nil
}

func (i *interval) Start() time.Time {
	return i.start
}

func (i *interval) End() time.Time {
	return i.end
}

func (i *interval) Duration() time.Duration {
	return i.end.Sub(i.start)
}

func (i interval) String() string {
	dateFormat := "15h04"
	return i.start.Format(dateFormat) + " - " + i.end.Format(dateFormat)
}
