package actor

import (
	"strings"
	"time"
)

type Id uint64

type Actor struct {
	id         Id
	components []Component
}

func NewActor(id Id) *Actor {
	return &Actor{
		id: id,
	}
}

func (a *Actor) AddComponent(c Component) {
	a.components = append(a.components, c)
}

func (a *Actor) String() string {
	return "<Actor " + a.ListComponents() + ">"
}

func (a *Actor) ListComponents() string {
	var list []string
	for _, c := range a.components {
		list = append(list, c.String())
	}

	return strings.Join(list, ", ")
}

func (a *Actor) Update(delta time.Duration) {
	for _, c := range a.components {
		c.Update(delta)
	}
}
