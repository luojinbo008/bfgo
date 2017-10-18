package app

import (
	"github.com/luojinbo008/bfgo/container"
)

func Register(name string, creator container.Creator) {
	container.DefaultContainer.Register(name, creator)
}

func ConfigureAll(cfg map[string]interface{}) error {
	return container.DefaultContainer.ConfigureAll(cfg)
}

func Get(names ...string) (interface{}, error) {
	return container.DefaultContainer.Get(names...)
}

func GetContainer(name string) (*container.Container, error) {
	return container.DefaultContainer.GetContainer(name)
}