package middleares

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// RequestLoggerMiddleware 中间件：记录请求信息
func RequestLoggerMiddleware(traceId string) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// 复制请求体
		var requestBody []byte
		if c.Request.Body != nil {
			requestBody, _ = ioutil.ReadAll(c.Request.Body)
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestBody))
		}

		// 执行下一个处理函数
		c.Next()

		// 在响应之后记录请求信息
		latency := time.Since(start)

		// 打印请求信息
		fmt.Printf("[%s] traceId:[%s] %s %s %v\n", time.Now().Format("2006-01-02 15:04:05"), traceId, c.Request.Method, c.Request.URL.Path, latency)
		fmt.Printf("请求头------------------->")
		for key, values := range c.Request.Header {
			for _, value := range values {
				fmt.Printf("\t%s: %s\n", key, value)
			}
		}

		// 打印请求体
		fmt.Printf("请求体-------------------> %s\n", string(requestBody))
	}
}

func LogResponse(response *http.Response, traceId string) error {
	// 读取响应体
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("[%s] Failed to read response body: %v", traceId, err)
		return err
	}
	// 打印或处理响应体
	log.Printf(" traceId:[%s] 响应体-------------------> %s", traceId, string(body))

	// 将读取的响应体重新设置给 Response.Body
	response.Body = ioutil.NopCloser(bytes.NewReader(body))

	return nil
}
