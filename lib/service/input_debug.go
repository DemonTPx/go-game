package service

import (
	"fmt"
	"github.com/DemonTPx/go-game/lib/actor"
	"github.com/veandco/go-sdl2/sdl"
)

type DumpActorsInputHandler struct {
	actorCollection *actor.Collection
}

func NewDumpActorsInputHandler(actorCollection *actor.Collection) *DumpActorsInputHandler {
	return &DumpActorsInputHandler{actorCollection: actorCollection}
}

func (h *DumpActorsInputHandler) HandleEvent(e sdl.Event) {
	switch e.(type) {
	case *sdl.KeyboardEvent:
		evt := e.(*sdl.KeyboardEvent)
		if evt.Type == sdl.KEYDOWN && evt.Keysym.Sym == sdl.K_F12 {
			fmt.Println("listing all actors")
			h.actorCollection.Dump()
		}
	}
}
