package game

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Main struct {
	Window  *sdl.Window
	Running bool
	Frame   uint64
	Timer   *Timer
	aLeft   bool
	aRight  bool
	aUp     bool
	aDown   bool
	vX      int32
	vY      int32
}

const (
	windowTitle        = "Testing"
	windowW            = 800
	windowH            = 600
	frameDelay  uint32 = 1000 / 60
)

func NewMain() *Main {
	return &Main{
		Running: true,
		Timer:   NewTimer(),
	}
}

func (m *Main) Run() error {
	var err error

	err = sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		return err
	}
	defer sdl.Quit()

	m.Window, err = sdl.CreateWindow(
		windowTitle,
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		windowW,
		windowH,
		sdl.WINDOW_SHOWN,
	)
	if err != nil {
		return err
	}
	defer m.Window.Destroy()

	return m.mainLoop()
}

func (m *Main) mainLoop() error {
	x := int32(0)
	y := int32(0)

	m.Timer.Start()
	for m.Running {
		m.handleEvents()

		surface, err := m.Window.GetSurface()
		if err != nil {
			return err
		}

		// Clear screen
		surface.FillRect(nil, 0)

		if m.aLeft {
			m.vX = m.vX - 1
		}
		if m.aRight {
			m.vX = m.vX + 1
		}
		if m.aUp {
			m.vY = m.vY - 1
		}
		if m.aDown {
			m.vY = m.vY + 1
		}

		x += m.vX
		y += m.vY

		if x > (windowW - 20) {
			m.vX = 0
			x = windowW - 20
		}
		if x < 0 {
			m.vX = 0
			x = 0
		}
		if y > (windowH - 20) {
			m.vY = 0
			y = windowH - 20
		}
		if y < 0 {
			m.vY = 0
			y = 0
		}

		rect := sdl.Rect{x, y, 20, 20}
		surface.FillRect(&rect, 0xffff0000)

		m.flip()
	}

	return nil
}

func (m *Main) handleEvents() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			m.Running = false
			break
		case *sdl.KeyboardEvent:
			event := event.(*sdl.KeyboardEvent)
			if event.Type == sdl.KEYDOWN {
				switch event.Keysym.Sym {
				case sdl.K_ESCAPE:
					m.Running = false
				case sdl.K_LEFT:
					m.aLeft = true
				case sdl.K_RIGHT:
					m.aRight = true
				case sdl.K_UP:
					m.aUp = true
				case sdl.K_DOWN:
					m.aDown = true
				}
			}
			if event.Type == sdl.KEYUP {
				switch event.Keysym.Sym {
				case sdl.K_ESCAPE:
					m.Running = false
				case sdl.K_LEFT:
					m.aLeft = false
				case sdl.K_RIGHT:
					m.aRight = false
				case sdl.K_UP:
					m.aUp = false
				case sdl.K_DOWN:
					m.aDown = false
				}
			}
		}
	}
}

func (m *Main) flip() {
	err := m.Window.UpdateSurface()
	if err != nil {
		panic(err)
	}
	m.Frame++

	ticks := m.Timer.GetTicks()
	if ticks < frameDelay {
		delay := frameDelay - ticks
		sdl.Delay(delay)
	}
	m.Timer.Start()
}
