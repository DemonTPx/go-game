package game

import (
	"time"
)

type Timer struct {
	start time.Time
}

func NewTimer() *Timer {
	return &Timer{
		start: time.Now(),
	}
}

func (t *Timer) Reset() {
	t.start = time.Now()
}

func (t *Timer) Duration() time.Duration {
	return time.Now().Sub(t.start)
}
