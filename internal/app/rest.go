package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	internalGrpc "github.com/saufiroja/cqrs/internal/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"net/http"
	"os"
)

type Rest struct {
	port     string
	listener net.Listener
	server   *http.Server
}

func NewRest(port string, listener net.Listener) App {
	return &Rest{
		port:     port,
		listener: listener,
	}
}

func (r *Rest) Start(ctx context.Context) {
	conn, err := grpc.DialContext(
		ctx,
		r.listener.Addr().String(),
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err)
	}

	mux := runtime.NewServeMux()

	err = internalGrpc.RegisterTodosHandler(ctx, mux, conn)
	if err != nil {
		panic(err)
	}

	r.server = &http.Server{
		Addr:    fmt.Sprintf(":%s", r.port),
		Handler: mux,
	}

	err = r.server.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		return
	} else if err != nil {
		fmt.Println("failed to start rest server", err.Error())
		os.Exit(1)
	}
}

func (r *Rest) Shutdown(ctx context.Context) {
	err := r.server.Shutdown(ctx)
	if err != nil {
		return
	}
}
