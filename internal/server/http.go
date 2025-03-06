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
func NewHTTPServer(c *conf.Conf, greeter *service.GreeterService, log *log.ZapLogger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Server.Http.Addr != "" {
		opts = append(opts, http.Address(c.Server.Http.Addr))
	}
	if c.Server.Http.Timeout > 0 {
		opts = append(opts, http.Timeout(c.Server.Http.Timeout))
	}
	srv := http.NewServer(opts...)
	v1.RegisterGreeterServiceHTTPServer(srv, greeter)
	return srv
}
