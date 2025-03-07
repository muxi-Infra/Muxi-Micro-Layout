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

func InitApp(c *conf.Conf) *App {
	wire.Build(
		wire.Struct(new(App), "*"),
		log.ProviderSet,
		// conf.NewConf,
		server.ProviderSet,
		data.ProviderSet,
		biz.ProviderSet,
		service.ProviderSet,
	)
	return &App{}
}
