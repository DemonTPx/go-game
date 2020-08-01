package actor

type ControlComponentBuilder struct {
}

func NewControlComponentBuilder() *ControlComponentBuilder {
	return &ControlComponentBuilder{}
}

func (b *ControlComponentBuilder) Build(data VariableConfig) (Component, error) {
	return NewControlComponent(), nil
}
