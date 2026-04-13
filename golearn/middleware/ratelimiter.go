package middleware

import (
	"net/http"

	"golearn/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func RateLimiter() gin.HandlerFunc {
	limiter := rate.NewLimiter(5, 10) // 5 requests per second, burst of 10

	return func(c *gin.Context) {
		if !limiter.Allow() {
			c.JSON(http.StatusTooManyRequests, models.ErrorResponse{Error: "Rate limit exceeded. Try again later."})
			c.Abort()
			return
		}
		c.Next()
	}
}
