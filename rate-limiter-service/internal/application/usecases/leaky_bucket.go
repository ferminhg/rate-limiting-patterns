package usecases

import (
	"math/rand"
	"time"

	"leaky-bucket/internal/domain/leakybucket"

	"github.com/gin-gonic/gin"
)

type LeakyBucketRateLimiter struct {
	bucket leakybucket.Bucket
}

func NewLeakyBucketRateLimiter(bucket leakybucket.Bucket) *LeakyBucketRateLimiter {
	return &LeakyBucketRateLimiter{
		bucket: bucket,
	}
}

func (l *LeakyBucketRateLimiter) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		status, random := l.Execute()
		c.JSON(status, gin.H{"message": "[🍮 Leaky bucket rate limiter]", "random": random})
	}
}

func (l *LeakyBucketRateLimiter) Execute() (int, int) {
	if l.bucket.IsFull() {
		return 429, 0
	}

	l.bucket.Inc()
	time.Sleep(1 * time.Second)
	defer l.bucket.Dec()
	return 200, rand.Intn(5000)
}
