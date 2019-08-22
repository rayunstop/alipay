package request

import (
	"github.com/z-ray/alipay/api/response"
	"github.com/z-ray/alipay/api/utils"
)

// AlipayPassSyncUpdateRequest
// api: alipay.pass.sync.update request
type AlipayPassTplContentUpdateRequest struct {
	ChannelId    string //代理商代替商户发放卡券后，再代替商户更新卡券时，此值为商户的pid/appid；商户自己发券时，此值为空或者商户appId
	TplParams    string //模版动态参数信息【支付宝pass模版参数键值对JSON字符串】
	Pass         string //需要修改的pass信息，可以更新全部pass信息，也可以斤更新某一节点。pass信息中的pass.json中的数据格式，如果不需要更新该属性值，设置为null即可。
	SerialNumber string //alipass唯一标识
	Status       string //alipass状态，目前仅支持CLOSED及USED两种数据。status为USED时，verify_type即为核销时的核销方式。
	VerifyCode   string //核销码串值【当状态变更为USED时，建议传入】
	VerifyType   string //核销方式，目前支持：wave（声波方式）、qrcode（二维码方式）、barcode（条码方式）、input（文本方式，即手工输入方式）。pass和verify_type不能同时为空
	UdfParams    map[string]string
}

func (r *AlipayPassTplContentUpdateRequest) GetApiMethod() string {
	return "alipay.pass.tpl.content.update"
}

func (r *AlipayPassTplContentUpdateRequest) GetTextParams() map[string]string {
	params := make(map[string]string)
	params["channel_id"] = r.ChannelId
	params["tpl_params"] = r.TplParams
	params["pass"] = r.Pass
	params["serial_number"] = r.SerialNumber
	params["status"] = r.Status
	params["verify_code"] = r.VerifyCode
	params["verify_type"] = r.VerifyType
	//utils.putAll(params,userParams)
	if len(r.UdfParams) > 0 {
		utils.PutAll(params, r.UdfParams)
	}
	return params
}

func (r *AlipayPassTplContentUpdateRequest) GetResponse() response.AlipayResponse {
	resp := new(response.AlipayPassTplContentUpdateResponse)
	resp.Name = "AlipayPassTplContentUpdateResponse"
	return resp
}

func (r *AlipayPassTplContentUpdateRequest) GetApiVersion() string {
	return "1.0"
}

// PutOtherTextParam 添加附加参数
func (r *AlipayPassTplContentUpdateRequest) PutOtherTextParam(k, v string) {

	if r.UdfParams == nil {
		r.UdfParams = make(map[string]string)
	}
	r.UdfParams[k] = v
}
