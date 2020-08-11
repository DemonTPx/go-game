package event

import "fmt"

type Name string

type Event interface {
	Name() Name
}

type ListenerName string

type Listener interface {
	Name() ListenerName
	OnEvent(Event)
}

type Dispatcher struct {
	listeners map[Name]map[ListenerName]Listener
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		listeners: map[Name]map[ListenerName]Listener{},
	}
}

func (d *Dispatcher) Dispatch(event Event) {
	listeners, ok := d.listeners[event.Name()]
	if !ok {
		return
	}
	for _, l := range listeners {
		l.OnEvent(event)
	}
}

func (d *Dispatcher) AddListener(listener Listener, event Name) {
	listeners, ok := d.listeners[event]
	if !ok {
		d.listeners[event] = map[ListenerName]Listener{}
	}
	if _, ok := listeners[listener.Name()]; ok {
		fmt.Printf("listener %s was already registered for event %s", listener.Name(), event)
		return
	}
	d.listeners[event][listener.Name()] = listener
}

func (d *Dispatcher) RemoveListener(listener Listener, event Name) {
	listeners, ok := d.listeners[event]
	if !ok {
		return
	}
	if _, ok := listeners[listener.Name()]; !ok {
		fmt.Printf("listener %+v was not registered for event %s", listener, event)
		return
	}
	delete(listeners, listener.Name())
}

type SimpleListener struct {
	name    ListenerName
	onEvent func(Event)
}

func NewSimpleListener(name ListenerName, onEvent func(Event)) *SimpleListener {
	return &SimpleListener{
		name:    name,
		onEvent: onEvent,
	}
}

func (l *SimpleListener) Name() ListenerName {
	return l.name
}

func (l *SimpleListener) OnEvent(event Event) {
	l.onEvent(event)
}
