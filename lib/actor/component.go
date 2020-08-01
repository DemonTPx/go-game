package actor

import "time"

type ComponentId uint8

const (
	Transform ComponentId = 1
	Render    ComponentId = 2
	Physics   ComponentId = 3
	Control   ComponentId = 4
)

type Component interface {
	Id() ComponentId
	Name() string
	String() string
	Owner() *Actor
	SetOwner(actor *Actor)
	Update(delta time.Duration)
}

type BaseComponent struct {
	owner *Actor
}

func (c *BaseComponent) Owner() *Actor {
	return c.owner
}

func (c *BaseComponent) SetOwner(actor *Actor) {
	c.owner = actor
}

func (c *BaseComponent) Update(delta time.Duration) {
}

type Builder interface {
	Build(data VariableConfig) (Component, error)
}
