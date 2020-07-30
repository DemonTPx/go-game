package actor

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Loader struct {
	lastId   Id
	builders map[string]Builder
}

func NewLoader() *Loader {
	return &Loader{
		lastId: Id(0),
		builders: map[string]Builder{
			"transform": NewTransformComponentBuilder(),
			"physics":   NewPhysicsComponentBuilder(),
			"render":    NewRenderComponentBuilder(),
		},
	}
}

func (l *Loader) GetNextActorId() Id {
	l.lastId++

	return l.lastId
}

func (l *Loader) LoadActorFromFile(filename string) (*Actor, error) {
	var err error

	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file %s: %s", filename, err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("error while reading file %s: %s", filename, err)
	}

	var config []ComponentConfig

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("error while parsing yaml from file %s: %s", filename, err)
	}

	actor := NewActor(l.GetNextActorId())

	for _, c := range config {
		builder, ok := l.builders[c.Type]
		if !ok {
			return nil, fmt.Errorf("no builder found for type '%s'", c.Type)
		}

		component, err := builder.Build(c.Data)
		if err != nil {
			return nil, fmt.Errorf("error while building component of type '%s': %s", c.Type, err)
		}

		actor.AddComponent(component)
	}

	return actor, nil
}
