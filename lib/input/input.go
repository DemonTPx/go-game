package input

import "github.com/veandco/go-sdl2/sdl"

type Action uint8

type ActionHandler struct {
	Actions  map[Action]bool
	KeyBinds map[sdl.Keycode]Action
}

func NewActionHandler(keyBinds map[sdl.Keycode]Action) *ActionHandler {
	actionsActive := make(map[Action]bool, 0)
	for _, a := range keyBinds {
		actionsActive[a] = false
	}

	return &ActionHandler{
		Actions:  actionsActive,
		KeyBinds: keyBinds,
	}
}

func (h *ActionHandler) HandleEvent(e sdl.Event) {
	switch e.(type) {
	case *sdl.KeyboardEvent:
		event := e.(*sdl.KeyboardEvent)
		if event.Type == sdl.KEYDOWN || event.Type == sdl.KEYUP {
			if action, ok := h.KeyBinds[event.Keysym.Sym]; ok {
				h.Actions[action] = event.Type == sdl.KEYDOWN
			}
		}
	}
}

func (h *ActionHandler) ActiveActions() []Action {
	list := make([]Action, 0)
	for action, active := range h.Actions {
		if active {
			list = append(list, action)
		}
	}

	return list
}
