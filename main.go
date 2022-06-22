package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"
)

type handler struct{
	slowSleep time.Duration
	fastSleep time.Duration
	randomRange time.Duration
}

func (h *handler) random() time.Duration {
	return time.Duration(rand.Int63n(int64(h.randomRange)))
}

func (h *handler) Slow(w http.ResponseWriter, _ *http.Request) {
	time.Sleep(h.slowSleep + h.random())
	w.WriteHeader(http.StatusOK)
}

func (h *handler) Fast(w http.ResponseWriter, _ *http.Request) {
	time.Sleep(h.fastSleep + h.random())
	w.WriteHeader(http.StatusOK)
}

func main() {
	h := handler{
		slowSleep: time.Second,
		fastSleep: 100*time.Millisecond,
		randomRange:    10*time.Millisecond,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/slow", h.Slow)
	mux.HandleFunc("/fast", h.Fast)

	log.Fatal(http.ListenAndServe(":8081", mux))
}