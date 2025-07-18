// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package types

type CreateSponsorReq struct {
	UserId     int64  `json:"userId"`
	Type       int64  `json:"type"`
	AppletType int64  `json:"appletType"`
	IsShow     int64  `json:"isShow"`
	Name       string `json:"name"`
	Desc       string `json:"desc"`
	Avatar     string `json:"avatar"`
	QrCode     string `json:"qr_code"`
	InputA     string `json:"inputA"`
	InputB     string `json:"inputB"`
}

type CreateSponsorResp struct {
	Id int64 `json:"id"`
}

type LoginReq struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

type LoginResp struct {
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
	RefreshAfter int64  `json:"refreshAfter"`
}

type RegisterReq struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

type RegisterResp struct {
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
	RefreshAfter int64  `json:"refreshAfter"`
}

type Sponsor struct {
	Id         int64  `json:"id"`
	UserId     int64  `json:"userId"`
	Type       int64  `json:"type"`
	AppletType int64  `json:"appletType"`
	IsShow     int64  `json:"isShow"`
	Name       string `json:"name"`
	Desc       string `json:"desc"`
	Avatar     string `json:"avatar"`
	QrCode     string `json:"qr_code"`
	InputA     string `json:"inputA"`
	InputB     string `json:"inputB"`
}

type SponsorDetailReq struct {
	Id int64 `json:"id"`
}

type SponsorDetailResp struct {
	Id         int64  `json:"id"`
	UserId     int64  `json:"userId"`
	Type       int64  `json:"type"`
	AppletType int64  `json:"appletType"`
	IsShow     int64  `json:"isShow"`
	Name       string `json:"name"`
	Desc       string `json:"desc"`
	Avatar     string `json:"avatar"`
	QrCode     string `json:"qr_code"`
	InputA     string `json:"inputA"`
	InputB     string `json:"inputB"`
}

type SponsorListReq struct {
	Page     int64 `json:"page,range=[1:]"`
	PageSize int64 `json:"pageSize,range=[0:]"`
}

type SponsorListResp struct {
	List []Sponsor `json:"list"`
}

type UpdateSponsorReq struct {
	Id         int64  `json:"id"`
	UserId     int64  `json:"userId"`
	Type       int64  `json:"type"`
	AppletType int64  `json:"appletType"`
	IsShow     int64  `json:"isShow"`
	Name       string `json:"name"`
	Desc       string `json:"desc"`
	Avatar     string `json:"avatar"`
	QrCode     string `json:"qr_code"`
	InputA     string `json:"inputA"`
	InputB     string `json:"inputB"`
}

type UpdateSponsorResp struct {
}

type User struct {
	Id           int64   `json:"id"`
	Mobile       string  `json:"mobile"`
	NickName     string  `json:"nickName"`
	Sex          int64   `json:"sex"`
	Avatar       string  `json:"avatar"`
	Info         string  `json:"info"`
	IsAdmin      int64   `json:"isAdmin"`
	Signature    string  `json:"signature"`
	LocationName string  `json:"locationName"`
	Longitude    float64 `json:"longitude"`
	Latitude     float64 `json:"latitude"`
}

type UserInfoReq struct {
}

type UserInfoResp struct {
	UserInfo User `json:"userInfo"`
}

type UserUpdateReq struct {
	Nickname  string  `json:"nickname"`
	Sex       string  `json:"sex"`
	Avator    string  `json:"avator"`
	Info      string  `json:"info"`
	Signature string  `json:"signature"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type UserUpdateResp struct {
}

type WXMiniAuthReq struct {
	Code          string `json:"code"`
	IV            string `json:"iv"`
	EncryptedData string `json:"encryptedData"`
	Nickname      string `json:"nickname,optional"`
	Avator        string `json:"avator,optional"`
}

type WXMiniAuthResp struct {
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
	RefreshAfter int64  `json:"refreshAfter"`
}

type SponsorDelReq struct {
	Id int64 `json:"id" validate:"required"`
}

type SponsorDelResp struct {
}
