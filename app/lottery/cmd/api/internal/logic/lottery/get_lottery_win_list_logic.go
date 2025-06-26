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

type GetLotteryWinListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取当前抽奖中奖者名单
func NewGetLotteryWinListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLotteryWinListLogic {
	return &GetLotteryWinListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLotteryWinListLogic) GetLotteryWinList(req *types.GetLotteryWinListReq) (resp *types.GetLotteryWinListResp, err error) {
	pbResp, err := l.svcCtx.LotteryRpc.GetWonListByLotteryId(l.ctx, &pb.GetWonListByLotteryIdReq{
		LotteryId: req.LotteryId,
	})
	if err != nil {
		return nil, errors.Wrapf(model.ErrGetLotteryWinList, "rpc error: %v", err)
	}

	list := make([]*types.WonList, 0)
	for _, pbItem := range pbResp.List {
		item := &types.WonList{}
		item.Prize = new(types.LotteryPrize)
		_ = copier.Copy(item.Prize, pbItem.Prize)
		item.WinnerCount = pbItem.WinnerCount
		item.Users = make([]*types.UserInfo, 0)
		for _, pbUser := range pbItem.Users {
			user := &types.UserInfo{}
			_ = copier.Copy(user, pbUser)
			item.Users = append(item.Users, user)
		}
		list = append(list, item)
	}

	return &types.GetLotteryWinListResp{List: list}, nil

	return
}
