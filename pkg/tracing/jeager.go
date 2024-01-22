package tracing

import (
	"context"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
	"log"
)

type Tracing struct {
	TracerProvider *tracesdk.TracerProvider
}

func NewTracing(url string) (*Tracing, error) {
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, err
	}
	tp := tracesdk.NewTracerProvider(
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exp),
		// Record information about this application in a Resource.
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("todos"),
			//attribute.String("environment", environment),
			attribute.Int64("ID", 42),
		)),
	)
	log.Println("Tracing initialized")
	return &Tracing{TracerProvider: tp}, nil
}

func (t *Tracing) Shutdown(ctx context.Context) {
	err := t.TracerProvider.Shutdown(ctx)
	if err != nil {
		panic(err)
	}
}

func (t *Tracing) StartGlobalTracerSpan(ctx context.Context, name string) (context.Context, trace.Span) {
	return t.TracerProvider.Tracer("todos").Start(ctx, name)
}
