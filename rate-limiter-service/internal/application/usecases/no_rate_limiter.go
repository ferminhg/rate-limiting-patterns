package usecases

import (
	"math/rand"

	"github.com/gin-gonic/gin"
)

const (
	maxRandom = 5000
)

type NoRateLimiter struct {
}

func NewNoRateLimiter() *NoRateLimiter {
	return &NoRateLimiter{}
}

func (n *NoRateLimiter) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		random := rand.Intn(maxRandom)
		c.JSON(200, gin.H{"message": "[🍮 No rate limiter]", "random": random})
	}
}
