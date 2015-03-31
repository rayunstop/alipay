package api

import (
	"encoding/json"
	"github.com/alipay/alipay-sdk/api/constants"
	"github.com/alipay/alipay-sdk/api/sign"
	"github.com/alipay/alipay-sdk/api/utils"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// AlipayRequest request接口
type AlipayRequest interface {
	// appId
	GetAppId() string

	// 字符集
	GetCharSet() string

	// 方法名称
	GetApiMethod() string

	// 签名类型
	GetSignType() string

	// 应用参数
	// 包括biz_content、自定义的参数
	GetTextParams() map[string]string

	// 每一个request必须绑定一个response对象
	GetResponse() *AlipayResponse
}

// AlipayResponse response接口
type AlipayResponse interface {

	// 判断是否成功
	IsSuccess() bool
}

// AlipayClient 客户端接口
type AlipayClient interface {
	execute(request *AlipayRequest) (*AlipayResponse, error)
	// 使用token
	executeWithToken(request *AlipayRequest, token string) (*AlipayResponse, error)
}

// DefaultAlipayClient 默认的client
type DefaultAlipayClient struct {
	AppId       string
	ServerURL   string
	PrivKey     string
	Format      string
	ConTimeOut  int32
	ReadTimeOut int32
	SignType    string
}

// 实现接口
func (d *DefaultAlipayClient) execute(request *AlipayRequest) (*AlipayResponse, error) {
	return d.executeWithToken(request, nil)
}

// 实现接口
func (d *DefaultAlipayClient) executeWithToken(request *AlipayRequest, token string) (*AlipayResponse, error) {
	// 获取必须参数
	must := make(map[string]string)
	must[constants.AppId] = request.GetAppId()
	must[constants.Method] = request.GetApiMethod()
	must[constants.SignType] = d.SignType

	// 可选参数
	opt := make(map[string]string)
	opt[constants.Format] = d.Format

	// 请求报文
	content := sign.PrepareContent(must)
	// 签名
	signed, err := sign.RsaSign(content, d.PrivKey)
	must[constants.Sign] = signed

	// 生成查询URL
	q := utils.BuildQuery(d.ServerURL, must)
	// 请求
	result, err := http.Post(q, "application/x-www-form-urlencoded", nil)
	if err != nil {
		return nil, err
	}
	msg, err := ioutil.ReadAll(result.Body)
	if err != nil {
		return nil, err
	}
	// 解析resp
	resp := request.GetResponse()
	json.Unmarshal(msg, resp)

	// 不成功
	if !resp.IsSuccess() {
		//TODO
		log.Println("todo to show all error message")
	}
	return resp, nil
}
