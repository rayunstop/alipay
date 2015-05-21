package request

import (
	"github.com/alipay/alipay-sdk/api/response"
)

// AlipayRequest request接口
type AlipayRequest interface {

	// 方法名称
	GetApiMethod() string

	// 版本号
	GetApiVersion() string

	// 应用参数
	// 包括biz_content、自定义的参数
	GetTextParams() map[string]string

	// 每一个request必须绑定一个response对象
	GetResponse() response.AlipayResponse
}
