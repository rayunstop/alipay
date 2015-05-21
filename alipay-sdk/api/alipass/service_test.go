package alipass

import (
	"testing"
)

const (
	AppId      = "2015040200041603"
	PrivateKey = "MIICdwIBADANBgkqhkiG9w0BAQEFAASCAmEwggJdAgEAAoGBAKK0PXoLKnBkgtOl0kvyc9X2tUUdh/lRZr9RE1frjr2ZtAulZ+Moz9VJZFew1UZIzeK0478obY/DjHmD3GMfqJoTguVqJ2MEg+mJ8hJKWelvKLgfFBNliAw+/9O6Jah9Q3mRzCD8pABDEHY7BM54W7aLcuGpIIOa/qShO8dbXn+FAgMBAAECgYA8+nQ380taiDEIBZPFZv7G6AmT97doV3u8pDQttVjv8lUqMDm5RyhtdW4n91xXVR3ko4rfr9UwFkflmufUNp9HU9bHIVQS+HWLsPv9GypdTSNNp+nDn4JExUtAakJxZmGhCu/WjHIUzCoBCn6viernVC2L37NL1N4zrR73lSCk2QJBAPb/UOmtSx+PnA/mimqnFMMP3SX6cQmnynz9+63JlLjXD8rowRD2Z03U41Qfy+RED3yANZXCrE1V6vghYVmASYsCQQCoomZpeNxAKuUJZp+VaWi4WQeMW1KCK3aljaKLMZ57yb5Bsu+P3odyBk1AvYIPvdajAJiiikRdIDmi58dqfN0vAkEAjFX8LwjbCg+aaB5gvsA3t6ynxhBJcWb4UZQtD0zdRzhKLMuaBn05rKssjnuSaRuSgPaHe5OkOjx6yIiOuz98iQJAXIDpSMYhm5lsFiITPDScWzOLLnUR55HL/biaB1zqoODj2so7G2JoTiYiznamF9h9GuFC2TablbINq80U2NcxxQJBAMhw06Ha/U7qTjtAmr2qAuWSWvHU4ANu2h0RxYlKTpmWgO0f47jCOQhdC3T/RK7f38c7q8uPyi35eZ7S1e/PznY="
	OpenApiUrl = "https://openapi.alipay.com/gateway.do"
	UserId     = "+0YuKZBkIc1cEDklW4gMk5qKo7ILToCxhO4skMe5bRK-S4-HyunMvYTqiRYEUD+U01"
)

func TestAddByTemplate(t *testing.T) {

	openId, serialNumber := "20881011315239463742107232815092", "20150521140000"
	// 卡券参数
	paramValuePair := make(map[string]string)
	paramValuePair["qrcode"] = serialNumber
	paramValuePair["serialNumber"] = serialNumber
	paramValuePair["channelID"] = AppId
	paramValuePair["webServiceUrl"] = " "

	// 用户参数
	userParams := make(map[string]string)
	userParams["open_id"] = openId

	addReq := &AddTplRequest{}
	addReq.AlipayApiUrl = OpenApiUrl
	addReq.AppId = AppId
	addReq.PrivateKeyData = PrivateKey
	addReq.TemplateParamValuePair = paramValuePair
	addReq.UserTypeParams = userParams
	addReq.UserType = OPENID
	addReq.TemplateId = "2015051210573746242893580"

	alipassService := &AlipassTransferService{}
	resp, err := alipassService.AddByTemplate(addReq)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Logf("%+v", resp)
}
