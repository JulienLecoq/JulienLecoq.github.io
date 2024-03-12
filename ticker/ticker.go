package ticker

import (
	"time"
)

// Example:
//
// ticker := ticker.NewTicker(2_000 * time.Millisecond)
//
//	ticker.OnTick(func() {
//		button1Txt := fmt.Sprintf("%s %d", "Button", count)
//		r.Button1.SetText(button1Txt)
//	})
//
// ticker.Start()
type Ticker struct {
	ticker   *time.Ticker
	duration time.Duration
}

func NewTicker(duration time.Duration) *Ticker {
	return &Ticker{
		duration: duration,
	}
}

func (t *Ticker) Start() {
	t.ticker = time.NewTicker(t.duration)
}

func (t *Ticker) OnTick(cb func()) {
	go func() {
		for range t.ticker.C {
			cb()
		}
	}()
}

func (t *Ticker) Stop() {
	if t.ticker == nil {
		return
	}

	t.ticker.Stop()
}

func (t *Ticker) Reset(duration time.Duration) {
	if t.ticker == nil {
		return
	}

	t.ticker.Reset(duration)
}