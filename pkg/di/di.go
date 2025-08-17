package dependency_container

import "fmt"

type DependencyContainer struct {
	Services map[string]interface{}
}

var Container = &DependencyContainer{
	Services: make(map[string]interface{}),
}

func (cont *DependencyContainer) Add(name string, value interface{}) {
	cont.Services[name] = value
}

func (cont *DependencyContainer) Get(name string) (interface{}, error) {
	currentService := cont.Services[name]
	if currentService == nil {
		return nil, fmt.Errorf("service %s not found", name)
	}

	return currentService, nil
}
