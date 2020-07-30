package actor

type PhysicsComponentBuilder struct {
}

func NewPhysicsComponentBuilder() *PhysicsComponentBuilder {
	return &PhysicsComponentBuilder{}
}

func (b *PhysicsComponentBuilder) Build(data VariableConfig) (Component, error) {
	return NewPhysicsComponent(), nil
}
