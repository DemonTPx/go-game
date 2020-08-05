package actor

type ControlComponentBuilder struct {
}

func NewControlComponentBuilder() *ControlComponentBuilder {
	return &ControlComponentBuilder{}
}

func (b *ControlComponentBuilder) Build(data VariableConfig) (Component, error) {
	switch data.GetStringOr("type", "") {
	case "paddle":
		return NewPaddleControlComponent(
			data.GetFloat64Or("velocity_factor", 1),
			data.GetFloat64Or("velocity_max", 50),
		), nil
	case "free":
		return NewFreeControlComponent(
			data.GetFloat64Or("velocity_factor", 1),
			data.GetFloat64Or("velocity_max", 50),
		), nil
	default:
		return NewControlComponent(), nil
	}
}
