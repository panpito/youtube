package prototype

import "errors"

type IPrototype interface {
	Clone() interface{}
}

type Registry struct {
	prototypes map[string]IPrototype
}

func NewRegistry() *Registry {
	return &Registry{prototypes: make(map[string]IPrototype)}
}

func (registry *Registry) Add(id string, proto IPrototype) {
	registry.prototypes[id] = proto
}

func (registry *Registry) Get(id string) (IPrototype, error) {
	proto, ok := registry.prototypes[id]
	if !ok {
		return nil, errors.New("not found")
	}

	return proto, nil
}
