package request

import (
	"github.com/alipay/alipay-sdk/api/response"
)

// AlipayRequest request接口
type AlipayRequest interface {

	// 方法名称
	GetApiMethod() string

	// 版本号
	GetApiVersion() string

	// 应用参数
	// 包括biz_content、自定义的参数
	GetTextParams() map[string]string

	// 每一个request必须绑定一个response对象
	GetResponse() *response.AlipayResponse
}

// AlipayMobilePublicMessageCustomSendRequest
// API: alipay.mobile.public.message.custom.send request
type AlipayMobilePublicMessageCustomSendRequest struct {
	BizContent string
}

func (r *AlipayMobilePublicMessageCustomSendRequest) GetApiMethod() string {
	return "alipay.mobile.public.message.custom.send"
}

func (r *AlipayMobilePublicMessageCustomSendRequest) GetTextParams() map[string]string {
	params := make(map[string]string)
	params["biz_content"] = r.BizContent
	//TODO 提供一个用户设置参数的接口
	//append(params,userParams)
	return params
}

func (r *AlipayMobilePublicMessageCustomSendRequest) GetResponse() *response.AlipayResponse {
	return nil
}

func (r *AlipayMobilePublicMessageCustomSendRequest) GetApiVersion() string {
	return "1.0"
}
