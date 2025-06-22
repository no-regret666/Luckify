package user

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

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户信息
func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	userId := utility.GetUserIdFromCtx(l.ctx)

	pbUserInfo, err := l.svcCtx.UsercenterRpc.GetUserInfo(l.ctx, &pb.GetUserInfoReq{
		Id: userId,
	})
	if err != nil {
		return nil, errors.Wrapf(model.ErrGetUserInfo, "GetUserInfo err: %v,userId: %d,userInfoResp: %+v", err, userId, pbUserInfo)
	}

	userInfo := types.User{}
	_ = copier.Copy(&userInfo, pbUserInfo.User)

	return &types.UserInfoResp{
		UserInfo: userInfo,
	}, nil
}
