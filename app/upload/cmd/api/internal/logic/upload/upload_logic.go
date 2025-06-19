package upload

import (
	"context"

	"Luckify/app/upload/cmd/api/internal/svc"
	"Luckify/app/upload/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 文件上传
func NewUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadLogic {
	return &UploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadLogic) Upload(req *types.UserUploadReq) (resp *types.UserUploadResp, err error) {
	// todo: add your logic here and delete this line

	return
}
