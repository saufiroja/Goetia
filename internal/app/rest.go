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
	*http.Server
}

func NewRest(port string, listener net.Listener, ctx context.Context) *Rest {
	conn, err := grpc.DialContext(
		ctx,
		listener.Addr().String(),
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

	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: mux,
	}

	return &Rest{
		httpServer,
	}
}

func (r *Rest) HttpStart() {
	err := r.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		return
	} else if err != nil {
		fmt.Println("failed to start rest server", err.Error())
		os.Exit(1)
	}
}

func (r *Rest) HttpShutdown(ctx context.Context) {
	err := r.Shutdown(ctx)
	if err != nil {
		return
	}
}
