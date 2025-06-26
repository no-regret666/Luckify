package lottery

import (
	"context"

	"Luckify/app/lottery/cmd/api/internal/svc"
	"Luckify/app/lottery/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLotteryWinListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取当前用户中奖列表
func NewGetUserLotteryWinListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLotteryWinListLogic {
	return &GetUserLotteryWinListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLotteryWinListLogic) GetUserLotteryWinList(req *types.GetUserLotteryWinListReq) (resp *types.GetUserLotteryWinListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
