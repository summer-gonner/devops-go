package user

import (
	"net/http"

	"devops-go/basicdata/api/internal/logic/sys/user"
	"devops-go/basicdata/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 用户管理-获取当前用户信息
func UserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
