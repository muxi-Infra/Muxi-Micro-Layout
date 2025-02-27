//go:build wireinject

package main

import (
	"github.com/google/wire"
	"greeter/conf"
	"greeter/internal"
	"greeter/internal/biz"
	"greeter/internal/data"
	"greeter/internal/server"
	"greeter/internal/service"
)

func InitApp() *App {
	wire.Build(
		wire.Struct(new(App), "*"),
		conf.NewData,
		conf.NewServer,
		internal.NewLogger,
		server.ProviderSet,
		data.ProviderSet,
		biz.ProviderSet,
		service.ProviderSet,
	)
	return &App{}
}
