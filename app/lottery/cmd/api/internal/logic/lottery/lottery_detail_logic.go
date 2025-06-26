package lottery

import (
	"context"

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
	// todo: add your logic here and delete this line

	return
}
