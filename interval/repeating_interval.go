package interval

import (
	"fmt"
	"time"
)

// Repeating describes an interval with recurring events distributed evenly by the duration of the interval.
// The number of Repetitions determine the bounds of the repeating interval (from StartsAt).
// When Repetitions is unset, then the repeating interval will be unbounded and recur infinitely long into the future.
type Repeating struct {
	StartInterval interval
	Repetitions   *uint
	lastGenerated *interval // Added field to keep track of the last generated interval.
}

func RepeatingFromDuration(
	start time.Time,
	repeatEvery time.Duration,
	repetitions *uint,
) (*Repeating, error) {
	in, err := FromDuration(start, repeatEvery)

	if err != nil {
		return nil, err
	}

	return &Repeating{
		StartInterval: *in,
		Repetitions:   repetitions,
	}, nil
}

func (r *Repeating) String() string {
	if r.Repetitions != nil {
		return fmt.Sprintf("%v, reps: %v, times: %v", r.StartInterval, r.RepeatEvery(), *r.Repetitions)
	}
	return fmt.Sprintf("%v, reps: %v", r.StartInterval, r.RepeatEvery())
}

// RepeatEvery returns duration of each repetition. This is identical to the duration of the interval.
func (r *Repeating) RepeatEvery() time.Duration {
	return r.StartInterval.Duration()
}

// Start returns the time the interval begins.
// If "Repetitions" is nil, then this indicates the repeating interval is unbounded
// and as a result Start() will return nil.
func (r *Repeating) Start() time.Time {
	return r.StartInterval.Start()
}

// End returns the time the interval ends.
// If "Repetitions" is nil, then this indicates the repeating interval is unbounded
// and as a result End() will return nil.
func (r *Repeating) End() *time.Time {
	if r.Repetitions == nil {
		return nil
	}

	end := r.StartInterval.start.Add(*r.Duration())
	return &end
}

// Duration returns the duration the repeating interval will be active for or nil if it is unbounded.
func (r *Repeating) Duration() *time.Duration {
	if r.Repetitions == nil {
		return nil
	}

	duration := time.Duration(*r.Repetitions) * r.RepeatEvery()
	return &duration
}

// Started returns a boolean indicating if the interval has begun at the given time.
func (r Repeating) Started(t time.Time) bool {
	start := r.Start()
	return t.Equal(start) || t.After(start)
}

// Ended returns a boolean indicating if the interval has ended at the given time.
// When the repeating interval is unbounded, then this function will always return false.
func (r Repeating) Ended(t time.Time) bool {
	endsAt := r.End()

	if endsAt == nil {
		return false
	}

	return t.After(*endsAt)
}

// Next returns the time of the next interval-occurrence relative to the given time.
// It returns the start time if the interval have not started yet and nil if the interval has ended.
func (r Repeating) Next(t time.Time) *time.Time {
	if !r.Started(t) {
		start := r.Start()
		return &start
	}

	if r.Ended(t) || r.RepeatEvery() == 0 {
		return nil
	}

	diff := t.Sub(r.Start())
	mod := diff % r.RepeatEvery()
	next := t.Add(r.RepeatEvery() - mod)

	if r.Ended(next) {
		return nil
	}

	return &next
}

// NextInterval returns the interval instance for the next occurrence relative to the given time.
// It returns nil if the interval has ended, if intervals repeat every zero duration, or if the next interval is invalid.
func (r Repeating) NextInterval(t time.Time) *interval {
	if !r.Started(t) {
		start := r.Start()
		nextInterval, err := FromDuration(start, r.RepeatEvery())
		if err != nil {
			return nil // Handle error or return nil if interval is invalid.
		}
		return nextInterval
	}

	if r.Ended(t) || r.RepeatEvery() == 0 {
		return nil
	}

	diff := t.Sub(r.Start())
	fullIntervalsPassed := diff / r.RepeatEvery()
	nextStart := r.Start().Add((fullIntervalsPassed + 1) * r.RepeatEvery())

	nextInterval, err := FromDuration(nextStart, r.RepeatEvery())

	if err != nil || r.Ended(nextInterval.end) {
		return nil // Handle error or check if the interval has ended.
	}

	return nextInterval
}

// NextGeneratedInterval computes and returns the next interval based on the last generated interval.
// If no interval has been generated yet, it computes the first interval starting from the start of the repeating interval.
// It updates the lastGenerated field to keep track of the last interval that was generated.
func (r *Repeating) NextGeneratedInterval() *interval {
	nextInterval := interval{}

	if r.lastGenerated == nil {
		// If no interval has been generated yet, start from the beginning.
		nextInterval = r.StartInterval
	} else {
		in, err := FromDuration(r.lastGenerated.end, r.RepeatEvery())
		if err != nil {
			return nil // Handle or log error.
		}

		nextInterval = *in
	}

	// Check if the newly generated interval exceeds the number of repetitions.
	if r.Repetitions != nil {
		end := *r.End()

		if nextInterval.end.After(end) {
			return nil // If the next interval goes beyond the total duration, return nil.
		}
	}

	// Update the lastGenerated field with the newly generated interval.
	r.lastGenerated = &nextInterval
	return r.lastGenerated
}
