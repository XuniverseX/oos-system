package file

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"oos-system/app/api/internal/logic/file"
	"oos-system/app/api/internal/svc"
	"oos-system/app/api/internal/types"
)

func DownloadUrlHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DownloadUrlReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := file.NewDownloadUrlLogic(r.Context(), svcCtx)
		resp, err := l.DownloadUrl(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
