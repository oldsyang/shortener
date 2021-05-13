package initialize

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
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
	zap.S().Infof("启动配置文件：%s", configFileName)
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
		zap.S().Infof("%s config changed", configFileName)
		_ = v.ReadInConfig()
		_ = v.Unmarshal(global.ServerConfig)
	})
}
