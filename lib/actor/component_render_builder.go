package actor

import (
	"fmt"

	"github.com/mitchellh/mapstructure"

	"github.com/DemonTPx/go-game/lib/actor/property"
)

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
		color := property.Color{}
		err = mapstructure.Decode(colorConfig, &color)
		if err != nil {
			return nil, fmt.Errorf("invalid configuration for type '%s': %s", data["type"], err)
		}

		segments := data.GetIntOr("segments", 20)

		return NewBallRenderComponent(color, segments), nil
	default:
		return NewRenderComponent(), nil
	}
}
