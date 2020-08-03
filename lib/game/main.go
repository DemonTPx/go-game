package game

import (
	"fmt"
	"github.com/DemonTPx/go-game/lib/actor"
	"github.com/DemonTPx/go-game/lib/render"
	gl "github.com/chsc/gogl/gl21"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
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

	err = ttf.Init()
	if err != nil {
		return fmt.Errorf("error while initializing sdl2_ttf: %s", err)
	}
	defer ttf.Quit()

	sdl.GLSetAttribute(sdl.GL_CONTEXT_MAJOR_VERSION, 3)
	sdl.GLSetAttribute(sdl.GL_CONTEXT_MINOR_VERSION, 3)
	sdl.GLSetAttribute(sdl.GL_CONTEXT_PROFILE_MASK, sdl.GL_CONTEXT_PROFILE_CORE)

	sdl.GLSetAttribute(sdl.GL_MULTISAMPLESAMPLES, 4)
	sdl.GLSetAttribute(sdl.GL_DOUBLEBUFFER, 1)

	m.Window, m.Renderer, err = sdl.CreateWindowAndRenderer(windowW, windowH, sdl.WINDOW_OPENGL|sdl.WINDOW_RESIZABLE)
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

	err = gl.Init()
	if err != nil {
		return fmt.Errorf("failed to initialize opengl")
	}

	gl.ClearColor(0.2, 0.2, 0.2, 1.0)
	gl.Viewport(0, 0, gl.Sizei(windowW), gl.Sizei(windowH))

	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()

	gl.Ortho(gl.Double(0), gl.Double(windowW), gl.Double(windowH), gl.Double(0), gl.Double(1.0), gl.Double(-1.0))

	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()

	gl.Enable(gl.TEXTURE_2D)

	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	actorFilenameList := []string{
		"res/actor/paddle.yml",
		"res/actor/ball.yml",
		"res/actor/wall-top.yml",
		"res/actor/wall-bottom.yml",
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
	font, err := render.NewFont("res/font/Inconsolata-Regular.ttf", 20)
	if err != nil {
		return fmt.Errorf("error while opening font: %s", err)
	}
	defer font.Close()

	smiling, err := render.NewTextureFromFile("res/sprite/awesomeface.png")
	if err != nil {
		return fmt.Errorf("error while loading texture: %s", err)
	}

	text, err := font.RenderTexture("Hallo, dit is wat tekst!")
	if err != nil {
		return fmt.Errorf("error while rendering text: %s", err)
	}

	m.Timer.Start()
	delta := 1 * time.Second / 60
	for m.Running {
		m.handleEvents()

		for _, a := range m.Actors {
			c := a.GetComponent(actor.Control)
			if c != nil {
				c.Update(delta)
			}
		}
		for _, a := range m.Actors {
			c := a.GetComponent(actor.Physics)
			if c != nil {
				c.Update(delta)
			}
		}

		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		gl.LoadIdentity()

		smiling.Draw(100, 100, -0.8)
		text.Draw(100, 700, 0.1)

		for _, a := range m.Actors {
			c := a.GetComponent(actor.Render)
			if c != nil {
				c.Update(delta)
				c.(actor.Renderer).Render()
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
