package user

import (
	"context"
	"devops-go/basicdata/api/internal/svc"
	"devops-go/basicdata/api/internal/types"
	"devops-go/basicdata/common/errors/rpcerror"
	"devops-go/basicdata/server/sys/sysclient"
	"strconv"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo() (*types.UserInfoResp, error) {
	var userId int64 = 1
	resp, err := l.svcCtx.Sys.UserInfo(l.ctx, &sysclient.InfoReq{
		UserId: userId,
	})

	if err != nil {
		logx.WithContext(l.ctx).Errorf("根据userId：%s, 查询用户异常：%s", strconv.FormatInt(userId, 10), err.Error())
		return nil, rpcerror.New(err)
	}

	var MenuTree []*types.ListMenuTree

	//组装ant ui中的菜单
	for _, item := range resp.MenuListTree {
		MenuTree = append(MenuTree, &types.ListMenuTree{
			Id:       item.Id,
			Path:     item.Path,
			Name:     item.Name,
			ParentId: item.ParentId,
			Icon:     item.Icon,
		})
	}

	if MenuTree == nil {
		MenuTree = make([]*types.ListMenuTree, 0)
	}

	//组装element ui中的菜单
	var MenuTreeVue []*types.ListMenuTreeVue

	for _, item := range resp.MenuListTree {
		if len(strings.TrimSpace(item.VuePath)) != 0 {
			MenuTreeVue = append(MenuTreeVue, &types.ListMenuTreeVue{
				Id:           item.Id,
				ParentId:     item.ParentId,
				Title:        item.Name,
				Path:         item.VuePath,
				Name:         item.Name,
				Icon:         item.VueIcon,
				VueRedirent:  item.VueRedirect,
				VueComponent: item.VueComponent,
				Meta: types.MenuTreeMeta{
					Title: item.Name,
					Icon:  item.VueIcon,
				},
			})
		}
	}
	if MenuTreeVue == nil {
		MenuTreeVue = make([]*types.ListMenuTreeVue, 0)
	}

	if err != nil {
		logx.Errorf("设置用户：%s, 权限到Redis异常：%+v", resp.Name, err)
	}

	return &types.UserInfoResp{
		Code:    200,
		Message: "成功",
		Data: types.UserInfoData{
			Avatar:      resp.Avatar,
			Name:        resp.Name,
			MenuTree:    MenuTree,
			MenuTreeVue: MenuTreeVue,
			ResetPwd:    resp.ResetPwd,
		},
	}, nil
}
