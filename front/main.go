package main

import (
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

func main() {
	if len(os.Args) != 2 {
		panic("ERROR: Expecting one argument")
	}

	tracer, closer := setupTracer("frontend!")
	defer closer.Close()

	ht := os.Args[1]

	span := tracer.StartSpan("hello-function")
	defer span.Finish()

	span.SetTag("hello-to", ht)

	fullStr := fmt.Sprintf("Full-hello %s!", ht)
	span.LogFields(
		log.String("event", "string-formmated"),
		log.String("value", fullStr),
	)

	fmt.Println("fullStr")
	span.LogKV("event", "println")
}
