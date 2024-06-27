package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"strings"
)

// Define a struct to hold service URLs
type Config struct {
	ServiceURLs map[string]string `yaml:"service_urls"`
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

// loadConfig 加载配置文件
func LoadConfig(filename string, cfg *Config) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open config file: %v", err)
	}
	defer file.Close()

	if err := yaml.NewDecoder(file).Decode(cfg); err != nil {
		return fmt.Errorf("failed to parse config file: %v", err)
	}

	return nil
}

// GetServiceName 从请求路径中提取服务名
func GetServiceName(path string) string {
	parts := strings.Split(strings.Trim(path, "/"), "/")
	if len(parts) > 0 {
		log.Printf("parts-----> %s", parts)
		return parts[0]
	}
	return ""
}
