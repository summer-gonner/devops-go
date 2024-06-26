package main

import (
	"devops-go/gateway/config"
	"devops-go/gateway/middleares"
	"devops-go/gateway/utils"
	"fmt"
	"github.com/didip/tollbooth"
	_ "github.com/didip/tollbooth/limiter"
	"github.com/go-yaml/yaml"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
)

// 定义一个结构体来存储代理的目标URL
type Gateway struct {
	proxies map[string]*httputil.ReverseProxy
}

// 实现ServeHTTP方法，使Gateway成为一个Handler
func (g *Gateway) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	// 限流中间件
	limiter := tollbooth.NewLimiter(1, nil) // 每秒限制1个请求
	httpError := tollbooth.LimitByRequest(limiter, res, req)
	if httpError != nil {
		return
	}
	// 记录访问日志
	traceID := utils.GenerateTraceID() // 假设有一个生成 traceID 的函数
	middleares.Log(res, req, traceID)

	// 根据请求路径选择后端服务进行转发
	serviceName := getServiceName(req.URL.Path)
	proxy, ok := g.proxies[serviceName]
	if !ok {
		http.Error(res, "Service not found", http.StatusNotFound)
		return
	}

	// 使用ReverseProxy将请求转发到目标URL
	proxy.ServeHTTP(res, req)
}

// 根据请求路径获取服务名称
func getServiceName(path string) string {
	// 这里假设路径为 /service1/path/to/resource，serviceName为service1
	parts := strings.Split(path, "/")
	if len(parts) >= 2 && parts[1] != "" {
		return parts[1]
	}
	return ""
}

func main() {
	// Open and read the configuration file
	file, err := os.Open("gateway/config.yaml")
	if err != nil {
		log.Fatalf("Failed to open config file: %v", err)
	}
	defer file.Close()

	// Parse the YAML file into a Config struct
	var config config.Config
	if err := yaml.NewDecoder(file).Decode(&config); err != nil {
		log.Fatalf("Failed to parse config file: %v", err)
	}

	// 创建多个ReverseProxy
	proxies := make(map[string]*httputil.ReverseProxy)
	for name, targetURL := range config.ServiceURLs {
		target, err := url.Parse(targetURL)
		if err != nil {
			panic(fmt.Sprintf("Failed to parse target URL for %s: %v", name, err))
		}
		proxies[name] = httputil.NewSingleHostReverseProxy(target)
	}

	// 创建一个Gateway实例
	gateway := &Gateway{
		proxies: proxies,
	}

	// 设置HTTP服务端口和处理程序
	port := ":8080"
	log.Printf("Starting gateway on %s...\n", port)
	http.ListenAndServe(port, gateway)
}
