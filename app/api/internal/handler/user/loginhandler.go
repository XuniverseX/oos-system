package user

import (
	"net/http"

	"oos-system/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"oos-system/app/api/internal/logic/user"
	"oos-system/app/api/internal/svc"
	"oos-system/app/api/internal/types"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := user.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		result.HttpResult(r, w, resp, err)
	}
}
