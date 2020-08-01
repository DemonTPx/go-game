package game

import (
	"fmt"
	gl "github.com/chsc/gogl/gl21"
	"github.com/veandco/go-sdl2/sdl"
	"math"
)

type Main struct {
	Window   *sdl.Window
	Renderer *sdl.Renderer
	Context  sdl.GLContext
	Running  bool
	Frame    uint64
	Timer    *Timer
	aLeft    bool
	aRight   bool
	aUp      bool
	aDown    bool
	aSpace   bool
	vX       float64
	vY       float64
}

const (
	windowTitle        = "Go Game"
	windowW            = 1600
	windowH            = 900
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

	sdl.GLSetAttribute(sdl.GL_CONTEXT_MAJOR_VERSION, 3)
	sdl.GLSetAttribute(sdl.GL_CONTEXT_PROFILE_MASK, sdl.GL_CONTEXT_PROFILE_CORE)

	sdl.GLSetAttribute(sdl.GL_MULTISAMPLESAMPLES, 4)
	sdl.GLSetAttribute(sdl.GL_DOUBLEBUFFER, 1)

	m.Window, m.Renderer, err = sdl.CreateWindowAndRenderer(windowW, windowH, sdl.WINDOW_OPENGL)
	if err != nil {
		return err
	}
	defer m.Renderer.Destroy()
	defer m.Window.Destroy()

	m.Window.SetTitle(windowTitle)

	info, err := m.Renderer.GetInfo()
	if err != nil {
		return err
	}

	expectedFlags := uint32(sdl.RENDERER_ACCELERATED | sdl.RENDERER_TARGETTEXTURE)
	if (info.Flags & expectedFlags) != expectedFlags {
		return fmt.Errorf("failed to create opengl context")
	}

	m.Context, err = m.Window.GLCreateContext()
	if err != nil {
		return fmt.Errorf("failed to create opengl context")
	}

	gl.Init()
	gl.Viewport(0, 0, gl.Sizei(windowW), gl.Sizei(windowH))

	gl.ClearColor(0.2, 0.2, 0.2, 1.0)
	gl.MatrixMode(gl.PROJECTION)
	gl.Ortho(gl.Double(0), gl.Double(windowW), gl.Double(windowH), gl.Double(0), gl.Double(-1.0), gl.Double(1.0))

	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)

	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	return m.mainLoop()
}

func (m *Main) mainLoop() error {
	x := float64(windowW / 2)
	y := float64(windowH / 2)

	m.Timer.Start()
	for m.Running {
		m.handleEvents()

		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

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
		if m.aSpace {
			m.vX = 0
			m.vY = 0
		}

		x += m.vX
		y += m.vY

		radius := 5 + (math.Abs(m.vX)+math.Abs(m.vY))/2

		if x >= (windowW-radius/2) && m.vX > 0 {
			x = windowW - radius/2
			m.vX = -m.vX
		}
		if x <= radius/2 && m.vX < 0 {
			x = radius / 2
			m.vX = -m.vX
		}
		if y >= (windowH-radius/2) && m.vY > 0 {
			y = windowH - radius/2
			m.vY = -m.vY
		}
		if y <= radius/2 && m.vY < 0 {
			y = radius / 2
			m.vY = -m.vY
		}

		gl.Begin(gl.TRIANGLE_FAN)
		gl.Color4f(gl.Float(1.0), gl.Float(0.0), gl.Float(0.0), gl.Float(1.0))
		gl.Vertex2f(gl.Float(x), gl.Float(y))
		gl.Color4f(gl.Float(1.0), gl.Float(1.0), gl.Float(0.0), gl.Float(1.0))
		segments := float64(20)
		for n := float64(0); n <= segments; n++ {
			t := math.Pi * 2 * n / segments
			gl.Vertex2f(gl.Float(x+math.Sin(t)*radius), gl.Float(y+math.Cos(t)*radius))
		}
		gl.End()

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
				case sdl.K_SPACE:
					m.aSpace = true
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
				case sdl.K_SPACE:
					m.aSpace = false
				}
			}
		}
	}
}

func (m *Main) flip() {
	m.Window.GLSwap()
	m.Frame++

	ticks := m.Timer.GetTicks()
	if ticks < frameDelay {
		delay := frameDelay - ticks
		sdl.Delay(delay)
	}
	m.Timer.Start()
}
