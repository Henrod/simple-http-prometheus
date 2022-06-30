package prometheus

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Middleware struct {
	next http.HandlerFunc
}

const (
	PathLabel       = "path"
	MethodLabel     = "method"
	StatusCodeLabel = "status_code"

	metricsPort = 2112
)

var (
	responseTime = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace:   "henrod",
		Subsystem:   "simple_http_prometheus",
		Name:        "http_request_duration_seconds",
		Help:        "",
		ConstLabels: nil,
		Buckets:     nil,
	}, []string{
		PathLabel,
		MethodLabel,
		StatusCodeLabel,
	})

	once sync.Once
)

func WithPrometheusMiddleware(next http.HandlerFunc) *Middleware {
	once.Do(func() {
		go func() {
			log.Printf("Starting metrics server at port %d", metricsPort)
			http.Handle("/metrics", promhttp.Handler())
			log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", metricsPort), nil))
		}()
	})

	return &Middleware{next: next}
}

func (p *Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	responseWriter := NewResponseWriter(w)
	p.next.ServeHTTP(responseWriter, r)
	responseTime.With(prometheus.Labels{
		PathLabel:       r.URL.Path,
		MethodLabel:     r.Method,
		StatusCodeLabel: strconv.Itoa(responseWriter.statusCode),
	}).Observe(time.Since(start).Seconds())
}
