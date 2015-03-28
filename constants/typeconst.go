package constants

const (

	// 消息类型
	MsgTypeText  = "text"
	MsgTypeImage = "image"
	MsgTypeEvent = "event"

	// 事件类型
	EventTypeVerifyGw = "verifygw" // 验证网关
	EventTypeFollow   = "follow"   // 关注
	EventTypeUnFollow = "unfollow" // 取消关注
	EventTypeClick    = "click"    // 点击
	EventTypeEnter    = "enter"    // 进入

	// 服务类型
	ServerTypeCheck     = "alipay.service.check"
	ServerTypeMsgNotify = "alipay.mobile.public.message.notify"
)
