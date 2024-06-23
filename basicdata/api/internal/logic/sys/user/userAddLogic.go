package user

import (
	"context"
	"devops-go/basicdata/common/errors/rpcerror"
	"devops-go/basicdata/server/sys/sysclient"
	"encoding/json"

	"devops-go/basicdata/api/internal/svc"
	"devops-go/basicdata/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户管理-新增用户
func NewUserAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserAddLogic {
	return &UserAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserAddLogic) UserAdd(req *types.AddUserReq) (resp *types.AddUserResp, err error) {
	res, err := l.svcCtx.Sys.UserAdd(l.ctx, &sysclient.UserAddReq{
		Name:     req.Name,
		NickName: req.NickName,
		Password: req.Password,
		Email:    req.Email,
		RoleId:   req.RoleId,
		Status:   req.Status,
		CreateBy: "songfayuan",
	})

	if err != nil {
		reqJson, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("添加用户信息失败，请求参数：%s，异常信息：%s", reqJson, err.Error())
		return nil, rpcerror.New(err)
	}

	return &types.AddUserResp{
		Code:    200,
		Message: "添加用户成功",
		Data:    types.ReceiptUserData{Id: res.Id},
	}, nil
}
