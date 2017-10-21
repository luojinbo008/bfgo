package bfgo
import (
	"errors"
	"log"
	"github.com/spf13/viper"
	"github.com/luojinbo008/bfgo/app"
	"github.com/luojinbo008/bfgo/thrift"
	"github.com/luojinbo008/bfgo/database/mysql"
	"github.com/luojinbo008/bfgo/database/redis"
)

func Init(configFile string, args ...interface{}) error {

	viper.SetConfigFile(configFile)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		errStr := "Config parse error," + err.Error()
		log.Fatal(errStr)
		return err
	} else {
		log.Println("Using config file:", viper.ConfigFileUsed())
	}

	// 初始化服务类型
	runMode := viper.GetString("server.type")
	if len(runMode) != 0 {
		switch runMode {
		case "thrift":
			if len(args) == 1 {
				thrift.Init(args[0])
			} else{
				return errors.New("args error")
			}
		default:
			return errors.New("run error with type：" + runMode)
		}
	}

	// 初始化组件
	if len(viper.GetStringMap("components.mysql")) != 0 {
		app.Register("mysql", mysql.Creator)
	}
	if len(viper.GetStringMap("components.redis")) != 0 {
		app.Register("redis", redis.Creator)
	}

	return app.ConfigureAll(viper.GetStringMap("components"))
}

func Run() error {
	runMode := viper.GetString("server.type")
	if len(runMode) != 0{
		switch runMode {
		case "thrift":
			err := thrift.Run()
			if err != nil{
				log.Fatal("RunThriftServer: ", err)
			}
		default:
			return errors.New("run error with type：" + runMode)
		}
	}
	return errors.New("run error with type：" + runMode)
}