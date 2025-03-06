//go:build wireinject

package main

import (
	"github.com/google/wire"
	"greeter/conf"
	"greeter/internal/biz"
	"greeter/internal/data"
	"greeter/internal/server"
	"greeter/internal/service"
	log "greeter/pkg/logger"
)

func InitApp() *App {
	wire.Build(
		wire.Struct(new(App), "*"),
		wire.Bind(new(log.Logger), new(*log.ZapLogger)), // 绑定接口和实现
		log.NewZapLogger,
		conf.NewConf,
		server.ProviderSet,
		data.ProviderSet,
		biz.ProviderSet,
		service.ProviderSet,
	)
	return &App{}
}
