package comment

import (
	"net/http"

	"github.com/go-playground/validator/v10"

	"Luckify/app/comment/cmd/api/internal/logic/comment"
	"Luckify/app/comment/cmd/api/internal/svc"
	"Luckify/app/comment/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 获取评论详情
func GetCommentDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CommentDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			response.ParamErrorResult(r, w, err)
			return
		}

		err := validator.New().StructCtx(r.Context(), req)
		if err != nil {
			response.ParamErrorResult(r, w, err)
			return
		}

		l := comment.NewGetCommentDetailLogic(r.Context(), svcCtx)
		resp, err := l.GetCommentDetail(&req)
		response.HttpResult(r, w, resp, err)
	}
}
