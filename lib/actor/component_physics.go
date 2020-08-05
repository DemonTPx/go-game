package actor

import (
	"fmt"
	"github.com/DemonTPx/go-game/lib/common"
	"math"
	"time"
)

type PhysicsComponent struct {
	BaseComponent
	Velocity   common.Vector3
	friction   float64
	bounciness float64
}

func NewPhysicsComponent(velocity common.Vector3, friction float64, bounciness float64) *PhysicsComponent {
	return &PhysicsComponent{Velocity: velocity, friction: friction, bounciness: bounciness}
}

func (c *PhysicsComponent) Id() ComponentId {
	return Physics
}

func (c *PhysicsComponent) Name() string {
	return "PhysicsComponent"
}

func (c *PhysicsComponent) String() string {
	return fmt.Sprintf("<%s velocity=%s friction=%f bounciness=%f>", c.Name(), c.Velocity.String(), c.friction, c.bounciness)
}

func (c *PhysicsComponent) Update(delta time.Duration) {
	transform := c.owner.GetComponent(Transform)
	if transform == nil {
		return
	}

	t := transform.(*TransformComponent)

	if t.Position.X-t.Scale.X/2 <= 0 && c.Velocity.X < 0 {
		c.Velocity.X = -(c.Velocity.X * c.bounciness)
		t.Position.X = t.Scale.X / 2
	}
	if t.Position.X+t.Scale.X/2 > 1600 && c.Velocity.X > 0 {
		c.Velocity.X = -(c.Velocity.X * c.bounciness)
		t.Position.X = 1600 - t.Scale.X/2
	}

	if t.Position.Y-t.Scale.Y/2 <= 0 && c.Velocity.Y < 0 {
		c.Velocity.Y = -(c.Velocity.Y * c.bounciness)
		t.Position.Y = t.Scale.Y / 2
	}
	if t.Position.Y+t.Scale.Y/2 > 900 && c.Velocity.Y > 0 {
		c.Velocity.Y = -(c.Velocity.Y * c.bounciness)
		t.Position.Y = 900 - t.Scale.Y/2
	}

	deltaMs := float64(delta / time.Millisecond)

	c.Velocity = c.Velocity.MultiFloat64(math.Pow(1-c.friction, deltaMs))

	v := c.Velocity.MultiFloat64(deltaMs)

	t.Position = t.Position.Add(&v)
}
