package alipass

import (
	"encoding/json"
	"fmt"
	"github.com/alipay/alipay-sdk/api"
	"github.com/alipay/alipay-sdk/api/request"
	"github.com/alipay/alipay-sdk/api/response"
)

// Alipass用户识别方式枚举类型
const (
	TRADE  = "1" // 支付宝交易
	USERID = "2" // 支付宝用户ID
	MOBILE = "3" // 支付宝用户绑定手机号
	OPENID = "4" // 支付宝公众号开放ID
)

// AlipassTransferService 卡券服务类
type AlipassTransferService struct {
}

// AddByTemplate 模板方式添加卡券
func (a *AlipassTransferService) AddByTemplate(r *AddTplRequest) (*response.AlipayPassTplContentAddResponse, error) {

	// 请求对象
	contentAddRequest := &request.AlipayPassTplContentAddRequest{}

	// client
	c := &api.DefaultAlipayClient{
		AppId:     r.AppId,
		ServerURL: r.AlipayApiUrl,
		PrivKey:   r.PrivateKeyData,
		Charset:   "UTF-8",
	}

	contentAddRequest.RecognitionType = r.UserType

	// to json
	tplParamsBytes, err := json.Marshal(r.TemplateParamValuePair)
	if err != nil {
		return nil, fmt.Errorf("%s", "TemplateParamValuePair格式错误，需为json格式")
	}
	recognitionInfoBytes, err := json.Marshal(r.UserTypeParams)
	if err != nil {
		return nil, fmt.Errorf("%s", "UserTypeParams格式错误，需为json格式")
	}

	// params
	contentAddRequest.RecognitionInfo = string(recognitionInfoBytes)
	contentAddRequest.TplParams = string(tplParamsBytes)
	contentAddRequest.TplId = r.TemplateId
	// 执行
	resp, err := c.Execute(contentAddRequest)

	if err != nil {
		return nil, err
	}

	tplAddResponse := resp.(*response.AlipayPassTplContentAddResponse)
	return tplAddResponse, nil
}

// AddByTemplate 模板方式添加卡券
func (a *AlipassTransferService) UpdateByTemplate() {

}

// AddByTemplate 模板方式添加卡券
func (a *AlipassTransferService) CreateByTemplate() {

}
