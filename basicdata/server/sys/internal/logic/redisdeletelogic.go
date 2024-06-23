package logic

import (
	"context"
	"errors"

	"devops-go/basicdata/server/sys/internal/svc"
	"devops-go/basicdata/server/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type RedisDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRedisDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RedisDeleteLogic {
	return &RedisDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RedisDeleteLogic) RedisDelete(in *sysclient.RedisReq) (*sysclient.RedisResp, error) {
	var key = "allen"
	res, err := l.svcCtx.RedisClient.Del(key)
	if err != nil {
		return nil, errors.New("删除Redis异常")
	}
	logx.Infof("删除Redis数据结果：%s", res)
	return &sysclient.RedisResp{}, nil
}
