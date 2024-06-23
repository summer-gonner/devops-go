package redis

import (
	"context"

	"devops-go/basicdata/api/internal/svc"
	"devops-go/basicdata/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RedisUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Redis修改数据
func NewRedisUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RedisUpdateLogic {
	return &RedisUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RedisUpdateLogic) RedisUpdate(req *types.ApiRedisReq) (resp *types.ApiRedisResp, err error) {
	// todo: add your logic here and delete this line

	return
}
