package app

import (
	"github.com/luojinbo008/bfgo/container"
	"github.com/luojinbo008/bfgo/database/mysql"
	rd "github.com/luojinbo008/bfgo/database/redis"
	"github.com/go-redis/redis"
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

func GetRedis(name string) (*redis.Client, error) {
	r, err := container.DefaultContainer.Get("redis", name)
	if err == nil {
		if rr, ok := r.(*redis.Client); ok {
			return rr, nil
		}
		return nil, err
	}
	return nil, err
}

func GetMySQL(name string) (*mysql.DB, error) {
	instance, err := container.DefaultContainer.Get("mysql", name)
	if err != nil {
		return nil, err
	}
	if db1, ok := instance.(*mysql.DB); ok {
		return db1, nil
	}
	return nil, err
}

type Model interface {
	ConnName() string
}

func UseModel(name string, subName string, write bool) interface{} {
	d, err := Get(name, subName)
	if err == nil {
		switch name {
		case "mysql":
			return d.(*mysql.DB).Get(write)
		case "redis":
			return d.(*rd.RedisDB).Get(write)
		}
	}
	return nil
}