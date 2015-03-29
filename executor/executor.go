package executor

import (
	"errors"
	"fmt"
	"github.com/alipay-sdk/constants"
	"github.com/alipay-sdk/model"
)

// Executor 事件执行器接口
type Executor interface {

	// 执行方法
	execute() (string, error)
}

// AlipayVerifyExecutor 网关验证执行器
type AlipayVerifyExecutor struct{}

// AlipayChatTextExecutor 图文消息执行器
type AlipayChatTextExecutor struct {
	BizContent *model.BizContent
}

// execute 网关验证
func (e AlipayVerifyExecutor) execute() (string, error) {
	bulider := "<success>true</success><biz_content>%s</biz_content>"
	return fmt.Sprintf(bulider, constants.AliPubKey)
}

// execute 图文消息
func (e AlipayChatTextExecutor) execute() (string, error) {
	if e.BizContent == nil {
		return nil, errors.New("bizContent is nil")
	}
	userId := e.BizContent.FromUserId
	//TODO

	return nil, nil
}
