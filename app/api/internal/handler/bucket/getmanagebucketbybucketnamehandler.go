package bucket

import (
	"net/http"

	"oos-system/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"oos-system/app/api/internal/logic/bucket"
	"oos-system/app/api/internal/svc"
	"oos-system/app/api/internal/types"
)

func GetManageBucketByBucketNameHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetBucketByBucketNameReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := bucket.NewGetManageBucketByBucketNameLogic(r.Context(), svcCtx)
		resp, err := l.GetManageBucketByBucketName(&req)
		result.HttpResult(r, w, resp, err)
	}
}
