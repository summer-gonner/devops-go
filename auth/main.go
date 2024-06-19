package main

import (
	"context"
	"fmt"

	"go-micro.dev/v5"
	pb "path/to/your/proto/package" // 替换为你的 protobuf 生成的包路径
)

type greeter struct{}

func (g *greeter) Hello(ctx context.Context, req *pb.Request, rsp *pb.Response) error {
	rsp.Msg = "Hello, " + req.Name
	return nil
}

func main() {
	// 创建新的服务
	service := micro.NewService(
		micro.Name("greeter"),
	)

	// 初始化服务
	service.Init()

	// 注册处理程序
	pb.RegisterGreeterHandler(service.Server(), new(greeter))

	// 运行服务
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
