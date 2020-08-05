package game

import (
	"fmt"
	"time"
)

type FramesPerSecond struct {
	lastFrame uint64
	timer     *Timer
	interval  time.Duration
}

func NewFramesPerSecond(interval time.Duration) *FramesPerSecond {
	return &FramesPerSecond{
		timer:    NewTimer(),
		interval: interval,
	}
}

func (t *FramesPerSecond) Reset() {
	t.timer.Reset()
}

func (t *FramesPerSecond) OnFrame(frame uint64) {
	if t.timer.Duration() < t.interval {
		return
	}
	t.timer.Reset()

	fps := float64(frame-t.lastFrame) / float64(t.interval/time.Second)
	fmt.Printf("frame: %d, fps: %f\n", frame, fps)

	t.lastFrame = frame
}
