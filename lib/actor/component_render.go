package actor

import (
	"fmt"
	"github.com/DemonTPx/go-game/lib/actor/property"
	gl "github.com/chsc/gogl/gl21"
	"math"
)

type RenderComponent struct {
	BaseComponent
}

type Renderer interface {
	Render()
}

func NewRenderComponent() *RenderComponent {
	return &RenderComponent{}
}

func (c *RenderComponent) Id() ComponentId {
	return Render
}

func (c *RenderComponent) Name() string {
	return "RenderComponent"
}

func (c *RenderComponent) String() string {
	return "<" + c.Name() + ">"
}

func (c *RenderComponent) Render() {
}

type EllipseRenderComponent struct {
	RenderComponent
	color    property.Color
	segments int
}

func NewEllipseRenderComponent(color property.Color, segments int) *EllipseRenderComponent {
	return &EllipseRenderComponent{color: color, segments: segments}
}

func (c *EllipseRenderComponent) Name() string {
	return "EllipseRenderComponent"
}

func (c *EllipseRenderComponent) String() string {
	return fmt.Sprintf("<%s color=%s segments=%d>", c.Name(), c.color.String(), c.segments)
}

func (c *EllipseRenderComponent) Render() {
	transformComponent := c.owner.GetComponent(Transform)
	if transformComponent == nil {
		return
	}

	transform := transformComponent.(*TransformComponent)
	pos := transform.Position
	scale := transform.Scale

	gl.Begin(gl.TRIANGLE_FAN)
	gl.Color4f(gl.Float(c.color.R), gl.Float(c.color.G), gl.Float(c.color.B), gl.Float(c.color.A))
	gl.Vertex2f(gl.Float(pos.X), gl.Float(pos.Y))
	segments := float64(c.segments)
	for n := float64(0); n <= segments; n++ {
		t := math.Pi * 2 * n / segments
		gl.Vertex2f(gl.Float(pos.X+math.Sin(t)*scale.X/2), gl.Float(pos.Y+math.Cos(t)*scale.Y/2))
	}
	gl.End()
}

type RectRenderComponent struct {
	RenderComponent
	color property.Color
}

func NewRectRenderComponent(color property.Color) *RectRenderComponent {
	return &RectRenderComponent{color: color}
}

func (c *RectRenderComponent) Name() string {
	return "RectRenderComponent"
}

func (c *RectRenderComponent) String() string {
	return fmt.Sprintf("<%s color=%s>", c.Name(), c.color.String())
}

func (c *RectRenderComponent) Render() {
	transformComponent := c.owner.GetComponent(Transform)
	if transformComponent == nil {
		return
	}

	transform := transformComponent.(*TransformComponent)
	pos := transform.Position
	scale := transform.Scale

	gl.Begin(gl.TRIANGLE_STRIP)
	gl.Color4f(gl.Float(c.color.R), gl.Float(c.color.G), gl.Float(c.color.B), gl.Float(c.color.A))
	gl.Vertex2f(gl.Float(pos.X-scale.X/2), gl.Float(pos.Y-scale.Y/2))
	gl.Vertex2f(gl.Float(pos.X+scale.X/2), gl.Float(pos.Y-scale.Y/2))
	gl.Vertex2f(gl.Float(pos.X-scale.X/2), gl.Float(pos.Y+scale.Y/2))
	gl.Vertex2f(gl.Float(pos.X+scale.X/2), gl.Float(pos.Y+scale.Y/2))
	gl.End()
}
