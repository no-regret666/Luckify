package user_sponsor

import (
	"Luckify/app/usercenter/cmd/rpc/pb"
	"Luckify/app/usercenter/model"
	"Luckify/common/utility"
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"Luckify/app/usercenter/cmd/api/internal/svc"
	"Luckify/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddSponsorLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 添加赞助商
func NewAddSponsorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSponsorLogic {
	return &AddSponsorLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddSponsorLogic) AddSponsor(req *types.CreateSponsorReq) (resp *types.CreateSponsorResp, err error) {
	pbReq := new(pb.AddUserSponsorReq)
	_ = copier.Copy(pbReq, req)
	pbReq.UserId = utility.GetUserIdFromCtx(l.ctx)
	pbResp, err := l.svcCtx.UsercenterRpc.AddUserSponsor(l.ctx, pbReq)
	if err != nil {
		return nil, errors.Wrapf(model.ErrAddUserSponsor, "rpc error: %v", err)
	}

	return &types.CreateSponsorResp{
		Id: pbResp.Id,
	}, nil
}
