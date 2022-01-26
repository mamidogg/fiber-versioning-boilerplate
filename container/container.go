package container

import (
	"log"

	"github.com/aofdev/fiber-versioning-boilerplate/api"
	"go.uber.org/dig"
)

// Container ...
type Container struct {
	container *dig.Container
}

// Configure ...
func (c *Container) Configure() error {
	// constructors := []interface{}{
	// 	adapters.NewMysqlServer,
	// }

	// for _, constructor := range constructors {
	// 	if err := c.container.Provide(constructor); err != nil {
	// 		return err
	// 	}
	// }

	return nil
}

// Start ...
func (c *Container) Start() error {
	log.Println("Start Container")
	if err := c.container.Invoke(func() {
		api.StartServer()
		log.Println("Start Start Server")

	}); err != nil {
		log.Printf("%s", err)
		return err
	}
	return nil
}

// NewContainer ...
func NewContainer() (*Container, error) {
	log.Println("this file should be set dependency injection by using uber-dig")
	d := dig.New()
	container := &Container{
		container: d,
	}
	if err := container.Configure(); err != nil {
		return nil, err
	}
	return container, nil
}
