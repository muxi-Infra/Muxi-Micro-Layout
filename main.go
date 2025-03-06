package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	_ "go.uber.org/automaxprocs"
	"sync"
)

// App is an application.
type App struct {
	Http *http.Server
	Grpc *grpc.Server
}

func main() {
	flag.Parse()
	ctx := context.Background()

	app := InitApp()

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := app.Http.Start(ctx); err != nil {
			panic(err)
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := app.Grpc.Start(ctx); err != nil {
			fmt.Println(err)
			panic(err)
		}
	}()
	wg.Wait()
}
