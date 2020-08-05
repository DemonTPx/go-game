package game

import (
	"fmt"
	"github.com/DemonTPx/go-game/lib/actor"
	"github.com/DemonTPx/go-game/lib/common"
	"github.com/DemonTPx/go-game/lib/render"
	gl "github.com/chsc/gogl/gl21"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"time"
)

type Main struct {
	Window   *sdl.Window
	Renderer *sdl.Renderer
	Context  sdl.GLContext

	Running bool

	Frame      uint64
	FrameTimer *Timer
	FPS        *FramesPerSecond

	ActorLoader  *actor.Loader
	Actors       map[actor.Id]*actor.Actor
	ActorWatcher *actor.Watcher
}

const (
	windowTitle = "Go Game"

	windowW = 1600
	windowH = 900

	fps           = 60
	frameDuration = time.Second / fps

	enableActorWatcher = true
)

func NewMain() *Main {
	return &Main{
		Running:     true,
		FrameTimer:  NewTimer(),
		FPS:         NewFramesPerSecond(5 * time.Second),
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

	_ = sdl.GLSetAttribute(sdl.GL_CONTEXT_MAJOR_VERSION, 3)
	_ = sdl.GLSetAttribute(sdl.GL_CONTEXT_MINOR_VERSION, 3)
	_ = sdl.GLSetAttribute(sdl.GL_CONTEXT_PROFILE_MASK, sdl.GL_CONTEXT_PROFILE_CORE)

	_ = sdl.GLSetAttribute(sdl.GL_MULTISAMPLESAMPLES, 4)
	_ = sdl.GLSetAttribute(sdl.GL_DOUBLEBUFFER, 1)

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

	gl.Ortho(gl.Double(0), gl.Double(windowW), gl.Double(windowH), gl.Double(0), gl.Double(-1.0), gl.Double(1.0))

	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()

	gl.Enable(gl.TEXTURE_2D)
	gl.Enable(gl.DEPTH_TEST)

	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	if enableActorWatcher {
		m.ActorWatcher, err = actor.NewWatcher()
		if err != nil {
			return fmt.Errorf("failed to initialize actor watcher: %s", err)
		}
	}

	actorFilenameList := []string{
		"res/actor/paddle.yml",
		"res/actor/ball.yml",
		"res/actor/wall-top.yml",
		"res/actor/wall-bottom.yml",
	}

	for _, filename := range actorFilenameList {
		var id actor.Id
		id, err = m.loadActor(filename)
		if err != nil {
			return err
		}

		err = m.ActorWatcher.Watch(id, filename)
		if err != nil {
			return err
		}
	}

	if m.ActorWatcher != nil {
		m.ActorWatcher.Start()
	}

	return m.mainLoop()
}

func (m *Main) loadActor(filename string) (actor.Id, error) {
	a, err := m.ActorLoader.LoadActorFromFile(filename)
	if err != nil {
		return actor.InvalidId, err
	}

	m.Actors[a.Id()] = a

	return a.Id(), nil
}

func (m *Main) reloadChangedActors() error {
	if m.ActorWatcher == nil {
		return nil
	}

	for _, w := range m.ActorWatcher.GetChangedActors() {
		oldActor, ok := m.Actors[w.Id]
		if !ok {
			fmt.Printf("skip reloading actor since old actor was not found\n")
			continue
		}

		id, err := m.loadActor(w.Filename)
		if err != nil {
			fmt.Printf("error while reloading actor: %s\n", err)
			continue
		}

		oldActor.Destroy()
		delete(m.Actors, w.Id)

		fmt.Printf("replaced actor %d with %d\n", w.Id, id)

		err = m.ActorWatcher.Unwatch(w.Filename)
		if err != nil {
			return err
		}
		err = m.ActorWatcher.Watch(id, w.Filename)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *Main) mainLoop() error {
	font, err := render.NewFont("res/font/Inconsolata-Regular.ttf", 72)
	if err != nil {
		return fmt.Errorf("error while opening font: %s", err)
	}
	defer font.Close()

	text, err := font.RenderTextureShadow(
		"Hallo, dit is wat tekst!",
		common.NewColor(1.0, 0.5, 0, 1),
		common.NewColor(0.3, 0.3, 0.3, 0.3),
		common.NewVector2(2, 2),
	)
	if err != nil {
		return fmt.Errorf("error while rendering text: %s", err)
	}

	m.FrameTimer.Reset()
	m.FPS.Reset()
	for m.Running {
		delta := m.FrameTimer.Duration()
		m.FrameTimer.Reset()

		err = m.reloadChangedActors()
		if err != nil {
			return err
		}

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

		text.Draw(100, 300, -0.5)

		for _, a := range m.Actors {
			c := a.GetComponent(actor.Render)
			if c != nil {
				c.Update(delta)
				c.(actor.Renderer).Render()
			}
		}

		m.flip()
	}

	for _, a := range m.Actors {
		a.Destroy()
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
			if k.Type == sdl.KEYDOWN {
				switch k.Keysym.Sym {
				case sdl.K_ESCAPE:
					m.Running = false
				case sdl.K_F12:
					fmt.Println("listing all actors")
					for id, a := range m.Actors {
						fmt.Printf("actor %d: %+v\n\n", id, a)
					}
				}
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

	duration := m.FrameTimer.Duration()
	if duration < frameDuration {
		time.Sleep(frameDuration - duration)
	}

	m.FPS.OnFrame(m.Frame)
}
