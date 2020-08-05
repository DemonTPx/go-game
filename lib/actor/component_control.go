package actor

import (
	"fmt"
	"github.com/DemonTPx/go-game/lib/common"
	"github.com/DemonTPx/go-game/lib/input"
	"github.com/veandco/go-sdl2/sdl"
	"time"
)

const (
	ActionUp    input.Action = 1
	ActionDown  input.Action = 2
	ActionLeft  input.Action = 3
	ActionRight input.Action = 4
	ActionStop  input.Action = 5
)

type Controller interface {
	HandleEvent(sdl.Event)
}

type ControlComponent struct {
	BaseComponent
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
}

func (c *ControlComponent) HandleEvent(e sdl.Event) {
}

type FreeControlComponent struct {
	ControlComponent
	handler        *input.ActionHandler
	VelocityFactor float64
	VelocityMax    float64
}

func NewFreeControlComponent(velocityFactor, velocityMax float64) *FreeControlComponent {
	return &FreeControlComponent{
		VelocityFactor: velocityFactor,
		VelocityMax:    velocityMax,
		handler: input.NewActionHandler(
			map[sdl.Keycode]input.Action{
				sdl.K_UP:    ActionUp,
				sdl.K_DOWN:  ActionDown,
				sdl.K_LEFT:  ActionLeft,
				sdl.K_RIGHT: ActionRight,
				sdl.K_SPACE: ActionStop,
			},
		),
	}
}

func (c *FreeControlComponent) Name() string {
	return "FreeControlComponent"
}

func (c *FreeControlComponent) String() string {
	return fmt.Sprintf("<%s velocity_factor=%f velocity_max=%f>", c.Name(), c.VelocityFactor, c.VelocityMax)
}

func (c *FreeControlComponent) Update(delta time.Duration) {
	physics := c.owner.GetComponent(Physics)
	if physics == nil {
		return
	}

	p := physics.(*PhysicsComponent)

	deltaMs := float64(delta / time.Millisecond)
	velocityChange := deltaMs * c.VelocityFactor

	for _, a := range c.handler.ActiveActions() {
		switch a {
		case ActionUp:
			p.Velocity.Y -= velocityChange
		case ActionDown:
			p.Velocity.Y += velocityChange
		case ActionLeft:
			p.Velocity.X -= velocityChange
		case ActionRight:
			p.Velocity.X += velocityChange
		case ActionStop:
			p.Velocity.X = 0
			p.Velocity.Y = 0
		}
	}

	p.Velocity.X = common.Clamp(p.Velocity.X, -c.VelocityMax, c.VelocityMax)
	p.Velocity.Y = common.Clamp(p.Velocity.Y, -c.VelocityMax, c.VelocityMax)
}

func (c *FreeControlComponent) HandleEvent(e sdl.Event) {
	c.handler.HandleEvent(e)
}

type PaddleControlComponent struct {
	ControlComponent
	handler        *input.ActionHandler
	VelocityFactor float64
	VelocityMax    float64
}

func NewPaddleControlComponent(velocityFactor, velocityMax float64) *PaddleControlComponent {
	return &PaddleControlComponent{
		VelocityFactor: velocityFactor,
		VelocityMax:    velocityMax,
		handler: input.NewActionHandler(
			map[sdl.Keycode]input.Action{
				sdl.K_UP:   ActionUp,
				sdl.K_DOWN: ActionDown,
			},
		),
	}
}

func (c *PaddleControlComponent) Name() string {
	return "PaddleControlComponent"
}

func (c *PaddleControlComponent) String() string {
	return fmt.Sprintf("<%s velocity_factor=%f velocity_max=%f>", c.Name(), c.VelocityFactor, c.VelocityMax)
}

func (c *PaddleControlComponent) Update(delta time.Duration) {
	physics := c.owner.GetComponent(Physics)
	if physics == nil {
		return
	}

	p := physics.(*PhysicsComponent)
	d := float64(delta/time.Millisecond) * c.VelocityFactor

	for _, a := range c.handler.ActiveActions() {
		switch a {
		case ActionUp:
			p.Velocity.Y -= d
		case ActionDown:
			p.Velocity.Y += d
		}
	}

	p.Velocity.Y = common.Clamp(p.Velocity.Y, -c.VelocityMax, c.VelocityMax)
}

func (c *PaddleControlComponent) HandleEvent(e sdl.Event) {
	c.handler.HandleEvent(e)
}
