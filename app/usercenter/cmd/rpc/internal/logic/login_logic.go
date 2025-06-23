package logic

import (
	"Luckify/app/usercenter/model"
	"Luckify/common/constants"
	"Luckify/common/utility"
	"Luckify/common/xerr"
	"context"
	"github.com/pkg/errors"

	"Luckify/app/usercenter/cmd/rpc/internal/svc"
	"Luckify/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *pb.LoginReq) (*pb.LoginResp, error) {
	var userId int64
	var err error
	// try login
	switch in.AuthType {
	case constants.UserAuthTypeSystem:
		userId, err = l.loginByMobile(in.AuthKey, in.Password)
	}
	if err != nil {
		return nil, err
	}
	// generate token
	generateTokenLogic := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	pbResp, err := generateTokenLogic.GenerateToken(&pb.GenerateTokenReq{
		UserId: userId,
	})
	if err != nil {
		return nil, errors.Wrapf(model.ErrLogin, "GenerateToken userId: %d", userId)
	}

	return &pb.LoginResp{
		AccessToken:  pbResp.AccessToken,
		AccessExpire: pbResp.AccessExpire,
		RefreshAfter: pbResp.RefreshAfter,
	}, nil
}

func (l *LoginLogic) loginByMobile(mobile, password string) (int64, error) {
	user, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, mobile)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return 0, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "query user by mobile failed,mobile: %s,err: %v", mobile, err)
	}
	if user == nil {
		return 0, errors.Wrapf(model.ErrUserNotExistsError, "mobile: %s", mobile)
	}

	if !(utility.Md5ByString(password) == user.Password) {
		return 0, errors.Wrapf(model.ErrPasswordNotMatch, "password noot correct")
	}
	
	return user.Id, nil
}
