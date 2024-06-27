package main

import (
	"devops-go/gateway/config"
	"devops-go/gateway/middleares"
	"devops-go/gateway/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

func main() {
	//loadBalancer := loadbalance.NewConsistenceHashBalance(10, nil)
	traceId := utils.GenerateTraceID()

	// 初始化 Gin 引擎
	router := gin.Default()

	// 注册全局中间件，模拟过滤器链的效果
	router.Use(middleares.AuthorizationMiddleware)
	router.Use(middleares.RequestLoggerMiddleware(traceId))
	router.Use(middleares.RequestModifierMiddleware)

	// 加载配置文件
	var cfg config.Config
	if err := config.LoadConfig("gateway/config.yaml", &cfg); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 创建多个 ReverseProxy，并进行负载均衡
	proxies := make(map[string]*httputil.ReverseProxy)
	for _, targetURL := range cfg.ServiceURLs {
		//err := loadBalancer.Add(targetURL)
		target, err := url.Parse(targetURL)
		if err != nil {
			log.Printf("Failed to add target URL %s to load balancer: %v", targetURL, err)
			continue // 或者进行错误处理
		}

		proxy := &httputil.ReverseProxy{
			Director: func(req *http.Request) {
				// 设置目标URL
				req.URL.Scheme = target.Scheme
				req.URL.Host = target.Host

				// 去掉请求路径中的服务名部分
				req.URL.Path = config.GetPathWithoutServiceName(req.URL.Path)

				// 设置其他需要转发的头信息，例如 Authorization
				// req.Header.Set("Authorization", req.Header.Get("Authorization"))
			},
			ModifyResponse: func(response *http.Response) error {
				return middleares.LogResponse(response, traceId)
			},
			Transport: &http.Transport{
				ResponseHeaderTimeout: time.Second * 60, // 设置响应头超时时间
				DialContext: (&net.Dialer{
					Timeout:   30 * time.Second, // 设置连接超时时间
					KeepAlive: 30 * time.Second, // 设置长连接保持时间
				}).DialContext,
				MaxIdleConns:          100,              // 最大空闲连接数
				IdleConnTimeout:       90 * time.Second, // 空闲连接的超时时间
				TLSHandshakeTimeout:   10 * time.Second, // TLS 握手超时时间
				ExpectContinueTimeout: 1 * time.Second,  // 期望继续超时时间
			},
		}

		proxies[targetURL] = proxy
	}

	// 创建 Gateway 实例
	gateway := &Gateway{
		proxies: proxies,
	}

	// 设置路由处理函数
	router.Any("/*path", gateway.ServeHTTP)

	// 启动 HTTP 服务
	port := ":8080"
	log.Printf("Starting gateway on %s...\n", port)
	log.Fatal(router.Run(port))
}

// Gateway 结构体定义
type Gateway struct {
	proxies map[string]*httputil.ReverseProxy
}

func (g *Gateway) ServeHTTP(c *gin.Context) {
	// 根据请求路径选择后端服务进行转发
	serviceName := config.GetServiceName(c.Request.URL.Path)
	log.Printf("Service name extracted from URL path: %s", serviceName)

	// 查找对应的代理
	proxy, ok := g.proxies[serviceName]
	if !ok {
		log.Printf("Service not found for serviceName: %s", serviceName)
		c.String(http.StatusNotFound, "Service not found")
		return
	}

	// 使用 ReverseProxy 将请求转发到目标URL
	proxy.ServeHTTP(c.Writer, c.Request)
}
