package game

import (
	"fmt"
	"github.com/DemonTPx/go-game/lib/actor"
	gl "github.com/chsc/gogl/gl21"
	"github.com/veandco/go-sdl2/sdl"
	"time"
)

type Main struct {
	Window      *sdl.Window
	Renderer    *sdl.Renderer
	Context     sdl.GLContext
	Running     bool
	Frame       uint64
	Timer       *Timer
	ActorLoader *actor.Loader
	Actors      map[actor.Id]*actor.Actor
}

const (
	windowTitle        = "Go Game"
	windowW            = 1600
	windowH            = 900
	frameDelay  uint32 = 1000 / 60
)

func NewMain() *Main {
	return &Main{
		Running:     true,
		Timer:       NewTimer(),
		ActorLoader: actor.NewLoader(),
		Actors:      make(map[actor.Id]*actor.Actor, 0),
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

	actorFilenameList := []string{
		"res/actor/paddle.yml",
		"res/actor/ball.yml",
	}

	for _, filename := range actorFilenameList {
		err = m.loadActor(filename)
		if err != nil {
			return err
		}
	}

	return m.mainLoop()
}

func (m *Main) loadActor(filename string) error {
	a, err := m.ActorLoader.LoadActorFromFile(filename)
	if err != nil {
		return err
	}

	m.Actors[a.Id()] = a

	return nil
}

func (m *Main) mainLoop() error {
	m.Timer.Start()
	delta := 1 * time.Second / 60
	for m.Running {
		m.handleEvents()

		for _, a := range m.Actors {
			control := a.GetComponent(actor.Control)
			if control != nil {
				control.Update(delta)
			}
		}
		for _, a := range m.Actors {
			physics := a.GetComponent(actor.Physics)
			if physics != nil {
				physics.Update(delta)
			}
		}

		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		for _, a := range m.Actors {
			render := a.GetComponent(actor.Render)
			if render != nil {
				render.Update(delta)
				render.(actor.Renderer).Render()
			}
		}

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
			k := event.(*sdl.KeyboardEvent)
			if k.Type == sdl.KEYDOWN && k.Keysym.Sym == sdl.K_ESCAPE {
				m.Running = false
			}
		}

		for _, a := range m.Actors {
			control := a.GetComponent(actor.Control)
			if control == nil {
				continue
			}

			control.(actor.Controller).HandleEvent(event)
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
