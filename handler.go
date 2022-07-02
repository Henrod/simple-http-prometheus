package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Henrod/simple-http-prometheus/prometheus"
)

type handler struct {
	slowSleep time.Duration
	fastSleep time.Duration
}

func (h *handler) Slow(w http.ResponseWriter, r *http.Request) {
	time.Sleep(h.slowSleep)
	w.WriteHeader(http.StatusOK)
}

func (h *handler) Fast(w http.ResponseWriter, _ *http.Request) {
	time.Sleep(h.fastSleep)
	w.WriteHeader(http.StatusOK)
}

func (h *handler) Sleep(w http.ResponseWriter, r *http.Request) {
	sleepDurationStr := r.URL.Query().Get("sleep")
	if sleepDurationStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	sleepDurationInt, err := strconv.Atoi(sleepDurationStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	sleepDuration := time.Duration(sleepDurationInt) * time.Millisecond
	time.Sleep(sleepDuration)

	w.WriteHeader(http.StatusOK)
}

func (h *handler) Henrod(w http.ResponseWriter, _ *http.Request) {
	prometheus.IncrementHenrod()
	w.WriteHeader(http.StatusOK)
}
