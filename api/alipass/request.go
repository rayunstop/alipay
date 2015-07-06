package alipass

// BaseRequest 请求入参基类
type BaseRequest struct {
	AlipayApiUrl   string
	AppId          string
	PrivateKeyData string
}

// AddTplRequest 添加模板请求对象
type AddTplRequest struct {
	BaseRequest
	TemplateId             string
	TemplateUserId         string
	TemplateParamValuePair map[string]string
	UserTypeParams         map[string]string
	UserType               string
}

// UpdAlipssRequest alipass更新请求入参对象
type UpdAlipssRequest struct {
	BaseRequest
	SerialNumber string
	Pass         string
	Status       string
	ChannelId    string
	VerifyCode   string
	VerifyType   string
	ExtInfo      map[string]string
}

// UpdTplAlipssRequest 模板方式更新卡券
type UpdTplAlipssRequest struct {
	BaseRequest
	SerialNumber           string
	Status                 string
	ChannelId              string
	VerifyCode             string
	VerifyType             string
	TemplateParamValuePair map[string]string
}
