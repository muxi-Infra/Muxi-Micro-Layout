package conf

import (
	"github.com/fsnotify/fsnotify"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/spf13/viper"
	"path"
	"runtime"

	"sync"
	"time"
)

type Config struct {
	conf Conf
	sync.RWMutex
}

type MutexConf interface {
	GetConfig() Conf
	UpdateConfig(newConfig Conf)
}

func (cw *Config) GetConfig() Conf {
	cw.RLock()
	defer cw.RUnlock()
	return cw.conf
}

func (cw *Config) UpdateConfig(newConfig Conf) {
	cw.Lock()
	defer cw.Unlock()
	cw.conf = newConfig
}

func NewConfig() (*Config, error) {
	globalConfig := &Config{}

	_, filename, _, _ := runtime.Caller(0) // 获取当前文件（config.go）路径
	confPath := path.Dir(filename)         // 获取当前文件目录
	viper.AddConfigPath(confPath)          // 设置配置文件目录
	// fmt.Println("confPath:", confPath)

	// 初始化 Viper 配置
	viper.SetConfigName("config") // 配置文件名
	viper.SetConfigType("yaml")   // 配置文件类型

	// 加载初始配置
	err := viper.ReadInConfig()
	if err != nil {
		log.Errorf("Error reading config file: %s", err)
		return nil, err
	}

	// 解析初始配置到 Config
	var initialConf Conf
	err = viper.Unmarshal(&initialConf)
	if err != nil {
		log.Errorf("Error unmarshalling config: %s", err)
		return nil, err
	}
	globalConfig.UpdateConfig(initialConf)

	// 开启热更新监听
	viper.WatchConfig()

	// 在配置变更时触发的回调
	viper.OnConfigChange(func(e fsnotify.Event) {
		go func() {
			newConf := Conf{}
			err := viper.Unmarshal(&newConf)
			if err != nil {
				log.Fatalf("Error unmarshalling updated config: %s", err)
				return
			}
			globalConfig.UpdateConfig(newConf)
			log.Infof("Config updated successfully:%v\n", time.Now().Format("2006-01-02 15:04:05"))
		}()
		log.Infof("Config file changed:%v,time:%v\n", e.Name, time.Now().Format("2006-01-02 15:04:05"))
	})
	return globalConfig, nil
}
