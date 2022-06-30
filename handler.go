package main

import (
	"math/rand"
	"net/http"
	"time"
)

type handler struct {
	slowSleep   time.Duration
	fastSleep   time.Duration
	randomRange time.Duration
}

func (h *handler) random() time.Duration {
	return time.Duration(rand.Int63n(int64(h.randomRange))) //nolint:gosec
}

func (h *handler) Slow(w http.ResponseWriter, r *http.Request) {
	time.Sleep(h.slowSleep + h.random())
	w.WriteHeader(http.StatusOK)
}

func (h *handler) Fast(w http.ResponseWriter, _ *http.Request) {
	time.Sleep(h.fastSleep + h.random())
	w.WriteHeader(http.StatusOK)
}
