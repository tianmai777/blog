package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/tianmai777/blog/pkg/app"
	"github.com/tianmai777/blog/pkg/errcode"
	"github.com/tianmai777/blog/pkg/limiter"
)

func RateLimiter(l limiter.LimiterIface) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := l.Key(c)
		if bucket, ok := l.GetBucket(key); ok {
			count := bucket.TakeAvailable(1)
			if count == 0 {
				response := app.NewResponse(c)
				response.ToErrorResponse(errcode.TooManyRequests)
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
