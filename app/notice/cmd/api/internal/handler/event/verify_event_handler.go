package event

import (
	"Luckify/common/response"
	"net/http"

	"github.com/go-playground/validator/v10"

	"Luckify/app/notice/cmd/api/internal/logic/event"
	"Luckify/app/notice/cmd/api/internal/svc"
	"Luckify/app/notice/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 验证小程序回调消息
func VerifyEventHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.VerifyEventReq
		if err := httpx.Parse(r, &req); err != nil {
			response.ParamErrorResult(r, w, err)
			return
		}

		err := validator.New().StructCtx(r.Context(), req)
		if err != nil {
			response.ParamErrorResult(r, w, err)
			return
		}

		l := event.NewVerifyEventLogic(r.Context(), svcCtx)
		resp, err := l.VerifyEvent(&req, w)
		response.HttpResult(r, w, resp, err)
	}
}
