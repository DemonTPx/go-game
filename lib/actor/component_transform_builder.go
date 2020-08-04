package actor

import (
	"fmt"
	"github.com/DemonTPx/go-game/lib/common"
)

type TransformComponentBuilder struct {
}

func NewTransformComponentBuilder() *TransformComponentBuilder {
	return &TransformComponentBuilder{}
}

func (b *TransformComponentBuilder) Build(data VariableConfig) (Component, error) {
	var err error

	position := common.NewVector3(0, 0, 0)
	rotation := common.NewVector3(0, 0, 0)
	scale := common.NewVector3(1, 1, 1)

	err = data.Extract("position", &position)
	if err != nil {
		return nil, fmt.Errorf("invalid configuration for type '%s': %s", data["type"], err)
	}

	err = data.Extract("rotation", &rotation)
	if err != nil {
		return nil, fmt.Errorf("invalid configuration for type '%s': %s", data["type"], err)
	}

	err = data.Extract("scale", &scale)
	if err != nil {
		return nil, fmt.Errorf("invalid configuration for type '%s': %s", data["type"], err)
	}

	return NewTransformComponent(position, rotation, scale), nil
}
