package actor

import (
	"github.com/DemonTPx/go-game/lib/actor/property"
	"time"
)

type PhysicsComponent struct {
	BaseComponent
	Velocity property.Vector3
}

func NewPhysicsComponent(velocity property.Vector3) *PhysicsComponent {
	return &PhysicsComponent{Velocity: velocity}
}

func (c *PhysicsComponent) Id() ComponentId {
	return Physics
}

func (c *PhysicsComponent) Name() string {
	return "PhysicsComponent"
}

func (c *PhysicsComponent) String() string {
	return "<" + c.Name() + ">"
}

func (c *PhysicsComponent) Update(delta time.Duration) {
	transform := c.owner.GetComponent(Transform)
	if transform == nil {
		return
	}

	t := transform.(*TransformComponent)

	if t.Position.X-t.Scale.X <= 0 && c.Velocity.X < 0 {
		c.Velocity.X = -c.Velocity.X
	}
	if t.Position.X+t.Scale.X > 1600 && c.Velocity.X > 0 {
		c.Velocity.X = -c.Velocity.X
	}

	if t.Position.Y-t.Scale.Y <= 0 && c.Velocity.Y < 0 {
		c.Velocity.Y = -c.Velocity.Y
	}
	if t.Position.Y+t.Scale.Y > 900 && c.Velocity.Y > 0 {
		c.Velocity.Y = -c.Velocity.Y
	}

	t.Position = t.Position.Add(&c.Velocity)
}
