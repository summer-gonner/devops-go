package config

import (
	"net/http"
	"net/http/httputil"
	"strings"
)

// Define a struct to hold service URLs
type Config struct {
	ServiceURLs map[string]string `yaml:"service_urls"`
}

// Gateway struct to store reverse proxies
type Gateway struct {
	proxies map[string]*httputil.ReverseProxy
}

// ServeHTTP method to implement the Handler interface
func (g *Gateway) ServeHTTP(res http.ResponseWriter, req *http.Request) {
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
