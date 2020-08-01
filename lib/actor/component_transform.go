package actor

import "github.com/demontpx/go-game/lib/actor/property"

type TransformComponent struct {
	BaseComponent
	Position property.Vector3
	Rotation property.Vector3
	Scale    property.Vector3
}

func NewTransformComponent(position property.Vector3, rotation property.Vector3, scale property.Vector3) *TransformComponent {
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
	return "<" + c.Name() +
		" position=" + c.Position.String() +
		" rotation=" + c.Rotation.String() +
		" scale=" + c.Scale.String() +
		">"
}
