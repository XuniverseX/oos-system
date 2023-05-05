package file

import (
	"net/http"

	"oos-system/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"oos-system/app/api/internal/logic/file"
	"oos-system/app/api/internal/svc"
	"oos-system/app/api/internal/types"
)

func MergeChunkHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MergeChunkReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := file.NewMergeChunkLogic(r.Context(), svcCtx)
		resp, err := l.MergeChunk(&req)
		result.HttpResult(r, w, resp, err)
	}
}
