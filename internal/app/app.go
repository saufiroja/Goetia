package app

import (
	"context"
	"fmt"
	"github.com/fatih/color"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/saufiroja/cqrs/config"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type AppFactor struct {
	Grpc *Grpc
	Rest *Rest
}

func NewAppFactor() *AppFactor {
	return &AppFactor{}
}

func (a *AppFactor) StartApp(ctx context.Context) {
	conf := config.NewAppConfig()
	colors := color.New(color.FgGreen)
	reg := prometheus.NewRegistry()

	var module = a.StartModule(conf, reg)

	grpcListen, err := net.Listen("tcp", fmt.Sprintf(":%s", conf.Grpc.Port))
	if err != nil {
		panic(err)
	}

	go func() {
		a.Grpc = NewGrpc(module)
		a.Grpc.GrpcStart(grpcListen)
	}()

	go func() {
		a.Rest = NewRest(conf.Http.Port, grpcListen, ctx, a.Grpc, reg)
		a.Rest.HttpStart()
	}()

	fmt.Printf("%s\n", colors.Sprint("----------------------------------------"))
	fmt.Printf("GRPC server running on port %s\n", colors.Sprint(conf.Grpc.Port))
	fmt.Printf("REST server running on port %s\n", colors.Sprint(conf.Http.Port))
	fmt.Printf("%s\n", colors.Sprint("----------------------------------------"))

	a.StopApp(ctx, colors)
}

func (a *AppFactor) StopApp(ctx context.Context, colors *color.Color) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop

	a.Grpc.GrpcShutdown(ctx)
	a.Rest.HttpShutdown(ctx)

	fmt.Printf("%s\n", colors.Sprint("----------------------------------------"))
	fmt.Println("Server gracefully stopped")
	fmt.Println("Process clean up...")
	fmt.Printf("%s\n", colors.Sprint("----------------------------------------"))
}

func (a *AppFactor) StartModule(conf *config.AppConfig, reg *prometheus.Registry) *Module {
	module := NewModule()
	module.StartModule(conf, reg)

	return module
}
