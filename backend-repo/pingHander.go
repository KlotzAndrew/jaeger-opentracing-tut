package main

import (
	"jaeger-opentracing-tut/lib/client"
	"net/http"
)

func pingHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		_, span := client.ContextFromHTTP("pinging", r)
		defer span.Finish()

		w.Write([]byte("pong"))
	}
}
