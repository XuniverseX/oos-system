package file

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"oos-system/app/api/internal/logic/file"
	"oos-system/app/api/internal/svc"
	"oos-system/app/api/internal/types"
)

func DownloadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DownloadReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := file.NewDownloadLogic(r.Context(), svcCtx)
		err := l.Download(&req, w)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
