package redis

import (
	"devops-go/basicdata/api/internal/logic/redis"
	"devops-go/basicdata/api/internal/svc"
	"devops-go/basicdata/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

// Redis查询数据
func RedisGetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ApiRedisGetReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := redis.NewRedisGetLogic(r.Context(), svcCtx)
		resp, err := l.RedisGet(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
