package executor

import (
	"fmt"
	"github.com/alipay-sdk/constants"
)

// Executor 事件执行器接口
type Executor interface {

	// 执行方法
	execute() (string, error)
}

type AlipayVerifyExecutor struct{}

// execute 网关验证
func (e AlipayVerifyExecutor) execute() (string, error) {
	bulider := "<success>true</success><biz_content>%s</biz_content>"
	return fmt.Sprintf(bulider, constants.AliPubKey)
}
