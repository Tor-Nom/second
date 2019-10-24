package second

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
)

func RequestIdRe() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.Request.Header.Get("X-Request-Id")

		if requestID == "" {
			uuid4 := uuid.NewV4()
			requestID = uuid4.String()
		}

		c.Set("RequestId", requestID)

		c.Writer.Header().Set("X-Request-Id", requestID)

		fmt.Println("X-Request-Id  ==> ", c.Request.Header.Get("X-Request-Id"))
		c.Next()
	}
}
