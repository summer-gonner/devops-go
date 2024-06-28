package main

import (
	"devops-go/gateway/config"
	"devops-go/gateway/loadbalance/hash"
	"devops-go/gateway/middleares"
	"github.com/gin-gonic/gin"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func main() {
	loadBalancer := hash.NewConsistenceHashBalance(10, nil)

	// 初始化 Gin 引擎
	router := gin.New()
	// 注册全局中间件，模拟过滤器链的效果
	router.Use(middleares.AuthorizationMiddleware)
	router.Use(middleares.RequestModifierMiddleware)
	router.Use(middleares.Logger())

	// 加载配置文件
	var cfg config.Config
	if err := config.LoadConfig("gateway/config.yaml", &cfg); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 创建多个 ReverseProxy，并进行负载均衡
	proxies := make(map[string]*httputil.ReverseProxy)
	for _, route := range cfg.Gateway.Routes {
		targetURL := route.Target
		err := loadBalancer.Add(targetURL)
		if err != nil {
			log.Printf("Failed to add target URL %s to load balancer: %v", targetURL, err)
			continue
		}
		// 解析路径规则，去掉 Path= 前缀，获取真实的路径部分
		var path string
		for _, predicate := range route.Predicates {
			if strings.HasPrefix(predicate, "Path=") {
				path = strings.TrimPrefix(predicate, "Path=")
				break
			}
		}
		target, err := url.Parse(targetURL)
		if err != nil {
			log.Printf("Failed to parse target URL %s: %v", targetURL, err)
			continue
		} else {
			log.Printf("%s", target)
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

		proxies[path] = proxy // 假设每个路由只有一个 predicate，这里简化处理

	}
	// 创建 Gateway 实例
	gateway := &Gateway{
		proxies: proxies,
	}

	// 设置路由处理函数
	router.Any("/*path", gateway.ServeHTTP)
	// 从配置中读取端口号并启动 HTTP 服务
	port := ":" + strconv.Itoa(cfg.Server.Port)
	log.Fatal(router.Run(port))
}

// Gateway 结构体定义
type Gateway struct {
	proxies map[string]*httputil.ReverseProxy
}

func (g *Gateway) ServeHTTP(c *gin.Context) {

	var proxy *httputil.ReverseProxy
	for predicate, routeProxy := range g.proxies {
		if strings.HasPrefix(c.Request.URL.Path, predicate) {
			proxy = routeProxy
			break
		}
	}

	if proxy == nil {
		log.Printf("No route found for path: %s", c.Request.URL.Path)
		c.String(http.StatusNotFound, "Service not found")
		return
	}

	proxy.ServeHTTP(c.Writer, c.Request)
}
