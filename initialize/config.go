package initialize

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
	"shortener/global"
)

func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func InitConfig() {
	debug := GetEnvInfo("SHORTENER_DEV")
	configFileName := "pro.yaml"
	if debug {
		configFileName = "dev.yaml"
	}
	file,_ := os.Getwd()
	configFileName = fmt.Sprintf("%s/%s", file, configFileName)

	v := viper.New()
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := v.Unmarshal(global.ServerConfig); err != nil {
		panic(err)
	}

	// viper 动态监听
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config changed")
		_ = v.ReadInConfig()
		_ = v.Unmarshal(global.ServerConfig)
	})
}
