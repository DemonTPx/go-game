package actor

import (
	"time"
)

type Component interface {
	Name() string
	String() string
	Owner() *Actor
	Update(delta time.Duration)
}

type BaseComponent struct {
	owner *Actor
}

func (c *BaseComponent) Owner() *Actor {
	return c.owner
}

func (c *BaseComponent) Update(delta time.Duration) {
}

type RenderComponent struct {
	BaseComponent
}

func NewRenderComponent() *RenderComponent {
	return &RenderComponent{}
}

func (c *RenderComponent) Name() string {
	return "RenderComponent"
}

func (c *RenderComponent) String() string {
	return "<" + c.Name() + ">"
}

type BallRenderComponent struct {
	RenderComponent
	color Color
}

func NewBallRenderComponent(color Color) *BallRenderComponent {
	return &BallRenderComponent{color: color}
}

func (c *BallRenderComponent) Name() string {
	return "BallRenderComponent"
}

func (c *BallRenderComponent) String() string {
	return "<" + c.Name() + " color=" + c.color.String() + ">"
}

type PhysicsComponent struct {
	BaseComponent
}

func NewPhysicsComponent() *PhysicsComponent {
	return &PhysicsComponent{}
}

func (c *PhysicsComponent) Name() string {
	return "PhysicsComponent"
}

func (c *PhysicsComponent) String() string {
	return "<" + c.Name() + ">"
}
