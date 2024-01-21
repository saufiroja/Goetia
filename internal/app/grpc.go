package app

import (
	"context"
	"github.com/saufiroja/cqrs/internal/delivery/controllers"
	internalGrpc "github.com/saufiroja/cqrs/internal/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type Grpc struct {
	port         string
	listener     net.Listener
	dependencies controllers.ITodoController
	server       *grpc.Server
}

func NewGrpc(port string, listener net.Listener, dependencies controllers.ITodoController) App {
	return &Grpc{
		port:         port,
		listener:     listener,
		dependencies: dependencies,
	}
}

func (g *Grpc) Start(ctx context.Context) {
	var opt []grpc.ServerOption
	g.server = grpc.NewServer(opt...)
	reflection.Register(g.server)

	internalGrpc.RegisterTodosServer(g.server, g.dependencies)

	err := g.server.Serve(g.listener)
	if err != nil {
		panic(err)
	}
}

func (g *Grpc) Shutdown(ctx context.Context) {
	g.server.GracefulStop()
}
