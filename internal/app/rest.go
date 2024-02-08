package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	internalGrpc "github.com/saufiroja/cqrs/internal/grpc"
	metrics "github.com/slok/go-http-metrics/metrics/prometheus"
	"github.com/slok/go-http-metrics/middleware"
	"github.com/slok/go-http-metrics/middleware/std"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"net/http"
	"os"
	"strings"
)

type Rest struct {
	*http.Server
}

func NewRest(port string, listener net.Listener, ctx context.Context, grpcServer *Grpc, reg *prometheus.Registry) *Rest {
	conn, err := grpc.DialContext(
		ctx,
		listener.Addr().String(),
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err)
	}

	gatewayMux := runtime.NewServeMux()

	err = internalGrpc.RegisterTodosHandler(ctx, gatewayMux, conn)
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))

	// mux any http request to grpc server
	mux.Handle("/", gatewayMux)

	mdlw := middleware.New(middleware.Config{
		Recorder:           metrics.NewRecorder(metrics.Config{}),
		DisableMeasureSize: true,
	})
	h := std.Handler("", mdlw, mux)
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
			return
		}
		h.ServeHTTP(w, r)
	})

	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: h2c.NewHandler(handler, &http2.Server{}),
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
