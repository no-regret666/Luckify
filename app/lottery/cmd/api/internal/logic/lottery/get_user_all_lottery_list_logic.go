package lottery

import (
	"context"

	"Luckify/app/lottery/cmd/api/internal/svc"
	"Luckify/app/lottery/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserAllLotteryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取当前用户所有抽奖列表
func NewGetUserAllLotteryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserAllLotteryListLogic {
	return &GetUserAllLotteryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserAllLotteryListLogic) GetUserAllLotteryList(req *types.GetUserAllLotteryListReq) (resp *types.GetUserAllLotteryListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
