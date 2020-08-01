package actor

import (
	"fmt"

	"github.com/mitchellh/mapstructure"

	"github.com/DemonTPx/go-game/lib/actor/property"
)

type TransformComponentBuilder struct {
}

func NewTransformComponentBuilder() *TransformComponentBuilder {
	return &TransformComponentBuilder{}
}

func (b *TransformComponentBuilder) Build(data VariableConfig) (Component, error) {
	var err error

	position := property.NewVector3(0, 0, 0)
	rotation := property.NewVector3(0, 0, 0)
	scale := property.NewVector3(1, 1, 1)

	err = extractVector3(data, "position", &position)
	if err != nil {
		return nil, fmt.Errorf("invalid configuration for type '%s': %s", data["type"], err)
	}

	err = extractVector3(data, "rotation", &rotation)
	if err != nil {
		return nil, fmt.Errorf("invalid configuration for type '%s': %s", data["type"], err)
	}

	err = extractVector3(data, "scale", &scale)
	if err != nil {
		return nil, fmt.Errorf("invalid configuration for type '%s': %s", data["type"], err)
	}

	return NewTransformComponent(position, rotation, scale), nil
}

func extractVector3(data VariableConfig, key string, v *property.Vector3) error {
	config, err := data.GetVariableConfig(key)
	if err != nil {
		return nil
	}

	err = mapstructure.Decode(config, v)

	if err != nil {
		return fmt.Errorf("invalid vector3 config at key '%s'", key)
	}

	return nil
}
