package main

import (
	"context"
	"fmt"

	"go-micro.dev/v5"
	pb "path/to/your/proto/package" // 替换为你的 protobuf 生成的包路径
)

func main() {
	// 创建新的服务
	service := micro.NewService()

	// 初始化服务
	service.Init()

	// 创建远程服务客户端
	client := pb.NewGreeterService("greeter", service.Client())

	// 构造请求
	request := &pb.Request{
		Name: "Alice",
	}

	// 调用远程服务方法
	response, err := client.Hello(context.Background(), request)
	if err != nil {
		fmt.Println("Error calling service:", err)
		return
	}

	// 打印响应消息
	fmt.Println("Response from service:", response.Msg)
}
