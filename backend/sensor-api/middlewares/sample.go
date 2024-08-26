package middlewares

import (
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

func SampleMiddleware(c *gin.Context) {
	rand.Seed(time.Now().UnixNano())
	reqId := rand.Intn(1000001)

	// reqId := 10101

	c.Set("reqId", reqId)

	c.Next()
}
