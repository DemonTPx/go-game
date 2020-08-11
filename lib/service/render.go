package service

import (
	"github.com/DemonTPx/go-game/lib/actor"
	"github.com/DemonTPx/go-game/lib/common"
	"github.com/DemonTPx/go-game/lib/event"
)

type RendererId uint64

type Renderer interface {
	Render(viewport common.Rect)
}

type RenderManager struct {
	eventDispatcher *event.Dispatcher

	renderers      map[RendererId]Renderer
	lastRendererId RendererId

	actorRendererMap map[actor.Id]RendererId
}

func NewRenderManager(dispatcher *event.Dispatcher) *RenderManager {
	m := &RenderManager{
		eventDispatcher:  dispatcher,
		renderers:        map[RendererId]Renderer{},
		actorRendererMap: map[actor.Id]RendererId{},
	}

	dispatcher.AddListener(event.NewSimpleListener("service.RenderManager.OnActorNew", m.OnActorNew), "actor.NewEvent")
	dispatcher.AddListener(event.NewSimpleListener("service.RenderManager.OnActorDestroyed", m.OnActorDestroyed), "actor.DestroyedEvent")

	return m
}

func (m *RenderManager) Add(r Renderer) RendererId {
	id := m.nextRendererId()
	m.renderers[id] = r
	return id
}

func (m *RenderManager) Remove(id RendererId) {
	delete(m.renderers, id)
}

func (m *RenderManager) nextRendererId() RendererId {
	m.lastRendererId++
	return m.lastRendererId
}

func (m *RenderManager) Render(viewport common.Rect) {
	for _, r := range m.renderers {
		r.Render(viewport)
	}
}

func (m *RenderManager) OnActorNew(e event.Event) {
	a := e.(*actor.NewEvent).Actor
	c := a.GetComponent(actor.Render)
	if c == nil {
		return
	}

	handlerId := m.Add(c.(actor.Renderer))
	m.actorRendererMap[a.Id()] = handlerId
}

func (m *RenderManager) OnActorDestroyed(e event.Event) {
	actorId := e.(*actor.DestroyedEvent).Id
	handlerId, ok := m.actorRendererMap[actorId]
	if !ok {
		return
	}
	m.Remove(handlerId)
	delete(m.actorRendererMap, actorId)
}
