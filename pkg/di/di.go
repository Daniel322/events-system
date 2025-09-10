package dependency_container

import "fmt"

type DependencyContainer struct {
	Services map[string]interface{}
}

func NewDIContainer() *DependencyContainer {
	return &DependencyContainer{
		Services: make(map[string]interface{}),
	}
}

func (cont *DependencyContainer) Add(name string, value interface{}) {
	cont.Services[name] = value
}

func (cont *DependencyContainer) Get(name string) (interface{}, error) {
	currentService := cont.Services[name]
	if currentService == nil {
		return nil, fmt.Errorf("service %s not found", name)
	}

	return &currentService, nil
}

func (cont *DependencyContainer) MultiGet(names []string) (map[string]interface{}, error) {
	services := make(map[string]interface{})
	for _, name := range names {
		currentService := cont.Services[name]
		if currentService == nil {
			return nil, fmt.Errorf("service %s not found", name)
		}
		services[name] = currentService
	}

	return services, nil
}
