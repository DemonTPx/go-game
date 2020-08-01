package actor

import (
	"strings"
	"time"
)

type Id uint64

type Actor struct {
	id         Id
	components map[ComponentId]Component
}

func NewActor(id Id) *Actor {
	return &Actor{
		id:         id,
		components: make(map[ComponentId]Component),
	}
}

func (a *Actor) AddComponent(c Component) {
	a.components[c.Id()] = c
}

func (a *Actor) GetComponent(id ComponentId) Component {
	c, ok := a.components[id]
	if !ok {
		return nil
	}

	return c
}

func (a *Actor) Id() Id {
	return a.id
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
