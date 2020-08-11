package service

import (
	"github.com/DemonTPx/go-game/lib/actor"
	"github.com/DemonTPx/go-game/lib/event"
	"github.com/veandco/go-sdl2/sdl"
)

type InputHandlerId uint64

type InputHandler interface {
	HandleEvent(sdl.Event)
}

type InputManager struct {
	eventDispatcher *event.Dispatcher

	handlers      map[InputHandlerId]InputHandler
	lastHandlerId InputHandlerId

	actorHandlerMap map[actor.Id]InputHandlerId
}

func NewInputManager(dispatcher *event.Dispatcher) *InputManager {
	m := &InputManager{
		eventDispatcher: dispatcher,

		handlers:      map[InputHandlerId]InputHandler{},
		lastHandlerId: 0,

		actorHandlerMap: map[actor.Id]InputHandlerId{},
	}

	dispatcher.AddListener(event.NewSimpleListener("service.InputManager.OnActorNew", m.OnActorNew), "actor.NewEvent")
	dispatcher.AddListener(event.NewSimpleListener("service.InputManager.OnActorDestroyed", m.OnActorDestroyed), "actor.DestroyedEvent")

	return m
}

func (m *InputManager) Add(h InputHandler) InputHandlerId {
	id := m.nextHandlerId()
	m.handlers[id] = h
	return id
}

func (m *InputManager) Remove(id InputHandlerId) {
	delete(m.handlers, id)
}

func (m *InputManager) OnActorNew(e event.Event) {
	a := e.(*actor.NewEvent).Actor
	c := a.GetComponent(actor.Control)
	if c == nil {
		return
	}

	handlerId := m.Add(c.(actor.Controller))
	m.actorHandlerMap[a.Id()] = handlerId
}

func (m *InputManager) OnActorDestroyed(e event.Event) {
	actorId := e.(*actor.DestroyedEvent).Id
	handlerId, ok := m.actorHandlerMap[actorId]
	if !ok {
		return
	}
	m.Remove(handlerId)
	delete(m.actorHandlerMap, actorId)
}

func (m *InputManager) nextHandlerId() InputHandlerId {
	m.lastHandlerId++
	return m.lastHandlerId
}

func (m *InputManager) Handle(event sdl.Event) {
	for _, h := range m.handlers {
		h.HandleEvent(event)
	}
}
