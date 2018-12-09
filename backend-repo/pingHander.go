package main

import (
	"jaeger-opentracing-tut/lib/client"
	"net/http"

	opentracing "github.com/opentracing/opentracing-go"
)

func pingHandler(t opentracing.Tracer) func(w http.ResponseWriter, r *http.Request) {
	tracer := t
	return func(w http.ResponseWriter, r *http.Request) {
		_, span := client.ContextFromHTTP(tracer, "pinging", r)
		defer span.Finish()

		w.Write([]byte("pong"))
	}
}
