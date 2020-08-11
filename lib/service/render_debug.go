package service

import (
	"fmt"
	"github.com/DemonTPx/go-game/lib/common"
	"github.com/DemonTPx/go-game/lib/render"
	gl "github.com/chsc/gogl/gl21"
	"github.com/veandco/go-sdl2/sdl"
)

type DebugRenderer struct {
	font  *render.Font
	color common.Color

	tickStart uint32
	counter   uint32

	texture *render.Texture
}

func NewDebugRenderer(font *render.Font, color common.Color) *DebugRenderer {
	return &DebugRenderer{
		font:  font,
		color: color,
	}
}

func (r *DebugRenderer) Render(viewport common.Rect) {
	defer func() {
		r.counter++
	}()

	tick := sdl.GetTicks()
	if tick > r.tickStart+1000 {
		r.renderNewTexture()
		r.counter = 0
		r.tickStart = tick
	}

	if r.texture != nil {
		r.texture.Draw(gl.Float(viewport.X2()-100), gl.Float(viewport.Y+5), 1)
	}
}

func (r *DebugRenderer) renderNewTexture() {
	texture, err := r.font.RenderTexture(fmt.Sprintf("FPS: %d", r.counter), &r.color)
	if err != nil {
		fmt.Printf("error while rendering text: %s", err)
		return
	}

	if r.texture != nil {
		r.texture.Destroy()
	}
	r.texture = texture
}
