package logic

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strconv"

	"devops-go/basicdata/server/sys/internal/svc"
	"devops-go/basicdata/server/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoLogic) UserInfo(in *sysclient.InfoReq) (*sysclient.InfoResp, error) {
	rowBuilder := l.svcCtx.UserModel.RowBuilder().Where(sq.Eq{"id": in.UserId})
	userInfo, err := l.svcCtx.UserModel.FindOneByQuery(l.ctx, rowBuilder)

	switch err {
	case nil:
	case sqlx.ErrNotFound:
		logx.WithContext(l.ctx).Infof("用户不存在userId:%s", in.UserId)
		return nil, fmt.Errorf("用户不存在userId:%s", strconv.FormatInt(in.UserId, 10))
	default:
		return nil, err
	}

	//var list []*sys.MenuListTree
	//var listUrls []string

	return &sysclient.InfoResp{
		Avatar:         "11111",
		Name:           userInfo.Name,
		MenuListTree:   nil,
		BackgroundUrls: nil,
		ResetPwd:       false,
	}, nil
}
