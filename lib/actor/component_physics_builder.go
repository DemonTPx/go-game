package actor

import (
	"fmt"
	"github.com/DemonTPx/go-game/lib/actor/property"
)

type PhysicsComponentBuilder struct {
}

func NewPhysicsComponentBuilder() *PhysicsComponentBuilder {
	return &PhysicsComponentBuilder{}
}

func (b *PhysicsComponentBuilder) Build(data VariableConfig) (Component, error) {
	var err error

	velocity := property.NewVector3(0, 0, 0)
	err = data.Extract("velocity", &velocity)
	if err != nil {
		return nil, fmt.Errorf("invalid configuration for type '%s': %s", data["type"], err)
	}

	friction := data.GetFloat64Or("friction", 0)
	bounciness := data.GetFloat64Or("bounciness", 1)

	return NewPhysicsComponent(velocity, friction, bounciness), nil
}
