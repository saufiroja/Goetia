package app

import (
	"context"
	"fmt"
	"github.com/fatih/color"
	"github.com/saufiroja/cqrs/config"
	"github.com/saufiroja/cqrs/pkg/database"
	"github.com/saufiroja/cqrs/pkg/logger"
	"github.com/saufiroja/cqrs/pkg/redis"
	"github.com/saufiroja/cqrs/pkg/tracing"
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
func (a *AppFactor) Start(ctx context.Context) {
	colors := color.New(color.FgCyan).Add(color.Bold)
	conf := config.NewAppConfig()
	log := logger.NewLogger()

	trace := tracing.NewTracing(conf)
	redisCli := redis.NewRedis(conf)
	db := database.NewPostgres(conf)

	module := NewModule(db, log, redisCli, trace)

	grpcListen, err := net.Listen("tcp", fmt.Sprintf(":%s", conf.Grpc.Port))
	if err != nil {
		panic(err)
	}

	go func() {
		a.Grpc = NewGrpc(module)
		a.Grpc.GrpcStart(grpcListen)
	}()

	go func() {
		a.Rest = NewRest(conf.Http.Port, grpcListen, ctx, a.Grpc)
		a.Rest.HttpStart()
	}()

	fmt.Printf("%s\n", colors.Sprint("----------------------------------------"))
	fmt.Printf("GRPC server running on port %s\n", colors.Sprint(conf.Grpc.Port))
	fmt.Printf("REST server running on port %s\n", colors.Sprint(conf.Http.Port))
	fmt.Printf("%s\n", colors.Sprint("----------------------------------------"))

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop

	log.Info("shutting down server")

	a.Grpc.GrpcShutdown(ctx)
	a.Rest.HttpShutdown(ctx)
	redisCli.Close(ctx)
	db.Close(ctx)

	fmt.Printf("%s\n", colors.Sprint("----------------------------------------"))
	fmt.Println("Server gracefully stopped")
	fmt.Println("Process clean up...")
	fmt.Printf("%s\n", colors.Sprint("----------------------------------------"))
}

//func (a *AppFactor) Shutdown(ctx context.Context) {
//	stop := make(chan os.Signal, 1)
//	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
//
//	colors := color.New(color.FgCyan).Add(color.Bold)
//	fmt.Printf("%s\n", colors.Sprint("----------------------------------------"))
//	fmt.Println("Server gracefully stopped")
//	fmt.Println("Process clean up...")
//	fmt.Printf("%s\n", colors.Sprint("----------------------------------------"))
//}
