package alipay

import (
	"github.com/rui2014/alipay/api/request"
	"log"
	"testing"
)

var alipayClient *DefaultAlipayClient

// var userId = "kquy1L6pBsGPHDYCYeBhLJqKo7ILToCxhO4skMe5bRK-S4-HyunMvYTqiRYEUD+U01"

var userId = "20881011315239463742107232815092"

func init() {
	// 飞森
	// alipayClient = &DefaultAlipayClient{
	// 	AppId:       "2014121600020041",
	// 	ServerURL:   "https://openapi.alipay.com/gateway.do",
	// 	PrivKey:     "MIICdwIBADANBgkqhkiG9w0BAQEFAASCAmEwggJdAgEAAoGBAKK0PXoLKnBkgtOl0kvyc9X2tUUdh/lRZr9RE1frjr2ZtAulZ+Moz9VJZFew1UZIzeK0478obY/DjHmD3GMfqJoTguVqJ2MEg+mJ8hJKWelvKLgfFBNliAw+/9O6Jah9Q3mRzCD8pABDEHY7BM54W7aLcuGpIIOa/qShO8dbXn+FAgMBAAECgYA8+nQ380taiDEIBZPFZv7G6AmT97doV3u8pDQttVjv8lUqMDm5RyhtdW4n91xXVR3ko4rfr9UwFkflmufUNp9HU9bHIVQS+HWLsPv9GypdTSNNp+nDn4JExUtAakJxZmGhCu/WjHIUzCoBCn6viernVC2L37NL1N4zrR73lSCk2QJBAPb/UOmtSx+PnA/mimqnFMMP3SX6cQmnynz9+63JlLjXD8rowRD2Z03U41Qfy+RED3yANZXCrE1V6vghYVmASYsCQQCoomZpeNxAKuUJZp+VaWi4WQeMW1KCK3aljaKLMZ57yb5Bsu+P3odyBk1AvYIPvdajAJiiikRdIDmi58dqfN0vAkEAjFX8LwjbCg+aaB5gvsA3t6ynxhBJcWb4UZQtD0zdRzhKLMuaBn05rKssjnuSaRuSgPaHe5OkOjx6yIiOuz98iQJAXIDpSMYhm5lsFiITPDScWzOLLnUR55HL/biaB1zqoODj2so7G2JoTiYiznamF9h9GuFC2TablbINq80U2NcxxQJBAMhw06Ha/U7qTjtAmr2qAuWSWvHU4ANu2h0RxYlKTpmWgO0f47jCOQhdC3T/RK7f38c7q8uPyi35eZ7S1e/PznY=",
	// 	Format:      "json",
	// 	ConTimeOut:  2000,
	// 	ReadTimeOut: 2000,
	// 	SignType:    "RSA",
	// 	Charset:     "GBK",
	// }
	alipayClient = &DefaultAlipayClient{
		AppId:       "2015040200041603",
		ServerURL:   "https://openapi.alipay.com/gateway.do",
		PrivKey:     "MIICdwIBADANBgkqhkiG9w0BAQEFAASCAmEwggJdAgEAAoGBAKK0PXoLKnBkgtOl0kvyc9X2tUUdh/lRZr9RE1frjr2ZtAulZ+Moz9VJZFew1UZIzeK0478obY/DjHmD3GMfqJoTguVqJ2MEg+mJ8hJKWelvKLgfFBNliAw+/9O6Jah9Q3mRzCD8pABDEHY7BM54W7aLcuGpIIOa/qShO8dbXn+FAgMBAAECgYA8+nQ380taiDEIBZPFZv7G6AmT97doV3u8pDQttVjv8lUqMDm5RyhtdW4n91xXVR3ko4rfr9UwFkflmufUNp9HU9bHIVQS+HWLsPv9GypdTSNNp+nDn4JExUtAakJxZmGhCu/WjHIUzCoBCn6viernVC2L37NL1N4zrR73lSCk2QJBAPb/UOmtSx+PnA/mimqnFMMP3SX6cQmnynz9+63JlLjXD8rowRD2Z03U41Qfy+RED3yANZXCrE1V6vghYVmASYsCQQCoomZpeNxAKuUJZp+VaWi4WQeMW1KCK3aljaKLMZ57yb5Bsu+P3odyBk1AvYIPvdajAJiiikRdIDmi58dqfN0vAkEAjFX8LwjbCg+aaB5gvsA3t6ynxhBJcWb4UZQtD0zdRzhKLMuaBn05rKssjnuSaRuSgPaHe5OkOjx6yIiOuz98iQJAXIDpSMYhm5lsFiITPDScWzOLLnUR55HL/biaB1zqoODj2so7G2JoTiYiznamF9h9GuFC2TablbINq80U2NcxxQJBAMhw06Ha/U7qTjtAmr2qAuWSWvHU4ANu2h0RxYlKTpmWgO0f47jCOQhdC3T/RK7f38c7q8uPyi35eZ7S1e/PznY=",
		Format:      "json",
		ConTimeOut:  2000,
		ReadTimeOut: 2000,
		SignType:    "RSA",
		Charset:     "GBK",
	}
}

func TestDefaultClientWork(t *testing.T) {

	ImgText := "{'articles':[{'actionName':'立即查看','desc':'这是图文内容','imageUrl':'http://pic.alipayobjects.com/e/201311/1PaQ27Go6H_src.jpg','title':'这是标题','url':'https://www.alipay.com/'}],'msgType':'image-text','toUserId':'kquy1L6pBsGPHDYCYeBhLJqKo7ILToCxhO4skMe5bRK-S4-HyunMvYTqiRYEUD+U01'}"
	r := &request.AlipayMobilePublicMessageCustomSendRequest{
		BizContent: ImgText,
	}

	resp, err := alipayClient.Execute(r)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	log.Printf("%+v", resp)

}

func TestGetOpenId(t *testing.T) {
	r := &request.AlipaySystemOauthTokenRequest{
		GrantType: "authorization_code",
		Code:      "6d46274d082f4f96807ca49e3e30cX92",
	}
	resp, err := alipayClient.Execute(r)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	log.Printf("%+v", resp)

}

func TestGetMobileGis(t *testing.T) {
	r := &request.AlipayMobilePublicGisGetRequest{
		BizContent: "{'userId':" + userId + "}",
	}
	resp, err := alipayClient.Execute(r)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	log.Printf("%+v", resp)
}
