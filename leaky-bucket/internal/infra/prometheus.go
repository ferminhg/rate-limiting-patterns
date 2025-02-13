package infra

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

type PrometheusMetrics struct {
}

var (
	RequestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "leaky_bucket_requests_total",
			Help: "Total number of requests processed by the MyApp web server.",
		},
		[]string{"path", "status"},
	)

	ErrorCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "leaky_bucket_requests_errors_total",
			Help: "Total number of error requests processed by the MyApp web server.",
		},
		[]string{"path", "status"},
	)
)

func NewPrometheus() *PrometheusMetrics {
	prometheus.MustRegister(RequestCount)
	prometheus.MustRegister(ErrorCount)

	return &PrometheusMetrics{}
}

func (p *PrometheusMetrics) TrackMetrics() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		c.Next()
		status := c.Writer.Status()
		RequestCount.WithLabelValues(path, http.StatusText(status)).Inc()
		if status >= 400 {
			ErrorCount.WithLabelValues(path, http.StatusText(status)).Inc()
		}
	}
}
