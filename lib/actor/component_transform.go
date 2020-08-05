package actor

import (
	"fmt"
	"github.com/DemonTPx/go-game/lib/common"
)

type TransformComponent struct {
	BaseComponent
	Position common.Vector3
	Rotation common.Vector3
	Scale    common.Vector3
}

func NewTransformComponent(position common.Vector3, rotation common.Vector3, scale common.Vector3) *TransformComponent {
	return &TransformComponent{
		Position: position,
		Rotation: rotation,
		Scale:    scale,
	}
}

func (c *TransformComponent) Id() ComponentId {
	return Transform
}

func (c *TransformComponent) Name() string {
	return "TransformComponent"
}

func (c *TransformComponent) String() string {
	return fmt.Sprintf("<%s position=%s rotation=%s scale=%s>", c.Name(), c.Position.String(), c.Rotation.String(), c.Scale.String())
}
