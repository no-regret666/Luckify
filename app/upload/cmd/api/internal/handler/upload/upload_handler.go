package upload

import (
	"net/http"

	"Luckify/app/upload/cmd/api/internal/logic/upload"
	"Luckify/app/upload/cmd/api/internal/svc"
	"Luckify/app/upload/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 文件上传
func UploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserUploadReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upload.NewUploadLogic(r.Context(), svcCtx)
		resp, err := l.Upload(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
