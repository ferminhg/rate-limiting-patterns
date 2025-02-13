package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	RequestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "myapp_requests_total",
			Help: "Total number of requests processed by the MyApp web server.",
		},
		[]string{"path", "status"},
	)

	ErrorCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "myapp_requests_errors_total",
			Help: "Total number of error requests processed by the MyApp web server.",
		},
		[]string{"path", "status"},
	)
)

func PrometheusInit() {
	prometheus.MustRegister(RequestCount)
	prometheus.MustRegister(ErrorCount)
}

func TrackMetrics() gin.HandlerFunc {
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

var (
	internalServerError = "Internal Server Error"
	notFound            = "Not Found"
)

func main() {
	r := gin.Default()

	PrometheusInit()

	// Prometheus metrics endpoint
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Middleware to track request metrics
	r.Use(TrackMetrics())

	// A simple route that increments the request count
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, Prometheus!")
	})

	// Another example route
	r.GET("/get-user", func(c *gin.Context) {
		param := c.DefaultQuery("param", "") // Get the query parameter "param" with a default empty value

		if param == "error" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": internalServerError,
			})
			return
		}

		if param == "not-found" {
			c.JSON(http.StatusNotFound, gin.H{
				"message": notFound,
			})
			return
		}

		c.String(http.StatusOK, "Success Get Users")
	})

	r.GET("/get-role", func(c *gin.Context) {
		param := c.DefaultQuery("param", "") // Get the query parameter "param" with a default empty value

		if param == "error" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": internalServerError,
			})
			return
		}

		if param == "not-found" {
			c.JSON(http.StatusNotFound, gin.H{
				"message": notFound,
			})
			return
		}

		c.String(http.StatusOK, "Success Get Roles")
	})

	r.GET("/get-level", func(c *gin.Context) {
		param := c.DefaultQuery("param", "") // Get the query parameter "param" with a default empty value

		if param == "error" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": internalServerError,
			})
			return
		}

		if param == "not-found" {
			c.JSON(http.StatusNotFound, gin.H{
				"message": notFound,
			})
			return
		}

		c.String(http.StatusOK, "Success Get Levels")
	})

	// Start the Gin server
	r.Run(":3010")
}
