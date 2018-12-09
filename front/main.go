package main

import (
	"context"
	"fmt"
	"io"
	"os"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	jaeger "github.com/uber/jaeger-client-go"
	config "github.com/uber/jaeger-client-go/config"
)

func setupTracer(service string) (opentracing.Tracer, io.Closer) {
	cfg := &config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans: true,
		},
	}

	tracer, closer, err := cfg.New(service, config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(err)
	}
	return tracer, closer
}

func formatString(ctx context.Context, str string) string {
	span, _ := opentracing.StartSpanFromContext(ctx, "format-string")
	defer span.Finish()

	fullStr := fmt.Sprintf("Full-hello %s!", str)
	span.LogFields(
		log.String("event", "string-formmated"),
		log.String("value", fullStr),
	)

	return fullStr
}

func printString(ctx context.Context, str string) {
	span, _ := opentracing.StartSpanFromContext(ctx, "printHello")
	defer span.Finish()

	fmt.Println(str)
	span.LogKV("event", "printlnn")
}

func main() {
	if len(os.Args) != 2 {
		panic("ERROR: Expecting one argument")
	}

	tracer, closer := setupTracer("frontend!")
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	ht := os.Args[1]

	span := tracer.StartSpan("hello-function")
	span.SetTag("hello-to", ht)
	defer span.Finish()

	ctx := context.Background()
	ctx = opentracing.ContextWithSpan(ctx, span)

	str := formatString(ctx, ht)
	printString(ctx, str)
}
