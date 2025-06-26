package lottery

import (
	"context"

	"Luckify/app/lottery/cmd/api/internal/svc"
	"Luckify/app/lottery/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserCreatedLotteryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取当前用户发起的抽奖列表
func NewGetUserCreatedLotteryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserCreatedLotteryListLogic {
	return &GetUserCreatedLotteryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserCreatedLotteryListLogic) GetUserCreatedLotteryList(req *types.GetUserCreatedLotteryListReq) (resp *types.GetUserCreatedLotteryListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
