package actor

import "github.com/demontpx/go-game/lib/actor/property"

type TransformComponent struct {
	BaseComponent
	V property.Vector3
}

func NewTransformComponent(v property.Vector3) *TransformComponent {
	return &TransformComponent{V: v}
}

func (c *TransformComponent) Id() ComponentId {
	return Transform
}

func (c *TransformComponent) Name() string {
	return "TransformComponent"
}

func (c *TransformComponent) String() string {
	return "<" + c.Name() + " " + c.V.String() + ">"
}
