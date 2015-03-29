package dispatcher

import (
	"encoding/xml"
	"errors"
	"github.com/alipay-sdk/constants"
	"github.com/alipay-sdk/executor"
	"github.com/alipay-sdk/model"
	"net/url"
)

// Executor 根据params获取对应的执行器
func Executor(params url.Values) (*executor.Executor, error) {
	service := params.Get("service")
	if service == "" {
		return nil, errors.New("service param is null")
	}

	content := params.Get("biz_content")
	if content == "" {
		return nil, errors.New("content param is null")
	}

	// 解析xml
	bizContent := new(model.BizContent)
	err := xml.Unmarshal([]byte(content), bizContent)

	msgType := bizContent.MsgType
	if msgType == "" {
		return nil, errors.New("msgType param is null")
	}

	userId := bizContent.FromUserId
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
		return eventExecutor(service, bizContent)
	}
}

// eventExecutor 事件执行器
func eventExecutor(service string, content BizContent) (*executor.Executor, error) {

	eventType := content.EventType
	// 根据service的不同细分
	if constants.ServerTypeCheck == service && constants.EventTypeVerifyGw == eventType {
		return new(executor.AlipayVerifyExecutor)
	}
	// 消息类事件
	if constants.ServerTypeMsgNotify == service {
		switch eventType {

		case constants.EventTypeFollow:
			//TODO
			return nil, nil
		case constants.EventTypeUnFollow:
			//TODO
			return nil, nil
		case constants.EventTypeClick:
			//TODO
			return nil, nil
		case constants.EventTypeEnter:
			//TODO
			return nil, nil
		}
	}
	// 暂不支持其他类型
	return nil, errors.New(eventType + " event does not support yet")
}
