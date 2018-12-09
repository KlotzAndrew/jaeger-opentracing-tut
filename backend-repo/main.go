package main

import (
	"jaeger-opentracing-tut/lib/tracing"
	"log"
	"net/http"
)

func main() {
	tracer, closer := tracing.New("backend-repo")
	defer closer.Close()

	http.HandleFunc("/ping", pingHandler(tracer))

	log.Fatal(http.ListenAndServe(":3002", nil))
}
