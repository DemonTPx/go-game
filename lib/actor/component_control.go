package actor

import (
	"github.com/veandco/go-sdl2/sdl"
	"time"
)

type ControlComponent struct {
	BaseComponent
	Up    bool
	Down  bool
	Left  bool
	Right bool
	Space bool
}

type Controller interface {
	HandleEvent(sdl.Event)
}

func NewControlComponent() *ControlComponent {
	return &ControlComponent{}
}

func (c *ControlComponent) Id() ComponentId {
	return Control
}

func (c *ControlComponent) Name() string {
	return "ControlComponent"
}

func (c *ControlComponent) String() string {
	return "<" + c.Name() + ">"
}

func (c *ControlComponent) Update(delta time.Duration) {
	physics := c.owner.GetComponent(Physics)
	if physics == nil {
		return
	}

	p := physics.(*PhysicsComponent)

	d := float64(delta/time.Millisecond) / 20

	if c.Up {
		p.Velocity.Y -= d
	}
	if c.Down {
		p.Velocity.Y += d
	}
	if c.Left {
		p.Velocity.X -= d
	}
	if c.Right {
		p.Velocity.X += d
	}
	if c.Space {
		p.Velocity.X = 0
		p.Velocity.Y = 0
	}
}

func (c *ControlComponent) HandleEvent(e sdl.Event) {
	switch e.(type) {
	case *sdl.KeyboardEvent:
		event := e.(*sdl.KeyboardEvent)
		if event.Type == sdl.KEYDOWN {
			switch event.Keysym.Sym {
			case sdl.K_UP:
				c.Up = true
			case sdl.K_DOWN:
				c.Down = true
			case sdl.K_LEFT:
				c.Left = true
			case sdl.K_RIGHT:
				c.Right = true
			case sdl.K_SPACE:
				c.Space = true
			}
		}
		if event.Type == sdl.KEYUP {
			switch event.Keysym.Sym {
			case sdl.K_UP:
				c.Up = false
			case sdl.K_DOWN:
				c.Down = false
			case sdl.K_LEFT:
				c.Left = false
			case sdl.K_RIGHT:
				c.Right = false
			case sdl.K_SPACE:
				c.Space = false
			}
		}
	}
}
