package di

type ServiceKey string

type Container struct {
	Services map[ServiceKey]any
}

func init() {
	InitServices()
}

var container Container

func InitServices() {
	container = Container{
		Services: make(map[ServiceKey]any),
	}
}

func Get(name ServiceKey) any {
	if service, ok := container.Services[name]; ok {
		return service
	}

	return nil
}

func Set(name ServiceKey, instance any) {
	container.Services[name] = instance
}
