package game

import (
	"fmt"
	"github.com/DemonTPx/go-game/lib/actor"
	"github.com/DemonTPx/go-game/lib/common"
	"github.com/DemonTPx/go-game/lib/event"
	"github.com/DemonTPx/go-game/lib/render"
	"github.com/DemonTPx/go-game/lib/service"
	gl "github.com/chsc/gogl/gl21"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"time"
)

const (
	windowTitle = "Go Game"

	windowW = 1600
	windowH = 900

	frameMinSleep = 10 * time.Millisecond

	enableActorWatcher = true
)

type Main struct {
	Window   *sdl.Window
	Renderer *sdl.Renderer
	Context  sdl.GLContext

	Running bool

	Viewport common.Rect

	Frame      uint64
	FrameTimer *Timer

	EventDispatcher *event.Dispatcher
	InputManager    *service.InputManager
	RenderManager   *service.RenderManager

	ActorLoader     *actor.Loader
	ActorCollection *actor.Collection
	ActorWatcher    *actor.Watcher

	Font *render.Font
}

func NewMain() *Main {
	eventDispatcher := event.NewDispatcher()

	return &Main{
		Running:    true,
		FrameTimer: NewTimer(),

		Viewport: common.NewRect(0, 0, windowW, windowH),

		EventDispatcher: eventDispatcher,
		InputManager:    service.NewInputManager(eventDispatcher),
		RenderManager:   service.NewRenderManager(eventDispatcher),

		ActorLoader:     actor.NewLoader(),
		ActorCollection: actor.NewCollection(eventDispatcher),
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

	err = sdl.GLSetSwapInterval(-1)
	if err != nil {
		fmt.Println("failed to enable adaptive vsync, falling back to vsync")
		err = sdl.GLSetSwapInterval(1)
		if err != nil {
			return fmt.Errorf("failed to enable vsync")
		}
	}

	gl.ClearColor(0.2, 0.2, 0.2, 1.0)
	gl.Viewport(0, 0, gl.Sizei(windowW), gl.Sizei(windowH))

	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()

	gl.Ortho(gl.Double(0), gl.Double(windowW), gl.Double(windowH), gl.Double(0), gl.Double(-1.0), gl.Double(1.0))

	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()

	gl.Enable(gl.TEXTURE_2D)

	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	m.Font, err = render.NewFont("res/font/Inconsolata-Regular.ttf", 12)
	if err != nil {
		return fmt.Errorf("unable to load font")
	}

	m.RenderManager.Add(service.NewDebugRenderer(m.Font, common.NewColorWhite()))

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

		if m.ActorWatcher != nil {
			err = m.ActorWatcher.Watch(id, filename)
			if err != nil {
				return err
			}
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

	m.ActorCollection.Add(a.Id(), a)

	return a.Id(), nil
}

func (m *Main) reloadChangedActors() error {
	if m.ActorWatcher == nil {
		return nil
	}

	for _, w := range m.ActorWatcher.GetChangedActors() {
		oldActor, ok := m.ActorCollection.Get(w.Id)
		if !ok {
			fmt.Printf("skip reloading actor since old actor was not found\n")
			continue
		}

		id, err := m.loadActor(w.Filename)
		if err != nil {
			fmt.Printf("error while reloading actor: %s\n", err)
			continue
		}

		m.ActorCollection.DestroyAndRemove(oldActor)

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
	var err error

	m.FrameTimer.Reset()
	for m.Running {
		delta := m.FrameTimer.Duration()
		m.FrameTimer.Reset()

		err = m.reloadChangedActors()
		if err != nil {
			return err
		}

		m.handleEvents()

		for _, c := range m.ActorCollection.GetAllComponent(actor.Control) {
			c.Update(delta)
		}
		for _, c := range m.ActorCollection.GetAllComponent(actor.Physics) {
			c.Update(delta)
		}

		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		gl.LoadIdentity()

		for _, c := range m.ActorCollection.GetAllComponent(actor.Render) {
			c.Update(delta)
		}

		m.RenderManager.Render(m.Viewport)

		m.flip()
	}

	m.ActorCollection.Destroy()
	m.Font.Close()

	return nil
}

func (m *Main) handleEvents() {
	for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
		switch e.(type) {
		case *sdl.QuitEvent:
			m.Running = false
			break
		case *sdl.KeyboardEvent:
			k := e.(*sdl.KeyboardEvent)
			if k.Type == sdl.KEYDOWN {
				switch k.Keysym.Sym {
				case sdl.K_ESCAPE:
					m.Running = false
				case sdl.K_F12:
					fmt.Println("listing all actors")
					m.ActorCollection.Dump()
				}
			}
		}

		m.InputManager.Handle(e)
	}
}

func (m *Main) flip() {
	m.Window.GLSwap()
	m.Frame++

	duration := m.FrameTimer.Duration()
	if duration < frameMinSleep {
		time.Sleep(frameMinSleep - duration)
	}
}
