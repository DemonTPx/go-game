package actor

type ControlComponent struct {
	BaseComponent
}

func NewControlComponent() *ControlComponent {
	return &ControlComponent{}
}

func (c *ControlComponent) Id() ComponentId {
	return Control
}

func (c *ControlComponent) Name() string {
	return "ControlComponent"
}

func (c *ControlComponent) String() string {
	return "<" + c.Name() + ">"
}
