package middleares

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

// 记录访问日志
func Log(res http.ResponseWriter, req *http.Request, traceId string) {
	// 记录请求信息
	logRequestInfo(req, traceId)

	// 处理响应信息
	responseRecorder := NewLoggingResponseWriter(res)
	defer logResponseInfo(responseRecorder)
}

// 记录请求信息
func logRequestInfo(req *http.Request, traceId string) {
	headersString := logRequestHeaders(req.Header)
	body, _ := io.ReadAll(req.Body)

	log.Printf("[%s] traceId --------> [%s] %s  %s 请求头 --------> %s", time.Now().Format("2006-01-02 15:04:05"), traceId, req.Method, fullRequestURL(req), headersString)

	if len(body) > 0 {
		log.Printf("请求体--------------> %s", string(body))
	}

	curlifyRequest(req.Method, req.Header, fullRequestURL(req), body)
}

// 获取完整的请求URL（包括主机名和端口）
func fullRequestURL(req *http.Request) string {
	scheme := "http"
	if req.TLS != nil {
		scheme = "https"
	}
	return fmt.Sprintf("%s://%s%s", scheme, req.Host, req.URL.Path)
}

func logRequestHeaders(headers http.Header) string {
	var headerStrings []string

	for key, values := range headers {
		for _, value := range values {
			headerStrings = append(headerStrings, fmt.Sprintf("%s: %s", key, value))
		}
	}

	return strings.Join(headerStrings, ", ")
}

// 将HTTP请求转换为Curl命令格式
func curlifyRequest(method string, headers http.Header, path string, body []byte) string {
	var curl strings.Builder

	curl.WriteString("curl -X ")
	curl.WriteString(method)
	curl.WriteString(" '")
	curl.WriteString(path)
	curl.WriteString("'")

	headers.Del("Host")
	for key, values := range headers {
		for _, value := range values {
			curl.WriteString(" -H '")
			curl.WriteString(key)
			curl.WriteString(": ")
			curl.WriteString(value)
			curl.WriteString("'")
		}
	}

	if body != nil {
		requestBody := strings.Replace(string(body), "'", "'\\''", -1)

		curl.WriteString(" --data-binary '")
		curl.WriteString(requestBody)
		curl.WriteString("'")
	}

	log.Printf("Curl Command-------------> %s", curl.String())
	return curl.String()
}

// 自定义 ResponseWriter 以捕获响应信息
type LoggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
	body       []byte
}

// 实现 http.ResponseWriter 接口的 Write 方法
func (lrw *LoggingResponseWriter) Write(b []byte) (int, error) {
	if lrw.body == nil {
		lrw.body = make([]byte, 0)
	}
	lrw.body = append(lrw.body, b...)
	return lrw.ResponseWriter.Write(b)
}

// 实现 http.ResponseWriter 接口的 WriteHeader 方法
func (lrw *LoggingResponseWriter) WriteHeader(statusCode int) {
	lrw.statusCode = statusCode
	lrw.ResponseWriter.WriteHeader(statusCode)
}

// 记录响应信息
func logResponseInfo(res *LoggingResponseWriter) {
	log.Printf("响应状态码: %d", res.statusCode)

	if len(res.body) > 0 {
		log.Printf("响应体: %s", string(res.body))
	}
}

// 创建新的 LoggingResponseWriter 实例
func NewLoggingResponseWriter(res http.ResponseWriter) *LoggingResponseWriter {
	return &LoggingResponseWriter{ResponseWriter: res}
}
