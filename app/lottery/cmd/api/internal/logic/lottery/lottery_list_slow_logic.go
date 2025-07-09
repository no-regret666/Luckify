package lottery

import (
	"Luckify/app/lottery/cmd/rpc/pb"
	"Luckify/app/lottery/model"
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

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
	pbResp, err := l.svcCtx.LotteryRpc.GetLotteryListSlowQuery(l.ctx, &pb.GetLotteryListSlowQueryReq{
		Page:       req.PageIndex,
		Limit:      req.PageSize,
		IsSelected: req.IsSelected,
	})
	if err != nil {
		return nil, errors.Wrapf(model.ErrSearchList, "rpc error: %v", err)
	}

	LotteryList := make([]types.Lottery, 0)
	for _, lottery := range pbResp.List {
		var item types.Lottery
		_ = copier.Copy(&item, lottery)
		LotteryList = append(LotteryList, item)
	}

	return &types.LotteryListSlowQueryResp{
		List: LotteryList,
	}, nil
}
