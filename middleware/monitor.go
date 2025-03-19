package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	totalRequests = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total number of HTTP requests.",
	})
	inFlightRequest = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "http_requests_in_flight",
		Help: "Number of currently in-flight HTTP requests.",
	})
	requestDuration = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "http_request_duration_seconds",
		Help:    "Duration of HTTP requests.",
		Buckets: prometheus.DefBuckets,
	})
	requestSize = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "http_request_size_bytes",
		Help:    "Size of HTTP request bodies",
		Buckets: []float64{100, 1000, 10000, 100000},
	})
	responseSize = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "http_response_size_bytes",
		Help:    "Size of HTTP response bodies.",
		Buckets: []float64{100, 1000, 10000, 100000},
	})
)

func init() {
	prometheus.MustRegister(totalRequests)
	prometheus.MustRegister(inFlightRequest)
	prometheus.MustRegister(requestDuration)
	prometheus.MustRegister(requestSize)
	prometheus.MustRegister(responseSize)
}

func PromMiddleweare(c *gin.Context) {
	start := time.Now()
	inFlightRequest.Inc()
	defer inFlightRequest.Dec()

	// Capture request size
	reqSize := float64(c.Request.ContentLength)
	requestSize.Observe(reqSize)
	c.Next()

	duration := time.Since(start)
	totalRequests.Inc()
	requestDuration.Observe(duration.Seconds())

	// Capture response size (if exists)
	if c.Writer.Size() >= 0 {
		responseSize.Observe(float64(c.Writer.Size()))
	}
}
