// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: comment.proto

package server

import (
	"context"

	"Luckify/app/comment/cmd/rpc/internal/logic"
	"Luckify/app/comment/cmd/rpc/internal/svc"
	"Luckify/app/comment/cmd/rpc/pb"
)

type CommentServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedCommentServer
}

func NewCommentServer(svcCtx *svc.ServiceContext) *CommentServer {
	return &CommentServer{
		svcCtx: svcCtx,
	}
}

func (s *CommentServer) AddComment(ctx context.Context, in *pb.AddCommentReq) (*pb.AddCommentResp, error) {
	l := logic.NewAddCommentLogic(ctx, s.svcCtx)
	return l.AddComment(in)
}

func (s *CommentServer) UpdateComment(ctx context.Context, in *pb.UpdateCommentReq) (*pb.UpdateCommentResp, error) {
	l := logic.NewUpdateCommentLogic(ctx, s.svcCtx)
	return l.UpdateComment(in)
}

func (s *CommentServer) DelComment(ctx context.Context, in *pb.DelCommentReq) (*pb.DelCommentResp, error) {
	l := logic.NewDelCommentLogic(ctx, s.svcCtx)
	return l.DelComment(in)
}

func (s *CommentServer) GetCommentById(ctx context.Context, in *pb.GetCommentByIdReq) (*pb.GetCommentByIdResp, error) {
	l := logic.NewGetCommentByIdLogic(ctx, s.svcCtx)
	return l.GetCommentById(in)
}

func (s *CommentServer) SearchComment(ctx context.Context, in *pb.SearchCommentReq) (*pb.SearchCommentResp, error) {
	l := logic.NewSearchCommentLogic(ctx, s.svcCtx)
	return l.SearchComment(in)
}

func (s *CommentServer) IsPraise(ctx context.Context, in *pb.IsPraiseReq) (*pb.IsPraiseResp, error) {
	l := logic.NewIsPraiseLogic(ctx, s.svcCtx)
	return l.IsPraise(in)
}

func (s *CommentServer) PraiseComment(ctx context.Context, in *pb.PraiseCommentReq) (*pb.PraiseCommentResp, error) {
	l := logic.NewPraiseCommentLogic(ctx, s.svcCtx)
	return l.PraiseComment(in)
}

func (s *CommentServer) GetCommentLastId(ctx context.Context, in *pb.GetCommentLastIdReq) (*pb.GetCommentLastIdResp, error) {
	l := logic.NewGetCommentLastIdLogic(ctx, s.svcCtx)
	return l.GetCommentLastId(in)
}

func (s *CommentServer) IsPraiseList(ctx context.Context, in *pb.IsPraiseListReq) (*pb.IsPraiseListResp, error) {
	l := logic.NewIsPraiseListLogic(ctx, s.svcCtx)
	return l.IsPraiseList(in)
}

func (s *CommentServer) GetUserComment(ctx context.Context, in *pb.GetUserCommentReq) (*pb.GetUserCommentResp, error) {
	l := logic.NewGetUserCommentLogic(ctx, s.svcCtx)
	return l.GetUserComment(in)
}

func (s *CommentServer) AddPraise(ctx context.Context, in *pb.AddPraiseReq) (*pb.AddPraiseResp, error) {
	l := logic.NewAddPraiseLogic(ctx, s.svcCtx)
	return l.AddPraise(in)
}

func (s *CommentServer) UpdatePraise(ctx context.Context, in *pb.UpdatePraiseReq) (*pb.UpdatePraiseResp, error) {
	l := logic.NewUpdatePraiseLogic(ctx, s.svcCtx)
	return l.UpdatePraise(in)
}

func (s *CommentServer) DelPraise(ctx context.Context, in *pb.DelPraiseReq) (*pb.DelPraiseResp, error) {
	l := logic.NewDelPraiseLogic(ctx, s.svcCtx)
	return l.DelPraise(in)
}

func (s *CommentServer) GetPraiseById(ctx context.Context, in *pb.GetPraiseByIdReq) (*pb.GetPraiseByIdResp, error) {
	l := logic.NewGetPraiseByIdLogic(ctx, s.svcCtx)
	return l.GetPraiseById(in)
}

func (s *CommentServer) SearchPraise(ctx context.Context, in *pb.SearchPraiseReq) (*pb.SearchPraiseResp, error) {
	l := logic.NewSearchPraiseLogic(ctx, s.svcCtx)
	return l.SearchPraise(in)
}

func (s *CommentServer) CheckUserPraise(ctx context.Context, in *pb.CheckUserPraiseReq) (*pb.CheckUserPraiseResp, error) {
	l := logic.NewCheckUserPraiseLogic(ctx, s.svcCtx)
	return l.CheckUserPraise(in)
}
