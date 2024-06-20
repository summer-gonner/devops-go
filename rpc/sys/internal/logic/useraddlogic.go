package logic

import (
	"context"
	"devops-go/common/utils"
	"devops-go/rpc/model/model/sysmodel"
	"errors"
	sq "github.com/Masterminds/squirrel"
	"time"

	"devops-go/rpc/sys/internal/svc"
	"devops-go/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserAddLogic {
	return &UserAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserAddLogic) UserAdd(in *sysclient.UserAddReq) (*sysclient.UserAddResp, error) {
	if in.Name == "" {
		return nil, errors.New("账号不能为空")
	}
	if in.NickName == "" {
		return nil, errors.New("姓名不能为空")
	}
	if in.Email == "" {
		return nil, errors.New("邮箱不能为空")
	}

	//校验账号是否已存在
	selectBuilder := l.svcCtx.UserModel.CountBuilder("id").Where(sq.Eq{"name": in.Name})
	count, _ := l.svcCtx.UserModel.FindCount(l.ctx, selectBuilder)
	if count > 0 {
		logx.WithContext(l.ctx).Errorf("账号已存在，添加失败，userName = %s", in.Name)
		return nil, errors.New("账号已存在")
	}

	if in.Password == "" {
		in.Password = "123456"
	}
	hashedPassword, err := utils.GenerateFromPassword(in.Password)
	if err != nil {
		return nil, errors.New("密码加密出错")
	}

	//插入数据
	result, err := l.svcCtx.UserModel.Insert(l.ctx, &sysmodel.SysUser{
		Name:       in.Name,
		NickName:   in.NickName,
		Avatar:     "",
		Password:   hashedPassword,
		Salt:       "",
		Email:      in.Email,
		Mobile:     "",
		Status:     0,
		CreateBy:   in.CreateBy,
		UpdateTime: time.Time{},
		DelFlag:    0,
	})
	if err != nil {
		return nil, err
	}
	insertId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &sysclient.UserAddResp{Id: insertId}, nil
}
