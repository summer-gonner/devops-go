package middleares

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

// 记录请求信息
func LogRequestInfo(req *http.Request, traceId string) {
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

// LoggingResponse 自定义结构体，用于捕获响应信息
type LoggingResponse struct {
	*http.Response        // Embedding http.Response to inherit its methods and fields
	Body           []byte // Custom field to store response body
}

func LogResponseInfo(res *http.Response) {
	log.Printf("响应状态码: %d", res.StatusCode)

	// 读取并打印响应体内容
	if res.Body != nil {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Printf("Failed to read response body: %v", err)
		}
		log.Printf("响应体: %s", string(body))
	}
}
