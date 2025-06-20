package upload

import (
	"net/http"

	"Luckify/common/response"
	"github.com/go-playground/validator/v10"

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
			response.ParamErrorResult(r, w, err)
			return
		}

		err := validator.New().StructCtx(r.Context(), req)
		if err != nil {
			response.ParamErrorResult(r, w, err)
			return
		}

		l := upload.NewUploadLogic(r.Context(), svcCtx)
		resp, err := l.Upload(&req)
		response.HttpResult(r, w, resp, err)
	}
}
