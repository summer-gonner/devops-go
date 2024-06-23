// Code generated by goctl. DO NOT EDIT.
// Source: sys.proto

package sys

import (
	"context"

	"devops-go/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	InfoReq      = sysclient.InfoReq
	InfoResp     = sysclient.InfoResp
	MenuListTree = sysclient.MenuListTree
	RedisReq     = sysclient.RedisReq
	RedisResp    = sysclient.RedisResp
	UserAddReq   = sysclient.UserAddReq
	UserAddResp  = sysclient.UserAddResp

	Sys interface {
		UserInfo(ctx context.Context, in *InfoReq, opts ...grpc.CallOption) (*InfoResp, error)
		UserAdd(ctx context.Context, in *UserAddReq, opts ...grpc.CallOption) (*UserAddResp, error)
		RedisAdd(ctx context.Context, in *RedisReq, opts ...grpc.CallOption) (*RedisResp, error)
		RedisDelete(ctx context.Context, in *RedisReq, opts ...grpc.CallOption) (*RedisResp, error)
		RedisUpdate(ctx context.Context, in *RedisReq, opts ...grpc.CallOption) (*RedisResp, error)
		RedisGet(ctx context.Context, in *RedisReq, opts ...grpc.CallOption) (*RedisResp, error)
	}

	defaultSys struct {
		cli zrpc.Client
	}
)

func NewSys(cli zrpc.Client) Sys {
	return &defaultSys{
		cli: cli,
	}
}

func (m *defaultSys) UserInfo(ctx context.Context, in *InfoReq, opts ...grpc.CallOption) (*InfoResp, error) {
	client := sysclient.NewSysClient(m.cli.Conn())
	return client.UserInfo(ctx, in, opts...)
}

func (m *defaultSys) UserAdd(ctx context.Context, in *UserAddReq, opts ...grpc.CallOption) (*UserAddResp, error) {
	client := sysclient.NewSysClient(m.cli.Conn())
	return client.UserAdd(ctx, in, opts...)
}

func (m *defaultSys) RedisAdd(ctx context.Context, in *RedisReq, opts ...grpc.CallOption) (*RedisResp, error) {
	client := sysclient.NewSysClient(m.cli.Conn())
	return client.RedisAdd(ctx, in, opts...)
}

func (m *defaultSys) RedisDelete(ctx context.Context, in *RedisReq, opts ...grpc.CallOption) (*RedisResp, error) {
	client := sysclient.NewSysClient(m.cli.Conn())
	return client.RedisDelete(ctx, in, opts...)
}

func (m *defaultSys) RedisUpdate(ctx context.Context, in *RedisReq, opts ...grpc.CallOption) (*RedisResp, error) {
	client := sysclient.NewSysClient(m.cli.Conn())
	return client.RedisUpdate(ctx, in, opts...)
}

func (m *defaultSys) RedisGet(ctx context.Context, in *RedisReq, opts ...grpc.CallOption) (*RedisResp, error) {
	client := sysclient.NewSysClient(m.cli.Conn())
	return client.RedisGet(ctx, in, opts...)
}
