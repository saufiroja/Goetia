//go:generate mockgen -destination ../../mocks/mock_jaeger.go -package mocks github.com/saufiroja/cqrs/pkg/tracing ITracing
package tracing

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/saufiroja/cqrs/config"
	"github.com/saufiroja/cqrs/pkg/logger"
	"github.com/uber/jaeger-client-go"
	jeagerCfg "github.com/uber/jaeger-client-go/config"
)

type ITracing interface {
	StartSpan(ctx context.Context, operationName string, opts ...opentracing.StartSpanOption) (opentracing.Span, context.Context)
}

type Tracing struct {
}

func NewTracing(conf *config.AppConfig, log *logger.Logger) ITracing {
	cfg := jeagerCfg.Configuration{
		ServiceName: conf.App.ServiceName,
		Sampler: &jeagerCfg.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &jeagerCfg.ReporterConfig{
			CollectorEndpoint: fmt.Sprintf("http://%s:%s/api/traces", conf.Jaeger.Host, conf.Jaeger.Port),
		},
	}

	tracer, closer, err := cfg.NewTracer(
		jeagerCfg.Logger(jaeger.StdLogger),
		jeagerCfg.ZipkinSharedRPCSpan(true),
	)
	defer closer.Close()
	if err != nil {
		log.StartLogger("jaeger.go", "NewTracing").Error("error connecting to jaeger")
		panic(err)
	}

	opentracing.SetGlobalTracer(tracer)

	log.StartLogger("jaeger.go", "NewTracing").Info("connected to jaeger")

	return &Tracing{}
}

func (t *Tracing) StartSpan(ctx context.Context, operationName string, opts ...opentracing.StartSpanOption) (opentracing.Span, context.Context) {
	return opentracing.StartSpanFromContext(ctx, operationName, opts...)
}
