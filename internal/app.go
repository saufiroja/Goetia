package internal

import (
	"context"
	"fmt"
	"github.com/fatih/color"
	"github.com/saufiroja/cqrs/config"
	"github.com/saufiroja/cqrs/internal/app"
	"github.com/saufiroja/cqrs/pkg/database"
	"github.com/saufiroja/cqrs/pkg/logger"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func Start() {
	colors := color.New(color.FgCyan).Add(color.Bold)
	conf := config.NewAppConfig()

	// database
	db := database.NewPostgres(conf)
	log := logger.NewLogger()

	module := NewModule(db, log)

	grpcListen, err := net.Listen("tcp", fmt.Sprintf(":%s", conf.Grpc.Port))
	if err != nil {
		panic(err)
	}

	grpcApp := app.NewGrpc(conf.Grpc.Port, grpcListen, module)
	restApp := app.NewRest(conf.Http.Port, grpcListen)

	// start grpc server
	go grpcApp.Start(context.Background())

	// start rest server
	go restApp.Start(context.Background())

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

	log.Info("server gracefully stopped")
	log.Info("process clean up...")
}
