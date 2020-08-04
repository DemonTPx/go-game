package render

import (
	"github.com/DemonTPx/go-game/lib/common"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"math"
)

type Font struct {
	Font *ttf.Font
	Size int
}

func NewFont(file string, size int) (*Font, error) {
	font, err := ttf.OpenFont(file, size)
	if err != nil {
		return nil, err
	}

	return &Font{
		Font: font,
		Size: size,
	}, nil
}

func (f *Font) RenderTexture(text string, color *common.Color) (*Texture, error) {
	surface, err := f.Font.RenderUTF8Blended(text, color.ToSDL())
	if err != nil {
		return nil, err
	}
	defer surface.Free()

	return NewTextureFromSurface(surface), nil
}

func (f *Font) RenderTextureShadow(text string, color common.Color, shadowColor common.Color, offset common.Vector2) (*Texture, error) {
	textSurface, err := f.Font.RenderUTF8Blended(text, color.ToSDL())
	if err != nil {
		return nil, err
	}
	defer textSurface.Free()

	shadowSurface, err := f.Font.RenderUTF8Blended(text, shadowColor.ToSDL())
	if err != nil {
		return nil, err
	}
	defer shadowSurface.Free()

	surface, err := sdl.CreateRGBSurface(0, textSurface.W+int32(math.Abs(offset.X)), textSurface.H+int32(math.Abs(offset.Y)), 32, 0, 0, 0, 0)
	if err != nil {
		return nil, err
	}
	defer surface.Free()

	textDestRect := sdl.Rect{W: textSurface.W, H: textSurface.H}
	shadowDestRect := sdl.Rect{W: shadowSurface.W, H: shadowSurface.H}

	if offset.X >= 0 {
		shadowDestRect.X = int32(offset.X)
	} else {
		textDestRect.X = int32(-offset.X)
	}

	if offset.Y >= 0 {
		shadowDestRect.Y = int32(offset.Y)
	} else {
		textDestRect.Y = int32(-offset.Y)
	}

	_ = shadowSurface.Blit(&shadowSurface.ClipRect, surface, &shadowDestRect)
	_ = textSurface.Blit(&textSurface.ClipRect, surface, &textDestRect)

	return NewTextureFromSurface(surface), nil
}

func (f *Font) Close() {
	f.Font.Close()
}
