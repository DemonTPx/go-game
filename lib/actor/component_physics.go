package actor

type PhysicsComponent struct {
	BaseComponent
}

func NewPhysicsComponent() *PhysicsComponent {
	return &PhysicsComponent{}
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
