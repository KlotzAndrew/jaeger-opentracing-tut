package main

import (
	"context"
	client "jaeger-opentracing-tut/lib/http"
	"jaeger-opentracing-tut/lib/tracing"

	opentracing "github.com/opentracing/opentracing-go"
)

func main() {
	tracer, closer := tracing.New("frontend!")
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	span := tracer.StartSpan("hello-function")
	span.SetTag("hello-to", "yee")
	defer span.Finish()

	ctx := context.Background()
	ctx = opentracing.ContextWithSpan(ctx, span)

	client.PingService(ctx, "http://0.0.0.0:3001/ping", "str")
}
