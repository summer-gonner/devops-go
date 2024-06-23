package redis

import (
	"context"
	"devops-go/basicdata/common/errors/rpcerror"
	"devops-go/basicdata/server/sys/sysclient"
	"encoding/json"

	"devops-go/basicdata/api/internal/svc"
	"devops-go/basicdata/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RedisAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Redis新增数据
func NewRedisAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RedisAddLogic {
	return &RedisAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
func (l *RedisAddLogic) RedisAdd(req *types.ApiRedisReq) (resp *types.ApiRedisResp, err error) {
	addRes, err := l.svcCtx.Sys.RedisAdd(l.ctx, &sysclient.RedisReq{
		Name:     req.Name,
		NickName: req.NickName,
		Password: req.Password,
		Email:    req.Email,
	})

	if err != nil {
		resJson, _ := json.Marshal(addRes)
		logx.WithContext(l.ctx).Errorf("Redis新增数据测试：操作失败，请求参数param = %s，异常信息errMsg = %s", resJson, err.Error())
		return nil, rpcerror.New(err)
	}

	return &types.ApiRedisResp{
		Code:    200,
		Message: "操作成功",
		Data: types.ApiRedisReq{
			Name:     addRes.Name,
			NickName: addRes.NickName,
			Password: addRes.Password,
			Email:    addRes.Email,
		},
	}, nil
}
