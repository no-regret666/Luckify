package lottery

import (
	"Luckify/app/lottery/cmd/rpc/pb"
	"Luckify/app/lottery/model"
	"Luckify/app/usercenter/cmd/rpc/usercenter"
	"Luckify/common/utility"
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"Luckify/app/lottery/cmd/api/internal/svc"
	"Luckify/app/lottery/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LotteryDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 抽奖详情
func NewLotteryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LotteryDetailLogic {
	return &LotteryDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LotteryDetailLogic) LotteryDetail(req *types.LotteryDetailReq) (resp *types.LotteryDetailResp, err error) {
	userId := utility.GetUserIdFromCtx(l.ctx)
	pbResp, err := l.svcCtx.LotteryRpc.LotteryDetail(l.ctx, &pb.LotteryDetailReq{
		Id:     req.Id,
		UserId: userId,
	})
	if err != nil {
		return nil, errors.Wrapf(model.ErrLotteryDetail, "rpc LotteryDetail error: %s", err.Error())
	}

	resp = &types.LotteryDetailResp{}
	resp.Prizes = make([]*types.Prize, 0)
	_ = copier.Copy(resp, pbResp.Lottery)
	for _, prize := range pbResp.Prizes {
		item := &types.Prize{}
		_ = copier.Copy(item, prize)
		resp.Prizes = append(resp.Prizes, item)
	}
	pbSponsor, err := l.svcCtx.UsercenterRpc.SponsorDetail(l.ctx, &usercenter.SponsorDetailReq{})
	return
}
