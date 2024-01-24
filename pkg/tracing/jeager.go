package tracing

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/saufiroja/cqrs/config"
	"github.com/uber/jaeger-client-go"
	jeagerCfg "github.com/uber/jaeger-client-go/config"
	"log"
)

type Tracing struct {
}

func NewTracing(conf *config.AppConfig) *Tracing {
	cfg := jeagerCfg.Configuration{
		ServiceName: "todo-service",
		Sampler: &jeagerCfg.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &jeagerCfg.ReporterConfig{
			LogSpans:          true,
			CollectorEndpoint: fmt.Sprintf("http://%s:%s/api/traces", conf.Jaeger.Host, conf.Jaeger.Port),
		},
	}

	tracer, _, err := cfg.NewTracer(
		jeagerCfg.Logger(jaeger.StdLogger),
		jeagerCfg.ZipkinSharedRPCSpan(true),
	)
	if err != nil {
		panic(err)
	}

	opentracing.SetGlobalTracer(tracer)

	log.Println("Jaeger Tracing is enabled")

	return &Tracing{}
}

func (t *Tracing) StartSpan(ctx context.Context, operationName string) (opentracing.Span, context.Context) {
	return opentracing.StartSpanFromContext(ctx, operationName)
}
