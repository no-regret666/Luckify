package xerr

var message map[uint32]string

func init() {
	message = make(map[uint32]string)
	message[OK] = "SUCCESS"
	message[SERVER_COMMOM_ERROR] = "服务器开小差啦，稍后再来试一下"
	message[REQUEST_PARAM_ERROR] = "参数错误"
	message[TOKEN_EXPIRE_ERROR] = "token失效，请重新登录"
	message[TOKEN_GENERATE_ERROR] = "生成token失败"
	message[DB_ERROR] = "数据库繁忙，请稍后再试"
	message[DB_UPDATE_AFFECTED_ZERO_ERROR] = "更新数据影响行数为0"
	message[DB_ERROR_NOT_FOUND] = "数据未找到"
	message[DB_TRANSACTION_ERROR] = "数据库事务错误"

	// 抽奖模块
	message[DB_GETLASTID_ERROR] = "获取最后一个抽奖id失败"
	message[DB_GET_LOTTERY_LIST_ERROR] = "获取抽奖列表失败"
	message[DB_GET_WIN_LOTTERY_LIST_ERROR] = "获取中奖抽奖列表失败"
	message[DB_GET_PRIZE_LIST_ERROR] = "获取奖品列表失败"
	message[DB_GET_USER_WIN_LOTTERY_LIST_ERROR] = "获取用户中奖抽奖列表失败"
	message[DB_GET_ALL_USER_LOTTERY_LIST_ERROR] = "获取用户所有抽奖列表失败"
	message[DB_GET_USER_CREATED_LOTTERY_LIST_ERROR] = "获取用户创建抽奖列表失败"
	message[DB_GET_PRIZE_ERROR] = "获取奖品失败"
	message[DB_GET_PARTICIPATED_BY_LOTTERY_ID] = "获取用户参与抽奖信息失败"
	message[DB_GET_PARTICIPATIONS_COUNT] = "获取用户参与抽奖数量失败"
	message[DB_GET_CREATED_COUNT_BY_USER_ID] = "获取用户创建抽奖数量失败"
	message[DB_GET_WON_COUNT_BY_USER_ID] = "获取用户中奖数量失败"
	message[DB_GET_PARTICIPATION_COUNT_BY_USER_ID] = "获取用户参与抽奖数量失败"
	message[DB_GET_PRIZE_BY_BY_LOTTERY_ID] = "获取抽奖奖品失败"
	message[DB_CHECK_IS_PARTICIPATED] = "检查用户是否参与抽奖失败"
	message[DB_PARTICIPATE_LOTTERY] = "参与抽奖失败"
	message[DB_GET_PENDING_LOTTERY_LIST_FAIL] = "获取待开奖抽奖列表失败"
	message[DB_GET_PARTICIPATORS_COUNT_FAIL] = "获取参与抽奖人数失败"
	message[DB_ANNOUNCE_LOTTERY_FAIL] = "开奖失败"
	message[ANNOUNCE_LOTTERY_FAIL] = "开奖失败"
	message[DB_UPDATE_STATUS_ERROR] = "更新抽奖状态失败"
	message[DB_UPDATE_WINNER_ERROR] = "更新中奖者失败"
	message[DB_FIND_PARTICIPATION_BY_LOTTERY_ID_AND_USER_ID_ERR] = "查找抽奖参与者失败"
	message[DB_FIND_LOTTERY_BY_ID_ERR] = "查找抽奖失败"
	message[DB_FIND_PRIZES_BY_LOTTERY_ID_ERR] = "查找抽奖奖品失败"
	message[DB_DECR_PRIZE_STOCK_ERR] = "减少奖品库存失败"
}

func MapErrMsg(errcode uint32) string {
	if msg, ok := message[errcode]; ok {
		return msg
	} else {
		return "服务器开小差啦，稍后再来试一试"
	}
}

func IsCodeErr(errcode uint32) bool {
	if _, ok := message[errcode]; ok {
		return true
	} else {
		return false
	}
}
