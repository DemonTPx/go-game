package actor

import (
	"github.com/DemonTPx/go-game/lib/event"
)

type NewEvent struct {
	Id    Id
	Actor *Actor
}

func (e *NewEvent) Name() event.Name {
	return "actor.NewEvent"
}

func NewNewEvent(id Id, actor *Actor) *NewEvent {
	return &NewEvent{Id: id, Actor: actor}
}

type DestroyedEvent struct {
	Id Id
}

func (e *DestroyedEvent) Name() event.Name {
	return "actor.DestroyedEvent"
}

func NewDestroyedEvent(id Id) *DestroyedEvent {
	return &DestroyedEvent{Id: id}
}
