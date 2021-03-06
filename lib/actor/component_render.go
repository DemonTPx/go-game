package actor

import (
	"fmt"
	"github.com/DemonTPx/go-game/lib/common"
	"github.com/DemonTPx/go-game/lib/render"
	gl "github.com/chsc/gogl/gl21"
	"math"
)

type RenderComponent struct {
	BaseComponent
	layer float64
}

type Renderer interface {
	Render(viewport common.Rect)
	RenderLayer() float64
}

func NewRenderComponent(layer float64) *RenderComponent {
	return &RenderComponent{
		layer: layer,
	}
}

func (c *RenderComponent) Id() ComponentId {
	return Render
}

func (c *RenderComponent) Name() string {
	return "RenderComponent"
}

func (c *RenderComponent) String() string {
	return fmt.Sprintf("<%s layer=%f>", c.Name(), c.layer)
}

func (c *RenderComponent) Render() {
}

func (c *RenderComponent) RenderLayer() float64 {
	return c.layer
}

type EllipseRenderComponent struct {
	RenderComponent
	color         common.Color
	texture       *render.Texture
	textureScale  float64
	textureOffset common.Vector2
	segments      int
}

func NewEllipseRenderComponent(
	layer float64,
	color common.Color,
	texture *render.Texture,
	textureScale float64,
	textureOffset common.Vector2,
	segments int,
) *EllipseRenderComponent {
	return &EllipseRenderComponent{
		RenderComponent: *NewRenderComponent(layer),
		color:           color,
		texture:         texture,
		textureScale:    textureScale,
		textureOffset:   textureOffset,
		segments:        segments,
	}
}

func (c *EllipseRenderComponent) Name() string {
	return "EllipseRenderComponent"
}

func (c *EllipseRenderComponent) String() string {
	return fmt.Sprintf("<%s layer=%f color=%s texture=%+v texture_scale=%f texture_offset=%f segments=%d>",
		c.Name(), c.layer, c.color.String(), c.texture, c.textureScale, c.textureOffset, c.segments)
}

func (c *EllipseRenderComponent) Render(viewport common.Rect) {
	transformComponent := c.owner.GetComponent(Transform)
	if transformComponent == nil {
		return
	}

	transform := transformComponent.(*TransformComponent)
	pos := transform.Position
	scale := transform.Scale

	gl.Color4f(gl.Float(c.color.R), gl.Float(c.color.G), gl.Float(c.color.B), gl.Float(c.color.A))

	if c.texture != nil {
		gl.ActiveTexture(gl.TEXTURE0)
		c.texture.Bind()
	}

	gl.Begin(gl.TRIANGLE_FAN)

	if c.texture != nil {
		gl.TexCoord2f(gl.Float(c.textureOffset.X), gl.Float(c.textureOffset.Y))
	}
	gl.Vertex3f(gl.Float(pos.X), gl.Float(pos.Y), gl.Float(pos.Z))
	segments := float64(c.segments)
	for n := float64(0); n <= segments; n++ {
		t := math.Pi * 2 * n / segments
		if c.texture != nil {
			gl.TexCoord2f(gl.Float(math.Sin(t)*c.textureScale+c.textureOffset.X), gl.Float(math.Cos(t)*c.textureScale+c.textureOffset.Y))
		}
		gl.Vertex3f(gl.Float(pos.X+math.Sin(t)*scale.X/2), gl.Float(pos.Y+math.Cos(t)*scale.Y/2), gl.Float(pos.Z))
	}
	gl.End()

	if c.texture != nil {
		gl.BindTexture(gl.TEXTURE_2D, 0)
	}
}

func (c *EllipseRenderComponent) Destroy() {
	if c.texture != nil {
		c.texture.Destroy()
	}
}

type RectRenderComponent struct {
	RenderComponent
	color   common.Color
	texture *render.Texture
}

func NewRectRenderComponent(
	layer float64,
	color common.Color,
	texture *render.Texture,
) *RectRenderComponent {
	return &RectRenderComponent{
		RenderComponent: *NewRenderComponent(layer),
		color:           color,
		texture:         texture,
	}
}

func (c *RectRenderComponent) Name() string {
	return "RectRenderComponent"
}

func (c *RectRenderComponent) String() string {
	return fmt.Sprintf("<%s layer=%f color=%s texture=%+v>", c.Name(), c.layer, c.color.String(), c.texture)
}

func (c *RectRenderComponent) Render(viewport common.Rect) {
	transformComponent := c.owner.GetComponent(Transform)
	if transformComponent == nil {
		return
	}

	transform := transformComponent.(*TransformComponent)
	pos := transform.Position
	scale := transform.Scale

	rect := common.NewRect(pos.X-scale.X/2, pos.Y-scale.Y/2, scale.X, scale.Y)

	gl.Color4f(gl.Float(c.color.R), gl.Float(c.color.G), gl.Float(c.color.B), gl.Float(c.color.A))

	if c.texture == nil {
		gl.Begin(gl.QUADS)
		gl.Vertex3f(gl.Float(rect.X), gl.Float(rect.Y), gl.Float(pos.Z))
		gl.Vertex3f(gl.Float(rect.X2()), gl.Float(rect.Y), gl.Float(pos.Z))
		gl.Vertex3f(gl.Float(rect.X2()), gl.Float(rect.Y2()), gl.Float(pos.Z))
		gl.Vertex3f(gl.Float(rect.X), gl.Float(rect.Y2()), gl.Float(pos.Z))
		gl.End()
	} else {
		gl.ActiveTexture(gl.TEXTURE0)
		c.texture.Bind()

		gl.Begin(gl.QUADS)
		gl.TexCoord2f(0, 0)
		gl.Vertex3f(gl.Float(rect.X), gl.Float(rect.Y), gl.Float(pos.Z))
		gl.TexCoord2f(1, 0)
		gl.Vertex3f(gl.Float(rect.X2()), gl.Float(rect.Y), gl.Float(pos.Z))
		gl.TexCoord2f(1, 1)
		gl.Vertex3f(gl.Float(rect.X2()), gl.Float(rect.Y2()), gl.Float(pos.Z))
		gl.TexCoord2f(0, 1)
		gl.Vertex3f(gl.Float(rect.X), gl.Float(rect.Y2()), gl.Float(pos.Z))
		gl.End()

		gl.BindTexture(gl.TEXTURE_2D, 0)
	}

}

func (c *RectRenderComponent) Destroy() {
	if c.texture != nil {
		c.texture.Destroy()
	}
}
