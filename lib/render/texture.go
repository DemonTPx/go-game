package render

import (
	"fmt"
	gl "github.com/chsc/gogl/gl21"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type Texture struct {
	Id gl.Uint
	W  int32
	H  int32
}

func NewTextureFromSurface(s *sdl.Surface) *Texture {
	var id gl.Uint

	gl.GenTextures(1, &id)
	gl.BindTexture(gl.TEXTURE_2D, id)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, gl.Sizei(s.W), gl.Sizei(s.H), 0, formatFromSurface(s), gl.UNSIGNED_BYTE, gl.Pointer(s.Data()))

	return &Texture{
		Id: id,
		W:  s.W,
		H:  s.H,
	}
}

func formatFromSurface(s *sdl.Surface) gl.Enum {
	if s.Format.BytesPerPixel == 4 {
		if s.Format.Rmask == 0xff {
			return gl.RGBA
		} else {
			return gl.BGRA
		}
	}

	if s.Format.Rmask == 0xff {
		return gl.RGB
	}

	return gl.BGR
}

func NewTextureFromFile(file string) (*Texture, error) {
	surface, err := img.Load(file)
	if err != nil {
		return nil, fmt.Errorf("error while loading texture: %s", err)
	}
	defer surface.Free()

	return NewTextureFromSurface(surface), nil
}

func (t *Texture) Bind() {
	gl.BindTexture(gl.TEXTURE_2D, t.Id)
}

func (t *Texture) Draw(x, y, z gl.Float) {
	gl.ActiveTexture(gl.TEXTURE0)
	t.Bind()

	gl.Color4f(1, 1, 1, 1)

	gl.Begin(gl.QUADS)

	gl.TexCoord2f(0, 0)
	gl.Vertex3f(x, y, z)

	gl.TexCoord2f(1, 0)
	gl.Vertex3f(x+gl.Float(t.W), y, z)

	gl.TexCoord2f(1, 1)
	gl.Vertex3f(x+gl.Float(t.W), y+gl.Float(t.H), z)

	gl.TexCoord2f(0, 1)
	gl.Vertex3f(x, y+gl.Float(t.H), z)

	gl.End()

	gl.BindTexture(gl.TEXTURE_2D, 0)
}

func (t *Texture) Destroy() {
	gl.DeleteTextures(1, &t.Id)
}
