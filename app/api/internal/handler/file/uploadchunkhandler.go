package file

import (
	"net/http"

	"oos-system/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"oos-system/app/api/internal/logic/file"
	"oos-system/app/api/internal/svc"
	"oos-system/app/api/internal/types"
)

func UploadChunkHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UploadChunkReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := file.NewUploadChunkLogic(r.Context(), svcCtx)
		resp, err := l.UploadChunk(&req, r)
		result.HttpResult(r, w, resp, err)
	}
}
