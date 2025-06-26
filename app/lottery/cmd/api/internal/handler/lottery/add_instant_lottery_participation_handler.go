package lottery

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/no-regret666/Luckify/common/response"

	"Luckify/app/lottery/cmd/api/internal/logic/lottery"
	"Luckify/app/lottery/cmd/api/internal/svc"
	"Luckify/app/lottery/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 参与即抽即中抽奖
func AddInstantLotteryParticipationHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddInstantLotteryParticipationReq
		if err := httpx.Parse(r, &req); err != nil {
			response.ParamErrorResult(r, w, err)
			return
		}

		err := validator.New().StructCtx(r.Context(), req)
		if err != nil {
			response.ParamErrorResult(r, w, err)
			return
		}

		l := lottery.NewAddInstantLotteryParticipationLogic(r.Context(), svcCtx)
		resp, err := l.AddInstantLotteryParticipation(&req)
		response.HttpResult(r, w, resp, err)
	}
}
