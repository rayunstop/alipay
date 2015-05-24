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
	// 保留body
	SetBody(body string)
	// code
	GetCode() string
	// subCode
	GetSubCode() string
	// msg
	GetMsg() string
}

type BaseResponse struct {
	Code    string `align:"code"`
	Msg     string `align:"msg"`
	SubCode string `align:"sub_code"`
	SubMsg  string `align:"sub_msg"`
	Name    string
	Body    string
}

func (r *BaseResponse) IsSuccess() bool {
	// sub_code如果为空，表明执行成功
	return strings.EqualFold("", r.SubCode)
}

// ToStr 输出类名，用于动态获取支付宝返回值key
func (r *BaseResponse) ToStr() string {
	return r.Name
}

// SetBody 保存请求结果
func (r *BaseResponse) SetBody(body string) {
	r.Body = body
}

// GetCode
func (r *BaseResponse) GetCode() string {
	return r.Code
}

// GetSubCode
func (r *BaseResponse) GetSubCode() string {
	return r.SubCode
}

// GetMsg
func (r *BaseResponse) GetMsg() string {
	return r.Msg
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

// AlipayMobilePublicGisGetResponse
// refer AlipayMobilePublicGisGetRequest
type AlipayMobilePublicGisGetResponse struct {
	BaseResponse
	Accuracy  string `align:"accuracy"`
	City      string `align:"city"`
	Latitude  string `align:"latitude"`
	Longitude string `align:"longitude"`
	Province  string `align:"province"`
}
