package config

import "fmt"

type Server struct {
	Port    int    `yaml:"port"`
	AppName string `yaml:"appName"`
}

func (s Server) Addr() string {
	return fmt.Sprintf(":%d", s.Port)
}
