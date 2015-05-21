package alipass

const (
	url = "https://openapi.alipay.com/gateway.do"
)

// BaseRequest 请求入参基类
type BaseRequest struct {
	AlipayApiUrl   string
	AppId          string
	PrivateKeyData string
}

type AddTplRequest struct {
	BaseRequest
	TemplateId             string
	TemplateUserId         string
	TemplateParamValuePair map[string]string
	UserTypeParams         map[string]string
	UserType               string
}
