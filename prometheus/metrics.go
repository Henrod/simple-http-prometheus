package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	responseTime = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace:   "http",
		Subsystem:   "",
		Name:        "request_duration_seconds",
		Help:        "Time to serve a request on the server",
		ConstLabels: nil,
		Buckets:     prometheus.DefBuckets,
	}, []string{
		PathLabel,
		MethodLabel,
		StatusCodeLabel,
	})

	henrodsTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace:   "henrods",
		Subsystem:   "",
		Name:        "total",
		Help:        "Total henrods counter",
		ConstLabels: nil,
	}, []string{})
)

func IncrementHenrod() {
	henrodsTotal.With(prometheus.Labels{}).Add(1)
}
