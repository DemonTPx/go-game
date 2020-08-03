package render

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
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

func (f *Font) RenderTexture(text string) (*Texture, error) {
	surface, err := f.Font.RenderUTF8Blended(text, sdl.Color{R: 0xff, G: 0x00, B: 0x7f, A: 0xff})
	if err != nil {
		return nil, err
	}
	defer surface.Free()

	return NewTextureFromSurface(surface), nil
}

func (f *Font) Close() {
	f.Font.Close()
}
