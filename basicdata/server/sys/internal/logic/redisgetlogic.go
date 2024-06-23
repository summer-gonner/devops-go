package logic

import (
	"context"
	"encoding/json"
	"errors"

	"devops-go/basicdata/server/sys/internal/svc"
	"devops-go/basicdata/server/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type RedisGetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRedisGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RedisGetLogic {
	return &RedisGetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RedisGetLogic) RedisGet(in *sysclient.RedisReq) (*sysclient.RedisResp, error) {
	// 查询Redis的数据
	getData, err := l.svcCtx.RedisClient.Get(in.Key)

	// 将字符串数据解码为 RedisResp 结构体
	var redisResp sysclient.RedisResp
	err = json.Unmarshal([]byte(getData), &redisResp)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("Redis新增数据测试：返回值封装异常，异常信息errMsg = %s", err.Error())
		return nil, errors.New("返回值封装异常")
	}

	return &redisResp, nil
}
