package api

import (
	"encoding/json"
	"github.com/alipay/alipay-sdk/api/constants"
	"github.com/alipay/alipay-sdk/api/conver"
	"github.com/alipay/alipay-sdk/api/request"
	"github.com/alipay/alipay-sdk/api/response"
	"github.com/alipay/alipay-sdk/api/sign"
	"github.com/alipay/alipay-sdk/api/utils"
	"github.com/qiniu/iconv"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// AlipayClient 客户端接口
type AlipayClient interface {
	Execute(r *request.AlipayRequest) (*response.AlipayResponse, error)
	// 使用token
	ExecuteWithToken(r *request.AlipayRequest, token string) (*response.AlipayResponse, error)
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
func (d *DefaultAlipayClient) Execute(r *request.AlipayRequest) (*response.AlipayResponse, error) {
	return d.executeWithToken(r, nil)
}

// 实现接口
func (d *DefaultAlipayClient) ExecuteWithToken(r *request.AlipayRequest, token string) (*response.AlipayResponse, error) {

	// convert utf-8 to gbk
	cd, err := iconv.Open("utf-8", "gbk")
	if err != nil {
		return nil, err
	}
	defer cd.Close()

	// 获取必须参数
	must := make(map[string]string)
	must[constants.AppId] = r.GetAppId()
	must[constants.Method] = r.GetApiMethod()
	must[constants.SignType] = d.SignType

	// 可选参数
	opt := make(map[string]string)
	opt[constants.Format] = d.Format

	// 请求报文
	content := sign.PrepareContent(must)
	// 签名
	signed, err := sign.RsaSign(content, d.PrivKey)
	must[constants.Sign] = signed

	// 编码查询参数
	data := utils.BuildQuery(must)
	// 请求
	result, err := http.PostForm(d.ServerURL, data)
	if err != nil {
		return nil, err
	}
	msg, err := ioutil.ReadAll(result.Body)
	if err != nil {
		return nil, err
	}
	// 解析resp
	params := make(map[string]interface{})
	json.Unmarshal(msg, &params)
	d.resultMapping(r.GetResponse(), params)

	// 不成功
	if !resp.IsSuccess() {
		//TODO
		log.Println("todo to show all error message")
	}
	return resp, nil
}

// resultMapping 将结果映射到response
func (d *DefaultAlipayClient) resultMapping(r *response.AlipayResponse, params map[string]interface{}) {

	// 拿到除sign的另一个key的值
	for k, v := range params {
		if k != "sign" {
			if value, ok := v.(map[string]interface{}); ok {
				conver.Convertor.Conver(r, value)
			}
		}
	}

}
