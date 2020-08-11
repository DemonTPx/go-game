package actor

import (
	"fmt"
	"github.com/DemonTPx/go-game/lib/event"
)

type Collection struct {
	actors          map[Id]*Actor
	eventDispatcher *event.Dispatcher
}

func NewCollection(dispatcher *event.Dispatcher) *Collection {
	return &Collection{
		actors:          map[Id]*Actor{},
		eventDispatcher: dispatcher,
	}
}

func (c *Collection) Add(id Id, a *Actor) {
	c.actors[id] = a
	c.eventDispatcher.Dispatch(NewNewEvent(id, a))
}

func (c *Collection) Get(id Id) (*Actor, bool) {
	a, ok := c.actors[id]
	return a, ok
}

func (c *Collection) GetAll() map[Id]*Actor {
	list := map[Id]*Actor{}
	for id, a := range c.actors {
		list[id] = a
	}
	return list
}

func (c *Collection) GetAllComponent(cId ComponentId) map[Id]Component {
	components := map[Id]Component{}
	for id, a := range c.actors {
		c := a.GetComponent(cId)
		if c != nil {
			components[id] = c
		}
	}
	return components
}

func (c *Collection) DestroyAndRemove(a *Actor) {
	a.Destroy()
	c.eventDispatcher.Dispatch(NewDestroyedEvent(a.Id()))
	delete(c.actors, a.id)
}

func (c *Collection) Destroy() {
	for _, a := range c.actors {
		c.DestroyAndRemove(a)
	}
}

func (c *Collection) Dump() {
	for id, a := range c.actors {
		fmt.Printf("actor %d: %+v\n\n", id, a)
	}
}
