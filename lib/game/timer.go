package game

import "github.com/veandco/go-sdl2/sdl"

type Timer struct {
	started    bool
	startTicks uint32
}

func NewTimer() *Timer {
	return &Timer{}
}

func (t *Timer) Start() {
	t.started = true
	t.startTicks = sdl.GetTicks()
}

func (t *Timer) Stop() {
	t.started = false
}

func (t *Timer) GetTicks() uint32 {
	if t.started {
		return sdl.GetTicks() - t.startTicks
	}

	return 0
}
