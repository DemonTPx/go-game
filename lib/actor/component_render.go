package actor

import (
	"github.com/demontpx/go-game/lib/actor/property"
)

type RenderComponent struct {
	BaseComponent
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

type BallRenderComponent struct {
	RenderComponent
	color property.Color
}

func NewBallRenderComponent(color property.Color) *BallRenderComponent {
	return &BallRenderComponent{color: color}
}

func (c *BallRenderComponent) Name() string {
	return "BallRenderComponent"
}

func (c *BallRenderComponent) String() string {
	return "<" + c.Name() + " color=" + c.color.String() + ">"
}
