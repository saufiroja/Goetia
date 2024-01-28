package app

import (
	"context"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/saufiroja/cqrs/internal/delivery/controllers"
	internalGrpc "github.com/saufiroja/cqrs/internal/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type Grpc struct {
	*grpc.Server
}

func NewGrpc(dependencies controllers.ITodoController) *Grpc {
	grpcServer := grpc.NewServer(
		grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
		grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
	)
	reflection.Register(grpcServer)

	internalGrpc.RegisterTodosServer(grpcServer, dependencies)

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
