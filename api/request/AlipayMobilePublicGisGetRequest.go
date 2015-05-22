package request

import (
	"github.com/rui2014/alipay/api/response"
)

// AlipayMobilePublicGisGetRequest
// api: alipay.mobile.public.gis.get request
type AlipayMobilePublicGisGetRequest struct {
	BizContent string
}

func (r *AlipayMobilePublicGisGetRequest) GetApiMethod() string {
	return "alipay.mobile.public.gis.get"
}

func (r *AlipayMobilePublicGisGetRequest) GetTextParams() map[string]string {
	params := make(map[string]string)
	params["biz_content"] = r.BizContent
	//TODO 提供一个用户设置参数的接口
	//utils.putAll(params,userParams)
	return params
}

func (r *AlipayMobilePublicGisGetRequest) GetResponse() response.AlipayResponse {
	resp := new(response.AlipayMobilePublicGisGetResponse)
	// 类名，在获取结果时有用
	resp.Name = "AlipayMobilePublicGisGetResponse"
	return resp
}

func (r *AlipayMobilePublicGisGetRequest) GetApiVersion() string {
	return "1.0"
}
