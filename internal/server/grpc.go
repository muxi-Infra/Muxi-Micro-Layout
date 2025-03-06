package server

import (
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	v1 "greeter/api/v1"
	"greeter/conf"
	"greeter/internal/service"
	log "greeter/pkg/logger"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Conf, greeter *service.GreeterService, log *log.ZapLogger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
		),
	}
	//if c.Grpc.Network != "" {
	//	opts = append(opts, grpc.Network(c.Grpc.Network))
	//}
	if c.Server.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Server.Grpc.Addr))
	}
	//if c.Grpc.Timeout != nil {
	//	opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	//}
	srv := grpc.NewServer(opts...)
	v1.RegisterGreeterServiceServer(srv, greeter)
	return srv
}
