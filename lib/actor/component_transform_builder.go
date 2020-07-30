package actor

import (
	"fmt"

	"github.com/mitchellh/mapstructure"

	"github.com/demontpx/go-game/lib/actor/property"
)

type TransformComponentBuilder struct {
}

func NewTransformComponentBuilder() *TransformComponentBuilder {
	return &TransformComponentBuilder{}
}

func (b *TransformComponentBuilder) Build(data VariableConfig) (Component, error) {
	var err error
	v := property.Vector3{}
	err = mapstructure.Decode(data, &v)
	if err != nil {
		return nil, fmt.Errorf("invalid configuration for type '%s': %s", data["type"], err)
	}

	return NewTransformComponent(v), nil
}
