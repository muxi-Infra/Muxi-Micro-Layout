package main

import (
	"context"
	"flag"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	_ "go.uber.org/automaxprocs"
	"greeter/conf"
	log "greeter/pkg/logger"
)

// App is a application.
type App struct {
	Http *http.Server
	Grpc *grpc.Server
}

func main() {
	flag.Parse()
	c, err := conf.NewConfig(log.NewZapLogger())
	if err != nil {
		panic(err)
	}

	cc := c.GetConfig()

	ctx := context.WithValue(context.Background(), "config", cc)

	app := InitApp()
	if err := app.Http.Start(ctx); err != nil {
		panic(err)
	}
	if err := app.Grpc.Start(ctx); err != nil {
		panic(err)
	}
}
