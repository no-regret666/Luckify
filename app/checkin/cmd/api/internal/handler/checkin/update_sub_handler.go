package checkin

import (
	"net/http"

	"github.com/go-playground/validator/v10"

	"Luckify/app/checkin/cmd/api/internal/logic/checkin"
	"Luckify/app/checkin/cmd/api/internal/svc"
	"Luckify/app/checkin/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 改变任务签到状态
func UpdateSubHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateSubReq
		if err := httpx.Parse(r, &req); err != nil {
			response.ParamErrorResult(r, w, err)
			return
		}

		err := validator.New().StructCtx(r.Context(), req)
		if err != nil {
			response.ParamErrorResult(r, w, err)
			return
		}

		l := checkin.NewUpdateSubLogic(r.Context(), svcCtx)
		resp, err := l.UpdateSub(&req)
		response.HttpResult(r, w, resp, err)
	}
}
