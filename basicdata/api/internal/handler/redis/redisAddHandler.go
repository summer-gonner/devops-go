package redis

import (
	"devops-go/basicdata/api/internal/logic/redis"
	"devops-go/basicdata/api/internal/svc"
	"devops-go/basicdata/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

// Redis新增数据
func RedisAddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ApiRedisReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := redis.NewRedisAddLogic(r.Context(), svcCtx)
		resp, err := l.RedisAdd(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
