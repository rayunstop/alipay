package request

import (
	"github.com/z-ray/alipay/api/response"
)

// AlipayMobilePublicMessageCustomSendRequest
// API: alipay.mobile.public.message.caustom.send request
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
	//utils.putAll(params,userParams)
	return params
}

func (r *AlipayMobilePublicMessageCustomSendRequest) GetResponse() response.AlipayResponse {
	resp := new(response.AlipayMobilePublicMessageCustomSendResponse)
	// 类名，在获取结果时有用
	resp.Name = "AlipayMobilePublicMessageCustomSendResponse"
	return resp
}

func (r *AlipayMobilePublicMessageCustomSendRequest) GetApiVersion() string {
	return "1.0"
}
