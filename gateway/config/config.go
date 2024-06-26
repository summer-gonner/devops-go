package config

import (
	"log"
	"strings"
)

// Define a struct to hold service URLs
type Config struct {
	ServiceURLs map[string]string `yaml:"service_urls"`
}

// 根据请求路径获取服务名称
func GetServiceName(path string) string {
	// 这里假设路径为 /service1/path/to/resource，serviceName为service1
	parts := strings.Split(path, "/")
	if len(parts) >= 2 && parts[1] != "" {
		return parts[1]
	}
	return ""
}

// 根据请求路径获取除服务名外的路径部分
func GetPathWithoutServiceName(path string) string {
	// 这里假设路径为 /service1/path/to/resource，返回 /path/to/resource
	parts := strings.SplitN(path, "/", 3) // 仅分割成3部分，以确保只去掉服务名部分
	log.Printf("根据请求路径获取除服务名外的路径部分 %s", path)
	if len(parts) > 2 {
		return "/" + parts[2] // 返回除服务名外的路径部分
	}
	return "/"
}
