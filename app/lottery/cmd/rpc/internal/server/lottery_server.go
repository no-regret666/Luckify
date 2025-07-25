// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: lottery.proto

package server

import (
	"context"

	"Luckify/app/lottery/cmd/rpc/internal/logic"
	"Luckify/app/lottery/cmd/rpc/internal/svc"
	"Luckify/app/lottery/cmd/rpc/pb"
)

type LotteryServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedLotteryServer
}

func NewLotteryServer(svcCtx *svc.ServiceContext) *LotteryServer {
	return &LotteryServer{
		svcCtx: svcCtx,
	}
}

func (s *LotteryServer) SearchLottery(ctx context.Context, in *pb.SearchLotteryReq) (*pb.SearchLotteryResp, error) {
	l := logic.NewSearchLotteryLogic(ctx, s.svcCtx)
	return l.SearchLottery(in)
}

func (s *LotteryServer) GetLotteryListSlowQuery(ctx context.Context, in *pb.GetLotteryListSlowQueryReq) (*pb.GetLotteryListSlowQueryResp, error) {
	l := logic.NewGetLotteryListSlowQueryLogic(ctx, s.svcCtx)
	return l.GetLotteryListSlowQuery(in)
}

func (s *LotteryServer) GetLotteryListAfterLogin(ctx context.Context, in *pb.GetLotteryListAfterLoginReq) (*pb.GetLotteryListAfterLoginResp, error) {
	l := logic.NewGetLotteryListAfterLoginLogic(ctx, s.svcCtx)
	return l.GetLotteryListAfterLogin(in)
}

func (s *LotteryServer) AddLottery(ctx context.Context, in *pb.AddLotteryReq) (*pb.AddLotteryResp, error) {
	l := logic.NewAddLotteryLogic(ctx, s.svcCtx)
	return l.AddLottery(in)
}

func (s *LotteryServer) LotteryDetail(ctx context.Context, in *pb.LotteryDetailReq) (*pb.LotteryDetailResp, error) {
	l := logic.NewLotteryDetailLogic(ctx, s.svcCtx)
	return l.LotteryDetail(in)
}

func (s *LotteryServer) AnnounceLottery(ctx context.Context, in *pb.AnnounceLotteryReq) (*pb.AnnounceLotteryResp, error) {
	l := logic.NewAnnounceLotteryLogic(ctx, s.svcCtx)
	return l.AnnounceLottery(in)
}

func (s *LotteryServer) GetUserWonList(ctx context.Context, in *pb.GetUserWonListReq) (*pb.GetUserWonListResp, error) {
	l := logic.NewGetUserWonListLogic(ctx, s.svcCtx)
	return l.GetUserWonList(in)
}

func (s *LotteryServer) GetWonListByLotteryId(ctx context.Context, in *pb.GetWonListByLotteryIdReq) (*pb.GetWonListByLotteryIdResp, error) {
	l := logic.NewGetWonListByLotteryIdLogic(ctx, s.svcCtx)
	return l.GetWonListByLotteryId(in)
}

func (s *LotteryServer) GetUserAllList(ctx context.Context, in *pb.GetUserAllListReq) (*pb.GetUserAllListResp, error) {
	l := logic.NewGetUserAllListLogic(ctx, s.svcCtx)
	return l.GetUserAllList(in)
}

func (s *LotteryServer) GetUserCreatedList(ctx context.Context, in *pb.GetUserCreatedListReq) (*pb.GetUserCreatedListResp, error) {
	l := logic.NewGetUserCreatedListLogic(ctx, s.svcCtx)
	return l.GetUserCreatedList(in)
}

func (s *LotteryServer) SearchLotteryParticipation(ctx context.Context, in *pb.SearchLotteryParticipationReq) (*pb.SearchLotteryParticipationResp, error) {
	l := logic.NewSearchLotteryParticipationLogic(ctx, s.svcCtx)
	return l.SearchLotteryParticipation(in)
}

func (s *LotteryServer) GetLotteryStatistic(ctx context.Context, in *pb.GetLotteryStatisticReq) (*pb.GetLotteryStatisticResp, error) {
	l := logic.NewGetLotteryStatisticLogic(ctx, s.svcCtx)
	return l.GetLotteryStatistic(in)
}

func (s *LotteryServer) AddLotteryParticipation(ctx context.Context, in *pb.AddLotteryParticipationReq) (*pb.AddLotteryParticipationResp, error) {
	l := logic.NewAddLotteryParticipationLogic(ctx, s.svcCtx)
	return l.AddLotteryParticipation(in)
}

func (s *LotteryServer) AddInstantLotteryParticipation(ctx context.Context, in *pb.AddInstantLotteryParticipationReq) (*pb.AddInstantLotteryParticipationResp, error) {
	l := logic.NewAddInstantLotteryParticipationLogic(ctx, s.svcCtx)
	return l.AddInstantLotteryParticipation(in)
}

func (s *LotteryServer) CheckLotteryParticipated(ctx context.Context, in *pb.CheckLotteryParticipatedReq) (*pb.CheckLotteryParticipatedResp, error) {
	l := logic.NewCheckLotteryParticipatedLogic(ctx, s.svcCtx)
	return l.CheckLotteryParticipated(in)
}

func (s *LotteryServer) CheckLotteryCreated(ctx context.Context, in *pb.CheckLotteryCreatedReq) (*pb.CheckLotteryCreatedResp, error) {
	l := logic.NewCheckLotteryCreatedLogic(ctx, s.svcCtx)
	return l.CheckLotteryCreated(in)
}
