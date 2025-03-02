package conf

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	log "greeter/pkg/logger"
	"path"
	"runtime"

	"sync"
	"time"
)

type Config struct {
	conf *Conf
	sync.RWMutex
	log *log.ZapLogger
}

type MutexConf interface {
	GetConfig() Conf
	UpdateConfig(newConfig Conf)
}

func (cw *Config) GetConfig() *Conf {
	cw.RLock()
	defer cw.RUnlock()
	return cw.conf
}

func (cw *Config) UpdateConfig(newConfig Conf) {
	cw.Lock()
	defer cw.Unlock()
	cw.conf = &newConfig
}

// NewConfig 创建一个新的配置实例并进行初始化
func NewConfig(log *log.ZapLogger) (*Config, error) {
	globalConfig := &Config{log: log}

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
		globalConfig.log.Errorf("Error reading config file: %s", err)
		return nil, err
	}

	// 解析初始配置
	var initialConf Conf
	err = viper.Unmarshal(&initialConf)
	if err != nil {
		globalConfig.log.Errorf("Error unmarshalling config: %s", err)
		return nil, err
	}
	globalConfig.UpdateConfig(initialConf)

	// 开启热更新监听
	viper.WatchConfig()

	// 配置文件变化时触发的回调
	viper.OnConfigChange(func(e fsnotify.Event) {
		go func() {
			newConf := Conf{}
			err := viper.Unmarshal(&newConf)
			if err != nil {
				globalConfig.log.Errorf("Error unmarshalling updated config: %s", err)
				return
			}
			globalConfig.UpdateConfig(newConf)
			globalConfig.log.Infof("Config updated successfully at %v", time.Now().Format("2006-01-02 15:04:05"))
		}()
		globalConfig.log.Infof("Config file changed: %v, time: %v", e.Name, time.Now().Format("2006-01-02 15:04:05"))
	})

	return globalConfig, nil
}
