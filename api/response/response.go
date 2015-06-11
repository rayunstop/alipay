package response

import (
	"fmt"
	"github.com/z-ray/log"
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
	Code    interface{} `json:"code"`
	Msg     string      `json:"msg"`
	SubCode string      `json:"sub_code"`
	SubMsg  string      `json:"sub_msg"`
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
	// code may be string
	str, ok := r.Code.(string)
	if ok {
		return str
	}
	// code may be float64
	integer, ok := r.Code.(float64)
	if ok {
		return fmt.Sprintf("%0.f", integer)
	}
	// both not
	log.Warnf("alipay response code type:%s", r.Code)
	return ""
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
	AccessToken  string `json:"access_token"`
	AlipayUserId string `json:"alipay_user_id"`
	ExpiresIn    int64  `json:"expires_in"`
	ReExpiresIn  int64  `json:"re_expires_in"`
	RefreshToken string `json:"refresh_token"`
}

// AlipayPassTplContentAddResponse
// refer AlipayPassTplContentAddRequest
type AlipayPassTplContentAddResponse struct {
	BaseResponse
	BizResult string `json:"biz_result"`
	ErrorCode string `json:"error_code"`
	Success   string `json:"success"` //T-成功；F-失败
}

// AlipayPassSyncUpdateResponse
// refer AlipayPassSyncUpdateRequest
type AlipayPassSyncUpdateResponse struct {
	BaseResponse
	BizResult string `json:"biz_result"`
	ErrorCode string `json:"error_code"`
	Success   bool   `json:"success"` //T-成功；F-失败
}

// AlipayMobilePublicGisGetResponse
// refer AlipayMobilePublicGisGetRequest
type AlipayMobilePublicGisGetResponse struct {
	BaseResponse
	Accuracy  string `json:"accuracy"`
	City      string `json:"city"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	Province  string `json:"province"`
}

// AlipayPassTplContentUpdateResponse
// refer AlipayPassTplContentUpdateRequest
type AlipayPassTplContentUpdateResponse struct {
	BaseResponse
	Result    string `json:"result"`
	ErrorCode string `json:"error_code"`
	Success   bool   `json:"success"`
}
