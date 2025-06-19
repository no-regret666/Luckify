package response

import (
	"Luckify/common/xerr"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
	"net/http"
)

// HttpResult http返回
func HttpResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	if err == nil {
		// 成功返回
		r := Success(resp)
		httpx.WriteJson(w, http.StatusOK, r)
	} else {
		//错误返回
		errCode := xerr.SERVER_COMMOM_ERROR
		errMsg := "服务器开小差啦，稍后再来试一试"

		causeErr := errors.Cause(err)                // err类型
		if e, ok := causeErr.(*xerr.CodeError); ok { //自定义错误类型
			//自定义CodeError
			errCode = e.GetErrCode()
			errMsg = e.GetErrMsg()
		} else {
			// grpc err错误
			if gStatus, ok := status.FromError(causeErr); ok {
				grpcCode := uint32(gStatus.Code())
				if xerr.IsCodeErr(grpcCode) {
					// 区分自定义错误跟系统底层、db等错误，底层、db错误不能返回给前端
					errCode = grpcCode
					errMsg = gStatus.Message()
				}
			}
		}

		logx.WithContext(r.Context()).Errorf(" [API-ERR] : %+v", err)

		httpx.WriteJson(w, http.StatusBadRequest, Error(errCode, errMsg))
	}
}

// AuthHttpResult 授权的http方法
func AuthHttpResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	if err == nil {
		// 成功返回
		r := Success(resp)
		httpx.WriteJson(w, http.StatusOK, r)
	} else {
		// 错误返回
		errCode := xerr.SERVER_COMMOM_ERROR
		errMsg := "服务器开小差啦，稍后再来试一试"

		causeErr := errors.Cause(err)                //err类型
		if e, ok := causeErr.(*xerr.CodeError); ok { //自定义错误类型
			//自定义CodeError
			errCode = e.GetErrCode()
			errMsg = e.GetErrMsg()
		} else {
			if gStatus, ok := status.FromError(causeErr); ok { // grpc err编程
				grpcCode := uint32(gStatus.Code())
				if xerr.IsCodeErr(grpcCode) { // 区分自定义错误跟系统底层、db等错误，底层、db错误不能返回给前端
					errCode = grpcCode
					errMsg = gStatus.Message()
				}
			}
		}
		logx.WithContext(r.Context()).Errorf(" [GATEWAY-ERR] : %+v", err)

		httpx.WriteJson(w, http.StatusUnauthorized, Error(errCode, errMsg))
	}
}

// ParamErrorResult 参数错误返回
func ParamErrorResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	errMsg := fmt.Sprintf("%s,%s", xerr.MapErrMsg(xerr.REQUEST_PARAM_ERROR), err.Error())
	httpx.WriteJson(w, http.StatusBadRequest, Error(xerr.REQUEST_PARAM_ERROR, errMsg))
}
