package lottery

import (
	"context"

	"Luckify/app/lottery/cmd/api/internal/svc"
	"Luckify/app/lottery/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LotteryListSlowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取抽奖列表（慢查询）
func NewLotteryListSlowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LotteryListSlowLogic {
	return &LotteryListSlowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LotteryListSlowLogic) LotteryListSlow(req *types.LotteryListSlowQueryReq) (resp *types.LotteryListSlowQueryResp, err error) {
	// todo: add your logic here and delete this line

	return
}
