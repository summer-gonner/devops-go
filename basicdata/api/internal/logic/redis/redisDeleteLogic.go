package redis

import (
	"context"

	"devops-go/basicdata/api/internal/svc"
	"devops-go/basicdata/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RedisDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Redis删除数据
func NewRedisDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RedisDeleteLogic {
	return &RedisDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RedisDeleteLogic) RedisDelete(req *types.ApiRedisReq) (resp *types.ApiRedisResp, err error) {
	// todo: add your logic here and delete this line

	return
}
