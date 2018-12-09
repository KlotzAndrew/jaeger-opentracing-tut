package main

import (
	"jaeger-opentracing-tut/lib/tracing"
	"log"
	"net/http"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	otlog "github.com/opentracing/opentracing-go/log"
)

func main() {
	tracer, closer := tracing.New("backend")
	defer closer.Close()

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		spanCtx, _ := tracer.Extract(
			opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(r.Header),
		)
		span := tracer.StartSpan("pinging", ext.RPCServerOption(spanCtx))
		defer span.Finish()

		span.LogFields(
			otlog.String("event", "string-format"),
			otlog.String("value", "ping"),
		)
		w.Write([]byte("pong"))
	})

	log.Fatal(http.ListenAndServe(":3001", nil))
}
