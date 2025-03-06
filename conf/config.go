package conf

import (
	"github.com/spf13/viper"
	"path"
	"runtime"
)

// NewConf 创建一个新的配置实例并进行初始化
func NewConf() *Conf {
	// 获取当前文件路径和目录
	_, filename, _, _ := runtime.Caller(0)
	confPath := path.Dir(filename)
	viper.AddConfigPath(confPath)

	// 配置 viper 以读取配置文件
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	// 读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
		return nil
	}

	// 解析初始配置
	var Config Conf
	err = viper.Unmarshal(&Config)
	if err != nil {
		panic(err)
		return nil
	}

	return &Config
}
