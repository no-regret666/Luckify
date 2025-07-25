package checkin

import (
	"Luckify/common/response"
	"net/http"

	"github.com/go-playground/validator/v10"

	"Luckify/app/checkin/cmd/api/internal/logic/checkin"
	"Luckify/app/checkin/cmd/api/internal/svc"
	"Luckify/app/checkin/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 获取用户任务完成状态
func GetTasksHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetTasksReq
		if err := httpx.Parse(r, &req); err != nil {
			response.ParamErrorResult(r, w, err)
			return
		}

		err := validator.New().StructCtx(r.Context(), req)
		if err != nil {
			response.ParamErrorResult(r, w, err)
			return
		}

		l := checkin.NewGetTasksLogic(r.Context(), svcCtx)
		resp, err := l.GetTasks(&req)
		response.HttpResult(r, w, resp, err)
	}
}
