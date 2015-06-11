package alipass

import (
	"encoding/json"
	"fmt"
	"github.com/z-ray/alipay"
	"github.com/z-ray/alipay/api/request"
	"github.com/z-ray/alipay/api/response"
)

// AlipassTransferService 卡券服务类
type AlipassTransferService struct {
}

// AddByTemplate 模板方式添加卡券
func (a *AlipassTransferService) AddByTemplate(r *AddTplRequest) (*response.AlipayPassTplContentAddResponse, error) {

	// 请求对象
	contentAddRequest := &request.AlipayPassTplContentAddRequest{}

	// client
	c := &alipay.DefaultAlipayClient{
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

// UpdateTplAlipass 模板方式更新卡券
func (a *AlipassTransferService) UpdateTplAlipass(r *UpdTplAlipssRequest) (*response.AlipayPassTplContentUpdateResponse, error) {

	// 验证参数合法性

	// 请求对象
	passUpdTplRequest := &request.AlipayPassTplContentUpdateRequest{}

	// client
	c := &sdk.DefaultAlipayClient{
		AppId:     r.AppId,
		ServerURL: r.AlipayApiUrl,
		PrivKey:   r.PrivateKeyData,
		Charset:   "UTF-8",
	}

	passUpdTplRequest.SerialNumber = r.SerialNumber
	passUpdTplRequest.ChannelId = r.ChannelId
	passUpdTplRequest.VerifyCode = r.VerifyCode
	passUpdTplRequest.VerifyType = r.VerifyType
	passUpdTplRequest.Status = r.Status

	if len(r.TemplateParamValuePair) > 0 {
		tplParamsBytes, err := json.Marshal(r.TemplateParamValuePair)
		if err == nil {
			passUpdTplRequest.TplParams = string(tplParamsBytes)
		}
	}

	passUpdTplRequest.PutOtherTextParam("operator_id", r.AppId)

	resp, err := c.Execute(passUpdTplRequest)
	if err != nil {
		return nil, err
	}

	passUpdTplResponse := resp.(*response.AlipayPassTplContentUpdateResponse)
	return passUpdTplResponse, nil
}

// UpdateByTemplate 更新卡券状态
func (a *AlipassTransferService) UpdateAlipass(r *UpdAlipssRequest) (*response.AlipayPassSyncUpdateResponse, error) {

	// 验证参数合法性

	// 请求对象
	passUpdRequest := &request.AlipayPassSyncUpdateRequest{}

	// client
	c := &alipay.DefaultAlipayClient{
		AppId:     r.AppId,
		ServerURL: r.AlipayApiUrl,
		PrivKey:   r.PrivateKeyData,
		Charset:   "UTF-8",
	}

	passUpdRequest.SerialNumber = r.SerialNumber
	passUpdRequest.Pass = r.Pass
	passUpdRequest.ChannelId = r.ChannelId
	passUpdRequest.VerifyCode = r.VerifyCode
	passUpdRequest.VerifyType = r.VerifyType
	passUpdRequest.Status = r.Status

	if len(r.ExtInfo) > 0 {
		extInfoBytes, err := json.Marshal(r.ExtInfo)
		if err == nil {
			passUpdRequest.ExtInfo = string(extInfoBytes)
		}
	}

	passUpdRequest.PutOtherTextParam("operator_id", r.AppId)

	resp, err := c.Execute(passUpdRequest)
	if err != nil {
		return nil, err
	}

	passUpdResponse := resp.(*response.AlipayPassSyncUpdateResponse)
	return passUpdResponse, nil
}

// CreateByTemplate 创建卡券
func (a *AlipassTransferService) CreateByTemplate() {

}
