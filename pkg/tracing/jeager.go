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

type Tracing struct {
}

func NewTracing(conf *config.AppConfig, log *logger.Logger) *Tracing {
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

	tracer, _, err := cfg.NewTracer(
		jeagerCfg.Logger(jaeger.StdLogger),
		jeagerCfg.ZipkinSharedRPCSpan(true),
	)
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
