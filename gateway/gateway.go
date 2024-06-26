package main

import (
	"bytes"
	"devops-go/gateway/config"
	"fmt"
	"github.com/go-yaml/yaml"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"time"
)

// Gateway 结构体定义
type Gateway struct {
	proxies map[string]*httputil.ReverseProxy
}

func main() {
	// 打开并读取配置文件
	file, err := os.Open("gateway/config.yaml")
	if err != nil {
		log.Fatalf("Failed to open config file: %v", err)
	}
	defer file.Close()

	// 解析配置文件
	var cfg config.Config
	if err := yaml.NewDecoder(file).Decode(&cfg); err != nil {
		log.Fatalf("Failed to parse config file: %v", err)
	}

	// 创建多个 ReverseProxy
	proxies := make(map[string]*httputil.ReverseProxy)
	for name, targetURL := range cfg.ServiceURLs {
		target, err := url.Parse(targetURL)
		if err != nil {
			panic(fmt.Sprintf("Failed to parse target URL for %s: %v", name, err))
		}

		// 创建 ReverseProxy 并设置 Director 和 ModifyResponse 方法
		proxy := &httputil.ReverseProxy{
			Director: func(req *http.Request) {
				//traceId := utils.GenerateTraceID()
				//middleares.LogRequestInfo(req, traceId)
				// 设置目标URL
				req.URL.Scheme = target.Scheme
				req.URL.Host = target.Host

				// 去掉请求路径中的服务名部分
				req.URL.Path = config.GetPathWithoutServiceName(req.URL.Path)

				// 设置其他需要转发的头信息，例如 Authorization
				//req.Header.Set("Authorization", req.Header.Get("Authorization"))

			},
			ModifyResponse: func(response *http.Response) error {
				//middleares.LogResponseInfo(response)
				// 读取响应体
				body, err := io.ReadAll(response.Body)
				if err != nil {
					log.Printf("Failed to read response body: %v", err)
					return err
				}
				// 打印或处理响应体
				log.Printf("Received response: %s", string(body))

				// 将读取的响应体重新设置给 Response.Body
				response.Body = ioutil.NopCloser(bytes.NewReader(body))

				return nil
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

		proxies[name] = proxy
	}

	// 创建 Gateway 实例
	gateway := &Gateway{
		proxies: proxies,
	}

	// 设置 HTTP 服务端口和处理程序
	port := ":8080"
	log.Printf("Starting gateway on %s...\n", port)

	// 启动 HTTP 服务
	log.Fatal(http.ListenAndServe(port, gateway))
}

// Gateway 结构体的 ServeHTTP 方法
func (g *Gateway) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	//// 记录访问日志
	//traceID := utils.GenerateTraceID() // 假设有一个生成 traceID 的函数
	//middleares.Log(res, req, traceID)

	// 根据请求路径选择后端服务进行转发
	serviceName := config.GetServiceName(req.URL.Path)
	proxy, ok := g.proxies[serviceName]
	if !ok {
		http.Error(res, "Service not found", http.StatusNotFound)
		return
	}

	// 使用 ReverseProxy 将请求转发到目标URL
	proxy.ServeHTTP(res, req)
}
