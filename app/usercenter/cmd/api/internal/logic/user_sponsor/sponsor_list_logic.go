package user_sponsor

import (
	"Luckify/app/usercenter/cmd/rpc/usercenter"
	"Luckify/common/utility"
	"context"
	"github.com/jinzhu/copier"

	"Luckify/app/usercenter/cmd/api/internal/svc"
	"Luckify/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SponsorListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 我的赞助商列表
func NewSponsorListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SponsorListLogic {
	return &SponsorListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SponsorListLogic) SponsorList(req *types.SponsorListReq) (resp *types.SponsorListResp, err error) {
	pbResp, err := l.svcCtx.UsercenterRpc.SearchUserSponsor(l.ctx, &usercenter.SearchUserSponsorReq{
		Page:   req.Page,
		Limit:  req.PageSize,
		UserId: utility.GetUserIdFromCtx(l.ctx),
	})
	if err != nil {
		return nil, err
	}

	sponsorList := make([]types.Sponsor, 0)
	for _, pbSponsor := range pbResp.UserSponsor {
		sponsor := types.Sponsor{}
		_ = copier.Copy(&sponsor, pbSponsor)
		sponsorList = append(sponsorList, sponsor)
	}

	return &types.SponsorListResp{
		List: sponsorList,
	}, nil
}
