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
	"strings"
	"time"
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
}

// 实现接口
func (d *DefaultAlipayClient) Execute(r request.AlipayRequest) (response.AlipayResponse, error) {
	return d.ExecuteWithToken(r, "")
}

// 实现接口
func (d *DefaultAlipayClient) ExecuteWithToken(r request.AlipayRequest, token string) (response.AlipayResponse, error) {

	// convert utf-8 to gbk
	cd, err := iconv.Open("utf-8", "gbk")
	if err != nil {
		return nil, err
	}
	defer cd.Close()

	// 获取必须参数
	rp := make(map[string]string)
	rp[constants.AppId] = d.AppId
	rp[constants.Method] = r.GetApiMethod()
	rp[constants.SignType] = d.SignType
	rp[constants.Timestamp] = time.Now().Format("2006-01-02 15:03:04")
	rp[constants.Version] = r.GetApiVersion()
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
		return nil, err
	}
	msg, err := ioutil.ReadAll(result.Body)
	if err != nil {
		return nil, err
	}
	// 解析resp
	params := make(map[string]interface{})
	json.Unmarshal(msg, &params)
	resp := r.GetResponse()
	d.resultMapping(resp, params)

	// 不成功
	if !resp.IsSuccess() {
		//TODO
		log.Println("todo to show all error message")
	}
	return resp, nil
}

// resultMapping 将结果映射到response
func (d *DefaultAlipayClient) resultMapping(r response.AlipayResponse, params map[string]interface{}) {

	log.Printf("alipay return : %+v", params)
	// 拿到除sign的另一个key的值
	for k, v := range params {
		if k != "sign" {
			if value, ok := v.(map[string]interface{}); ok {
				conver.Convertor.Conver(r, value)
			}
		}
	}

}
