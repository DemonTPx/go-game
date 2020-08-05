package actor

import (
	"fmt"
	"github.com/DemonTPx/go-game/lib/common"
	"github.com/DemonTPx/go-game/lib/render"

	"github.com/mitchellh/mapstructure"
)

type RenderComponentBuilder struct {
}

func NewRenderComponentBuilder() *RenderComponentBuilder {
	return &RenderComponentBuilder{}
}

func (b *RenderComponentBuilder) Build(data VariableConfig) (Component, error) {
	var err error

	switch data.GetStringOr("type", "") {
	case "ellipse":
		color := common.NewColorWhite()
		err = data.Extract("color", &color)
		if err != nil {
			return nil, fmt.Errorf("invalid configuration for type '%s': %s", data["type"], err)
		}

		textureFilename := data.GetStringOr("texture", "")
		var texture *render.Texture
		if textureFilename != "" {
			texture, err = render.NewTextureFromFile(textureFilename)
			if err != nil {
				return nil, fmt.Errorf("could not load texture configured for type '%s': %s", data["type"], err)
			}
		}

		textureScale := data.GetFloat64Or("texture_scale", 0.5)

		textureOffset := common.NewVector2(0.5, 0.5)
		err = data.Extract("texture_offset", &textureOffset)
		if err != nil {
			return nil, fmt.Errorf("invalid configuration for type '%s': %s", data["type"], err)
		}

		segments := data.GetIntOr("segments", 20)

		return NewEllipseRenderComponent(color, texture, textureScale, textureOffset, segments), nil
	case "rect":
		color := common.NewColorWhite()
		err = data.Extract("color", &color)
		if err != nil {
			return nil, fmt.Errorf("invalid configuration for type '%s': %s", data["type"], err)
		}

		textureFilename := data.GetStringOr("texture", "")
		var texture *render.Texture
		if textureFilename != "" {
			texture, err = render.NewTextureFromFile(textureFilename)
			if err != nil {
				return nil, fmt.Errorf("could not load texture configured for type '%s': %s", data["type"], err)
			}
		}

		return NewRectRenderComponent(color, texture), nil
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
