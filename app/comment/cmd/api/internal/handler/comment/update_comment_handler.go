package comment

import (
	"Luckify/common/response"
	"net/http"

	"github.com/go-playground/validator/v10"

	"Luckify/app/comment/cmd/api/internal/logic/comment"
	"Luckify/app/comment/cmd/api/internal/svc"
	"Luckify/app/comment/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 更新评论
func UpdateCommentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CommentUpdateReq
		if err := httpx.Parse(r, &req); err != nil {
			response.ParamErrorResult(r, w, err)
			return
		}

		err := validator.New().StructCtx(r.Context(), req)
		if err != nil {
			response.ParamErrorResult(r, w, err)
			return
		}

		l := comment.NewUpdateCommentLogic(r.Context(), svcCtx)
		resp, err := l.UpdateComment(&req)
		response.HttpResult(r, w, resp, err)
	}
}
