package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Henrod/simple-http-prometheus/prometheus"
)

const port = 8081

func main() {
	h := handler{
		slowSleep:   time.Second,
		fastSleep:   100 * time.Millisecond,
		randomRange: 10 * time.Millisecond,
	}

	mux := http.NewServeMux()
	mux.Handle("/slow", prometheus.WithPrometheusMiddleware(h.Slow))
	mux.Handle("/fast", prometheus.WithPrometheusMiddleware(h.Fast))

	log.Printf("Starting server at port %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", port), mux))
}
