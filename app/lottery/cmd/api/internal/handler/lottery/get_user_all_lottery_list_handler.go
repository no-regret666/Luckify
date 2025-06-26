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

// 获取当前用户所有抽奖列表
func GetUserAllLotteryListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetUserAllLotteryListReq
		if err := httpx.Parse(r, &req); err != nil {
			response.ParamErrorResult(r, w, err)
			return
		}

		err := validator.New().StructCtx(r.Context(), req)
		if err != nil {
			response.ParamErrorResult(r, w, err)
			return
		}

		l := lottery.NewGetUserAllLotteryListLogic(r.Context(), svcCtx)
		resp, err := l.GetUserAllLotteryList(&req)
		response.HttpResult(r, w, resp, err)
	}
}
