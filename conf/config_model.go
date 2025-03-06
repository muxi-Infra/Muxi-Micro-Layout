package conf

import "time"

type Conf struct {
	Server Server `mapstructure:"server"`
	Data   Data   `mapstructure:"data"`
}

type Server struct {
	Http Web `mapstructure:"http"`
	Grpc Web `mapstructure:"grpc"`
}
type Web struct {
	// Network string        `mapstructure:"network"`
	Addr    string        `mapstructure:"addr"`
	Timeout time.Duration `mapstructure:"timeout"`
}

type Data struct {
	Mysql Mysql `mapstructure:"mysql"`
	Redis Redis `mapstructure:"redis"`
}

type Mysql struct {
	DSN string `mapstructure:"dsn"`
}

type Redis struct {
	Addr         string        `mapstructure:"addr"`
	ReadTimeOut  time.Duration `mapstructure:"read_timeout"`
	WriteTimeOut time.Duration `mapstructure:"write_timeout"`
}
