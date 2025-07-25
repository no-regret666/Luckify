// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package handler

import (
	"net/http"

	user "Luckify/app/usercenter/cmd/api/internal/handler/user"
	user_sponsor "Luckify/app/usercenter/cmd/api/internal/handler/user_sponsor"
	"Luckify/app/usercenter/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				// 登录
				Method:  http.MethodPost,
				Path:    "/user/login",
				Handler: user.LoginHandler(serverCtx),
			},
			{
				// 注册
				Method:  http.MethodPost,
				Path:    "/user/register",
				Handler: user.RegisterHandler(serverCtx),
			},
			{
				// 微信登录注册
				Method:  http.MethodPost,
				Path:    "/user/wxMiniAuth",
				Handler: user.WxMiniAuthHandler(serverCtx),
			},
		},
		rest.WithPrefix("/usercenter/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				// 获取用户信息
				Method:  http.MethodPost,
				Path:    "/user/detail",
				Handler: user.DetailHandler(serverCtx),
			},
			{
				// 更新用户信息
				Method:  http.MethodPut,
				Path:    "/user/update",
				Handler: user.UpdateHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/usercenter/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				// 删除赞助商
				Method:  http.MethodPost,
				Path:    "/userContact/sponsorDel",
				Handler: user_sponsor.SponsorDelHandler(serverCtx),
			},
			{
				// 修改赞助商
				Method:  http.MethodPost,
				Path:    "/userContact/upDateSponsor",
				Handler: user_sponsor.UpdateSponsorHandler(serverCtx),
			},
			{
				// 添加赞助商
				Method:  http.MethodPost,
				Path:    "/userSponsor/addSponsor",
				Handler: user_sponsor.AddSponsorHandler(serverCtx),
			},
			{
				// 我的赞助商列表
				Method:  http.MethodPost,
				Path:    "/userSponsor/sponsorList",
				Handler: user_sponsor.SponsorListHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/usercenter/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				// 赞助商详情
				Method:  http.MethodPost,
				Path:    "/userSponsor/sponsorDetail",
				Handler: user_sponsor.SponsorDetailHandler(serverCtx),
			},
		},
		rest.WithPrefix("/usercenter/v1"),
	)
}
