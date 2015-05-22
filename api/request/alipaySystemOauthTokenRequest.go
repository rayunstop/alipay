package request

import (
	"github.com/rui2014/alipay/api/response"
)

const (
	Auth_Code = "authorization_code"
)

// AlipaySystemOauthTokenRequest
// API:  alipay.system.oauth.token request
type AlipaySystemOauthTokenRequest struct {
	Code         string
	GrantType    string
	RefreshToken string
}

func (r *AlipaySystemOauthTokenRequest) GetApiMethod() string {
	return "alipay.system.oauth.token"
}

func (r *AlipaySystemOauthTokenRequest) GetTextParams() map[string]string {
	params := make(map[string]string)
	params["code"] = r.Code
	params["grant_type"] = r.GrantType
	params["refresh_token"] = r.RefreshToken
	//TODO 提供一个用户设置参数的接口
	//utils.putAll(params,userParams)
	return params
}

func (r *AlipaySystemOauthTokenRequest) GetResponse() response.AlipayResponse {
	resp := new(response.AlipaySystemOauthTokenResponse)
	resp.Name = "AlipaySystemOauthTokenResponse"
	return resp
}

func (r *AlipaySystemOauthTokenRequest) GetApiVersion() string {
	return "1.0"
}
