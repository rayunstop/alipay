package alipay

import (
	"encoding/json"
	"github.com/huandu/xstrings"
	"github.com/z-ray/alipay/api/constants"
	"github.com/z-ray/alipay/api/conver"
	"github.com/z-ray/alipay/api/logger"
	"github.com/z-ray/alipay/api/request"
	"github.com/z-ray/alipay/api/response"
	"github.com/z-ray/alipay/api/sign"
	"github.com/z-ray/alipay/api/utils"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

// default
const (
	format         = "json"
	signType       = "RSA"
	connectTimeout = 3000
	readTimeout    = 15000
)

// AlipayClient 客户端接口
type AlipayClient interface {
	Execute(r request.AlipayRequest) (response.AlipayResponse, error)
	// 使用token
	ExecuteWithToken(r request.AlipayRequest, token string) (response.AlipayResponse, error)
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
	Charset     string
}

// 实现接口
func (d *DefaultAlipayClient) Execute(r request.AlipayRequest) (response.AlipayResponse, error) {
	return d.ExecuteWithToken(r, "")
}

// 实现接口
func (d *DefaultAlipayClient) ExecuteWithToken(r request.AlipayRequest, token string) (response.AlipayResponse, error) {

	// 请求
	msg, rp, err := d.post(r, token)
	if err != nil {
		return nil, err
	}
	log.Printf("alipay return : %s", string(msg))

	// body
	resp := r.GetResponse()
	resp.SetBody(string(msg))

	// 解析resp
	params := make(map[string]interface{})
	err = json.Unmarshal(msg, &params)
	if err != nil {
		log.Println(err)
	}

	// 获得响应报文
	v := params[xstrings.ToSnakeCase(resp.ToStr())]
	// 可能返回失败报文
	if v == nil {
		v = params["error_response"]
	}
	if sub, ok := v.(map[string]interface{}); ok {
		params = sub
	}
	// 映射
	err = conver.Do(resp, params)
	if err != nil {
		// TODO 系统、业务错误
		log.Println(err)
	}

	// 当发生安全机制接入错误时
	// 详细见https://fuwu.alipay.com/platform/doc.htm#c09
	if !resp.IsSuccess() {
		logger.SecureError(rp, resp)
	}
	return resp, nil
}

func (d *DefaultAlipayClient) post(r request.AlipayRequest, token string) ([]byte, map[string]string, error) {
	// 获取必须参数
	rp := make(map[string]string)
	rp[constants.AppId] = d.AppId
	rp[constants.Method] = r.GetApiMethod()
	rp[constants.SignType] = signType // TODO
	rp[constants.Timestamp] = time.Now().Format("2006-01-02 15:03:04")
	rp[constants.Version] = r.GetApiVersion()
	rp[constants.Charset] = d.Charset
	utils.PutAll(rp, r.GetTextParams())
	// 可选参数
	// rp[constants.Format] = d.Format

	// 请求报文
	content := sign.PrepareContent(rp)
	// 签名
	signed, err := sign.RsaSign(content, d.PrivKey)
	rp[constants.Sign] = signed

	// 编码查询参数
	values := utils.BuildQuery(rp)
	// 请求
	result, err := http.Post(d.ServerURL, "application/x-www-form-urlencoded;charset=utf-8", strings.NewReader(values.Encode()))
	if err != nil {
		log.Println(err)
	}
	msg, err := ioutil.ReadAll(result.Body)
	if err != nil {
		log.Println(err)
	}
	return msg, rp, err
}
