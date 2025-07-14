package upload

import (
	"Luckify/app/upload/cmd/rpc/upload"
	"Luckify/app/upload/model"
	"Luckify/common/utility"
	"context"
	"github.com/pkg/errors"

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
	userId := utility.GetUserIdFromCtx(l.ctx)
	pbResp, err := l.svcCtx.UploadRpc.Upload(l.ctx, &upload.FileUploadReq{
		UserId:   userId,
		Filename: req.FileName,
		Size:     req.Size,
		Ext:      req.Ext,
		FileData: req.FileData,
	})
	if err != nil {
		return nil, errors.Wrapf(model.ErrUpload, "upload file fail with rpc err: %+v", req)
	}
	logx.Infof("Successfully upload file %+v", pbResp.Url)
	return &types.UserUploadResp{
		Url: pbResp.Url,
	}, nil
}
