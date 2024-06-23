package logic

import (
	"context"
	"encoding/json"
	"errors"

	"devops-go/basicdata/server/sys/internal/svc"
	"devops-go/basicdata/server/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type RedisUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRedisUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RedisUpdateLogic {
	return &RedisUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RedisUpdateLogic) RedisUpdate(in *sysclient.RedisReq) (*sysclient.RedisResp, error) {
	if in.Name == "" {
		return nil, errors.New("账号不能为空")
	}
	if in.NickName == "" {
		return nil, errors.New("姓名不能为空")
	}
	if in.Email == "" {
		return nil, errors.New("邮箱不能为空")
	}

	// 将结构体数据转为json字符串
	jsonString, err := json.Marshal(in)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("Redis更新数据测试：数据转换json出错，异常信息errMsg = %s", err.Error())
		return nil, errors.New("数据转换异常")
	}

	// 更新数据到Redis，这里的修改，就是设置新值即可
	key := "allen"
	editRes := l.svcCtx.RedisClient.Set(key, string(jsonString))
	if editRes != nil {
		return nil, errors.New("修改Redis数据异常")
	}

	// 查询Redis的数据
	getData, err := l.svcCtx.RedisClient.Get(key)

	// 将字符串数据转换为结构体
	var redisResp sysclient.RedisResp
	err = json.Unmarshal([]byte(getData), &redisResp)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("修改Redis数据测试：返回值封装异常，异常信息errMsg = %s", err.Error())
		return nil, errors.New("返回值封装异常")
	}

	return &redisResp, nil
}
