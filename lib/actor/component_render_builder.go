package actor

import (
	"fmt"
	"github.com/DemonTPx/go-game/lib/common"

	"github.com/mitchellh/mapstructure"
)

type RenderComponentBuilder struct {
}

func NewRenderComponentBuilder() *RenderComponentBuilder {
	return &RenderComponentBuilder{}
}

func (b *RenderComponentBuilder) Build(data VariableConfig) (Component, error) {
	switch data.GetStringOr("type", "") {
	case "ellipse":
		color := common.Color{}
		err := data.Extract("color", &color)
		if err != nil {
			return nil, fmt.Errorf("invalid configuration for type '%s': %s", data["type"], err)
		}

		segments := data.GetIntOr("segments", 20)

		return NewEllipseRenderComponent(color, segments), nil
	case "rect":
		color := common.Color{}
		err := data.Extract("color", &color)
		if err != nil {
			return nil, fmt.Errorf("invalid configuration for type '%s': %s", data["type"], err)
		}

		return NewRectRenderComponent(color), nil
	default:
		return NewRenderComponent(), nil
	}
}

func extractColor(data VariableConfig, key string, v *common.Color) error {
	config, err := data.GetVariableConfig(key)
	if err != nil {
		return nil
	}

	err = mapstructure.Decode(config, v)

	if err != nil {
		return fmt.Errorf("invalid color config at key '%s'", key)
	}

	return nil
}
