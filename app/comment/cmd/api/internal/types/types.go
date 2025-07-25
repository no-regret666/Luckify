// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package types

type Comment struct {
	Id          int64  `json:"id"`
	UserId      int64  `json:"userId"`      // 用户id
	LotteryId   int64  `json:"lotteryId"`   // 抽奖id
	PrizeName   string `json:"prizeName"`   // 奖品名称
	Content     string `json:"content"`     // 晒单评论内容
	Pics        string `json:"pics"`        // 晒单评论图片
	PraiseCount int64  `json:"praiseCount"` // 点赞数量
	CreateTime  int64  `json:"createTime"`  // 创建时间
	UpdateTime  int64  `json:"updateTime"`  // 更新时间
}

type CommentAddReq struct {
	LotteryId int64  `json:"lotteryId"`
	PrizeName string `json:"prizeName"`
	Content   string `json:"content"`
	Pics      string `json:"pics"`
}

type CommentAddResp struct {
}

type CommentDelReq struct {
	Id int64 `json:"id"`
}

type CommentDelResp struct {
}

type CommentDetailReq struct {
	Id int64 `json:"id"`
}

type CommentDetailResp struct {
	Comment Comment `json:"comment"`
}

type CommentListReq struct {
	LastId   int64 `json:"lastId"`
	PageSize int64 `json:"pageSize"`
	Sort     int64 `json:"sort"`
}

type CommentListResp struct {
	List []Comments `json:"list"`
}

type CommentPraiseReq struct {
	Id int64 `json:"id"`
}

type CommentPraiseResp struct {
}

type CommentUpdateReq struct {
	Id        int64  `json:"id"`
	UserId    int64  `json:"userId"`
	LotteryId int64  `json:"lotteryId"`
	PrizeName string `json:"prizeName"`
	Content   string `json:"content"`
	Pics      string `json:"pics"`
}

type CommentUpdateResp struct {
}

type Comments struct {
	Id          int64  `json:"id"`
	UserId      int64  `json:"userId"`      // 用户id
	LotteryId   int64  `json:"lotteryId"`   // 抽奖id
	PrizeName   string `json:"prizeName"`   // 奖品名称
	Content     string `json:"content"`     // 晒单评论内容
	Pics        string `json:"pics"`        // 晒单评论图片
	PraiseCount int64  `json:"praiseCount"` // 点赞数量
	CreateTime  int64  `json:"createTime"`  // 创建时间
	UpdateTime  int64  `json:"updateTime"`  // 更新时间
	User        User   `json:"user"`        // 用户信息
	IsPraise    int64  `json:"isPraise"`    // 是否点赞
}

type GetCommentLastIdReq struct {
}

type GetCommentLastIdResp struct {
	LastId int64 `json:"lastId"`
}

type User struct {
	Id       int64  `json:"id"`
	Mobile   string `json:"Mobile"`
	Nickname string `json:"nickname"`
	Sex      int64  `json:"sex"`
	Avatar   string `json:"avatar"`
	Info     string `json:"info"`
	IsAdmin  int64  `json:"isAdmin"`
}
