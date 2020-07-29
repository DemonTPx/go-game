package actor

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
)

type ComponentBuilder interface {
	Build(data VariableConfig) (Component, error)
}

type RenderComponentBuilder struct {
}

func NewRenderComponentBuilder() *RenderComponentBuilder {
	return &RenderComponentBuilder{}
}

func (b *RenderComponentBuilder) Build(data VariableConfig) (Component, error) {
	switch data.GetStringOr("type", "") {
	case "ball":
		colorConfig, err := data.GetVariableConfig("color")
		if err != nil {
			return nil, fmt.Errorf("invalid configuration for type '%s': %s", data["type"], err)
		}
		color := Color{}
		err = mapstructure.Decode(colorConfig, &color)
		if err != nil {
			return nil, fmt.Errorf("invalid configuration for type '%s': %s", data["type"], err)
		}

		return NewBallRenderComponent(color), nil
	default:
		return NewRenderComponent(), nil
	}
}

type PhysicsComponentBuilder struct {
}

func NewPhysicsComponentBuilder() *PhysicsComponentBuilder {
	return &PhysicsComponentBuilder{}
}

func (b *PhysicsComponentBuilder) Build(data VariableConfig) (Component, error) {
	return NewPhysicsComponent(), nil
}
