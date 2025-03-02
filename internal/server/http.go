package server

import (
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	v1 "greeter/api/v1"
	"greeter/conf"
	"greeter/internal/service"
	log "greeter/pkg/logger"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, greeter *service.GreeterService, log *log.ZapLogger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	//if c.Http.Network != "" {
	//	opts = append(opts, http.Network(c.Http.Network))
	//}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	//if c.Http.Timeout != nil {
	//	opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	//}
	srv := http.NewServer(opts...)
	v1.RegisterGreeterServiceHTTPServer(srv, greeter)
	return srv
}
