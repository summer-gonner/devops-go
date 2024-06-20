package svc

import (
	"context"
	"devops-go/api/internal/config"
	"devops-go/rpc/sys/sys"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type ServiceContext struct {
	Config config.Config
	Sys    sys.Sys
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Sys:    sys.NewSys(zrpc.MustNewClient(c.SysRpc, zrpc.WithUnaryClientInterceptor(interceptor))),
	}
}

func interceptor(ctx context.Context, method string, req any, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	md := metadata.New(map[string]string{"x": "xx"})
	ctx = metadata.NewOutgoingContext(ctx, md)
	// logx.Debug("调用rpc服务前")
	err := invoker(ctx, method, req, reply, cc)
	if err != nil {
		return err
	}
	// logx.Debug("调用rpc服务后")
	return nil
}
