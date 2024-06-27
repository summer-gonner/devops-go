package middleares

import (
	"net/http/httputil"
	"net/url"
	"sync"
)

type LoadBalancer struct {
	backendURLs []string
	index       int
	mu          sync.Mutex
}

func NewLoadBalancer(urls []string) *LoadBalancer {
	return &LoadBalancer{
		backendURLs: urls,
		index:       0,
	}
}

// SingleHostReverseProxy 创建一个反向代理
func NewSingleHostReverseProxy(target string) *httputil.ReverseProxy {
	url, _ := url.Parse(target)
	return httputil.NewSingleHostReverseProxy(url)
}
