package user_sponsor

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/no-regret666/Luckify/common/response"

	"Luckify/app/usercenter/cmd/api/internal/logic/user_sponsor"
	"Luckify/app/usercenter/cmd/api/internal/svc"
	"Luckify/app/usercenter/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 赞助商详情
func SponsorDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SponosorDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			response.ParamErrorResult(r, w, err)
			return
		}

		err := validator.New().StructCtx(r.Context(), req)
		if err != nil {
			response.ParamErrorResult(r, w, err)
			return
		}

		l := user_sponsor.NewSponsorDetailLogic(r.Context(), svcCtx)
		resp, err := l.SponsorDetail(&req)
		response.HttpResult(r, w, resp, err)
	}
}
