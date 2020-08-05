package actor

import (
	"fmt"
	"github.com/DemonTPx/go-game/lib/common"
)

type PhysicsComponentBuilder struct {
}

func NewPhysicsComponentBuilder() *PhysicsComponentBuilder {
	return &PhysicsComponentBuilder{}
}

func (b *PhysicsComponentBuilder) Build(data VariableConfig) (Component, error) {
	var err error

	velocity := common.NewVector3(0, 0, 0)
	err = data.Extract("velocity", &velocity)
	if err != nil {
		return nil, fmt.Errorf("invalid configuration for type '%s': %s", data["type"], err)
	}

	return NewPhysicsComponent(
		velocity,
		data.GetFloat64Or("friction", 0),
		data.GetFloat64Or("bounciness", 1),
	), nil
}
