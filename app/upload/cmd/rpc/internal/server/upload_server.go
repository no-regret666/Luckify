// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: upload.proto

package server

import (
	"context"

	"Luckify/app/upload/cmd/rpc/internal/logic"
	"Luckify/app/upload/cmd/rpc/internal/svc"
	"Luckify/app/upload/cmd/rpc/pb"
)

type UploadServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedUploadServer
}

func NewUploadServer(svcCtx *svc.ServiceContext) *UploadServer {
	return &UploadServer{
		svcCtx: svcCtx,
	}
}

func (s *UploadServer) Upload(ctx context.Context, in *pb.FileUploadReq) (*pb.FileUploadResp, error) {
	l := logic.NewUploadLogic(ctx, s.svcCtx)
	return l.Upload(in)
}
