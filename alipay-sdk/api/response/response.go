package response

import (
	"strings"
)

// AlipayResponse response接口
type AlipayResponse interface {

	// 判断是否成功
	IsSuccess() bool
	// 接口名称
	ToStr() string
}

type BaseResponse struct {
	Code    string `align:"code"`
	Msg     string `align:"msg"`
	SubCode string `align:"sub_code"`
	SubMsg  string `align:"sub_msg"`
	Name    string
}

func (r *BaseResponse) IsSuccess() bool {
	// sub_code如果为空，表明执行成功
	return strings.EqualFold("", r.SubCode)
}

func (r *BaseResponse) ToStr() string {
	return r.Name
}

// AlipayMobilePublicMessageCustomSendResponse
// 与AlipayMobilePublicMessageCustomSendRequest关联
type AlipayMobilePublicMessageCustomSendResponse struct {
	BaseResponse
}

// AlipaySystemOauthTokenResponse
// refer AlipaySystemOauthTokenRequest
type AlipaySystemOauthTokenResponse struct {
	BaseResponse
	AccessToken  string `align:"access_token"`
	AlipayUserId string `align:"alipay_user_id"`
	ExpiresIn    string `align:"expires_in"`
	ReExpiresIn  string `align:"re_expires_in"`
	RefreshToken string `align:"refresh_token"`
}

// AlipayPassTplContentAddResponse
// refer AlipayPassTplContentAddRequest
type AlipayPassTplContentAddResponse struct {
	BaseResponse
	BizResult string `align:"biz_result"`
	ErrorCode string `align:"error_code"`
	Success   string `align:"success"` //T-成功；F-失败
}

// AlipayPassSyncUpdateResponse
// refer AlipayPassSyncUpdateRequest
type AlipayPassSyncUpdateResponse struct {
	BaseResponse
	BizResult string `align:"biz_result"`
	ErrorCode string `align:"error_code"`
	Success   string `align:"success"` //T-成功；F-失败
}
