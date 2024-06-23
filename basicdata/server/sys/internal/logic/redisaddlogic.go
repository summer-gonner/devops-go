package logic

import (
	"context"
	"encoding/json"
	"errors"

	"devops-go/basicdata/server/sys/internal/svc"
	"devops-go/basicdata/server/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type RedisAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRedisAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RedisAddLogic {
	return &RedisAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// redis增删改查
func (l *RedisAddLogic) RedisAdd(in *sysclient.RedisReq) (*sysclient.RedisResp, error) {
	if in.Name == "" {
		return nil, errors.New("账号不能为空")
	}
	if in.NickName == "" {
		return nil, errors.New("姓名不能为空")
	}
	if in.Email == "" {
		return nil, errors.New("邮箱不能为空")
	}

	// 将结构体转换为 JSON 字符串
	jsonString, err := json.Marshal(in)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("Redis新增数据测试：数据转换json出错，异常信息errMsg = %s", err.Error())
		return nil, errors.New("数据转换异常")
	}

	// 添加数据到Redis
	var key = "allen"
	addErr := l.svcCtx.RedisClient.Set(key, string(jsonString))
	if addErr != nil {
		return nil, errors.New("存储Redis异常")
	}

	// 查询Redis的数据
	getData, err := l.svcCtx.RedisClient.Get(key)

	// 将字符串数据解码为 RedisResp 结构体
	var redisResp sysclient.RedisResp
	err = json.Unmarshal([]byte(getData), &redisResp)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("Redis新增数据测试：返回值封装异常，异常信息errMsg = %s", err.Error())
		return nil, errors.New("返回值封装异常")
	}

	return &redisResp, nil
}
