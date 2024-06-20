package user

import (
	"net/http"

	"devops-go/api/internal/logic/sys/user"
	"devops-go/api/internal/svc"
	"devops-go/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 用户管理-新增用户
func UserAddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddUserReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewUserAddLogic(r.Context(), svcCtx)
		resp, err := l.UserAdd(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
