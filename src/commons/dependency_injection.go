package commons

import "sync"

var ContainerInjectable *Container = NewContainer()

type Container struct {
	services map[string]interface{}
	lock     sync.RWMutex
}

func NewContainer() *Container {
	return &Container{
		services: make(map[string]interface{}),
	}
}

func (c *Container) Register(name string, service interface{}) {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.services[name] = service
}

func (c *Container) Get(name string) interface{} {
	c.lock.RLock()
	defer c.lock.RUnlock()

	service, ok := c.services[name]
	if !ok {
		return nil
	}

	return service
}
