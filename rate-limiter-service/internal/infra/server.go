package infra

import (
	"leaky-bucket/internal/application/usecases"
	"leaky-bucket/internal/domain/leakybucket"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Server struct {
	host        string
	port        string
	engine      *gin.Engine
	prometheus  *PrometheusMetrics
	leakyBucket leakybucket.Bucket
}

func NewServer(host string, port string, prometheus *PrometheusMetrics, leakyBucket leakybucket.Bucket) *Server {
	return &Server{
		host:        host,
		port:        port,
		engine:      ginEngine(),
		prometheus:  prometheus,
		leakyBucket: leakyBucket,
	}
}

func (s *Server) Start() error {
	s.registerPrometheus()
	s.registerEndpoints()
	err := s.engine.Run(":3010")

	log.Printf("[🖥️] Server running on %s:%s \n", s.host, s.port)
	return err
}

func (s *Server) registerPrometheus() {
	s.engine.Use(s.prometheus.TrackMetrics())
}

func (s *Server) registerEndpoints() {
	s.engine.GET("/metrics", gin.WrapH(promhttp.Handler()))

	s.engine.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "👽 Rate Limiter Service"})
	})

	s.engine.GET("/no-limiter", usecases.NewNoRateLimiter().Handler())
	s.engine.GET("/leaky-bucket", usecases.NewLeakyBucketRateLimiter(s.leakyBucket).Handler())
}

func ginEngine() *gin.Engine {
	return gin.Default()
}
