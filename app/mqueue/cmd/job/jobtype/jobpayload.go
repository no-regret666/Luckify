package jobtype

type WxMiniProgramNotifyUserPayload struct {
	MsgType  int
	OpenId   string
	PageAddr string
	Data     string
}
