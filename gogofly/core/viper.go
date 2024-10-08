package core

import (
	"fmt"
	"github.com/dotdancer/gogofly/global"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func InitConfig() {
	//viper.SetConfigName("config")
	//viper.SetConfigType("yaml")
	//viper.AddConfigPath("./")
	//err := viper.ReadInConfig()
	//if err != nil {
	//	panic(err)
	//}

	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("./")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&global.Config); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&global.Config); err != nil {
		panic(err)
	}
}
