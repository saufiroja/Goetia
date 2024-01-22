package internal

import (
	"context"
	"fmt"
	"github.com/fatih/color"
	"github.com/saufiroja/cqrs/config"
	"github.com/saufiroja/cqrs/internal/app"
	"github.com/saufiroja/cqrs/pkg/database"
	"github.com/saufiroja/cqrs/pkg/logger"
	"github.com/saufiroja/cqrs/pkg/redis"
	"github.com/saufiroja/cqrs/pkg/tracing"
	"go.opentelemetry.io/otel"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func Start() {
	colors := color.New(color.FgCyan).Add(color.Bold)
	conf := config.NewAppConfig()

	// tracing
	trace, err := tracing.NewTracing("http://localhost:14268/api/traces")
	if err != nil {
		panic(err)
	}

	otel.SetTracerProvider(trace.TracerProvider)

	// redis
	redisCli := redis.NewRedis(conf.Redis.Host, conf.Redis.Port)

	// database
	db := database.NewPostgres(conf)
	log := logger.NewLogger()

	module := NewModule(db, log, redisCli, trace)

	grpcListen, err := net.Listen("tcp", fmt.Sprintf(":%s", conf.Grpc.Port))
	if err != nil {
		panic(err)
	}

	grpcApp := app.NewGrpc(conf.Grpc.Port, grpcListen, module)
	restApp := app.NewRest(conf.Http.Port, grpcListen)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// start grpc server
	go grpcApp.Start(ctx)

	// start rest server
	go restApp.Start(ctx)

	fmt.Printf("%s\n", colors.Sprint("----------------------------------------"))
	fmt.Printf("GRPC server running on port %s\n", colors.Sprint(conf.Grpc.Port))
	fmt.Printf("REST server running on port %s\n", colors.Sprint(conf.Http.Port))
	fmt.Printf("%s\n", colors.Sprint("----------------------------------------"))

	// wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop

	log.Info("shutting down server")

	grpcApp.Shutdown(context.Background())
	restApp.Shutdown(context.Background())
	trace.Shutdown(context.Background())
	redisCli.Close(context.Background())
	db.Close(context.Background())

	log.Info("server gracefully stopped")
	log.Info("process clean up...")
}
