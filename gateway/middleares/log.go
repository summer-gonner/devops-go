package middleares

import (
	"bytes"
	"devops-go/gateway/utils"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceId := utils.GenerateTraceID()
		start := time.Now()

		// Make a copy of the request body
		var requestBodyBytes []byte
		if c.Request.Body != nil {
			requestBodyBytes, _ = ioutil.ReadAll(c.Request.Body)
		}
		// Restore the io.ReadCloser to its original state
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestBodyBytes))

		// Process request
		c.Next()

		// Log request body
		if len(requestBodyBytes) > 0 {
			// Log request details to console
			log.Printf("请求体-----------------> traceId[%s][%s] %s %s - %s %s",
				traceId,
				c.Request.Method,
				c.Request.URL.Scheme,
				c.Request.URL.Host,
				time.Since(start),
				string(requestBodyBytes),
			)
		}

		// Log response body
		c.Writer.WriteHeaderNow() // Ensure response headers are written
		if len(c.Errors) == 0 {   // If no errors were recorded
			log.Printf("Response Body: %s", c.Writer)
		}
	}
}
