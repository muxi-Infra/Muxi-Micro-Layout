package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	_ "go.uber.org/automaxprocs"
	"greeter/conf"
	"sync"
)

// App is an application.
type App struct {
	Http *http.Server
	Grpc *grpc.Server
}

func main() {
	flag.Parse()

	c := conf.NewConf()
	app := InitApp(c)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := app.Http.Start(context.Background()); err != nil {
			panic(err)
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := app.Grpc.Start(context.Background()); err != nil {
			fmt.Println(err)
			panic(err)
		}
	}()
	wg.Wait()
}
