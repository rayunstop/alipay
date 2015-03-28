package dispatcher

import (
	"errors"
	"github.com/alipay-sdk/constants"
	"github.com/alipay-sdk/executor"
	"net/url"
)

// Executor 根据params获取对应的执行器
func Executor(params url.Values) (executor.Executor, error) {
	service := params.Get("service")
	if service == "" {
		return nil, errors.New("service param is null")
	}

	content := params.Get("biz_content")
	if content == "" {
		return nil, errors.New("content param is null")
	}

	msgType := params.Get("MsgType")
	if msgType == "" {
		return nil, errors.New("msgType param is null")
	}

	userId := params.Get("FromUserId")
	if userId == "" {
		return nil, errors.New("userId param is null")
	}

	switch msgType {
	// 根据消息类型处理
	case constants.MsgTypeText:
		//TODO
		return nil, nil
	case constants.MsgTypeImage:
		//TODO
		return nil, nil
	case constants.MsgTypeEvent:
		return eventExecutor(service, content)
	}
}

// eventExecutor 事件执行器
func eventExecutor(service, content string) (executor.Executor, error) {

}
