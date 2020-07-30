package actor

import "time"

type ComponentId uint8

const (
	Transform ComponentId = 1
	Render    ComponentId = 2
	Physics   ComponentId = 3
)

type Component interface {
	Id() ComponentId
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

type Builder interface {
	Build(data VariableConfig) (Component, error)
}
