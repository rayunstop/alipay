package request

import (
	"github.com/alipay/alipay-sdk/api/response"
)

// AlipayPassTplContentAddRequest
// api: alipay.pass.tpl.content.add request
type AlipayPassTplContentAddRequest struct {

	//支付宝用户识别信息：
	//当recognition_type=1时， recognition_info={“partner_id”:”2088102114633762”,“out_trade_no”:”1234567”}；
	//当recognition_type=3时，recognition_info={“mobile”:”136XXXXXXXX“}
	//当recognition_type=4时， recognition_info={“open_id”:”afbd8d9bb12fc02c5094d8ea89d1fae8“}
	RecognitionInfo string

	//Alipass添加对象识别类型【1--订单信息;3--支付宝用户绑定手机号；4--支付宝OpenId;】
	RecognitionType string

	//支付宝pass模版ID
	TplId string

	//模版动态参数信息【支付宝pass模版参数键值对JSON字符串】
	TplParams string

	//add user-defined text parameters
	UdfParams map[string]string
}

func (r *AlipayPassTplContentAddRequest) GetApiMethod() string {
	return "alipay.pass.tpl.content.add"
}

func (r *AlipayPassTplContentAddRequest) GetTextParams() map[string]string {
	params := make(map[string]string)
	params["recognition_info"] = r.RecognitionInfo
	params["recognition_type"] = r.RecognitionType
	params["tpl_id"] = r.TplId
	params["tpl_params"] = r.TplParams
	//TODO 提供一个用户设置参数的接口
	//utils.putAll(params,userParams)
	return params
}

func (r *AlipayPassTplContentAddRequest) GetResponse() response.AlipayResponse {
	resp := new(response.AlipayPassTplContentAddResponse)
	resp.Name = "AlipayPassTplContentAddResponse"
	return resp
}

func (r *AlipayPassTplContentAddRequest) GetApiVersion() string {
	return "1.0"
}
