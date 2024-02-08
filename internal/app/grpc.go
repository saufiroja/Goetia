package app

import (
	"context"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	internalGrpc "github.com/saufiroja/cqrs/internal/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type Grpc struct {
	*grpc.Server
}

func NewGrpc(module *Module) *Grpc {
	grpcServer := grpc.NewServer(
		grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
		grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
	)
	reflection.Register(grpcServer)

	internalGrpc.RegisterTodosServer(grpcServer, module)

	grpc_prometheus.Register(grpcServer)

	return &Grpc{
		grpcServer,
	}
}

func (g *Grpc) GrpcStart(listener net.Listener) {
	err := g.Serve(listener)
	if err != nil {
		panic(err)
	}
}

func (g *Grpc) GrpcShutdown(ctx context.Context) {
	g.GracefulStop()
}
