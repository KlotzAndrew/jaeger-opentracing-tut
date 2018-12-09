package main

import (
	"jaeger-opentracing-tut/lib/tracing"
	"log"
	"net/http"

	opentracing "github.com/opentracing/opentracing-go"
)

func main() {
	tracer, closer := tracing.New("backend")
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	http.HandleFunc("/ping", pingHandler())

	log.Fatal(http.ListenAndServe(":3001", nil))
}
