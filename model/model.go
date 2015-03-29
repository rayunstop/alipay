package model

import (
	"encoding/xml"
)

// BizContent 业务参数结构体
type BizContent struct {
	XMLName     xml.Name `xml:"XML"`
	AppId       string
	FromUserId  string
	CreateTime  int64
	MsgType     string
	EventType   string
	ActionParam string
	AgreementId string
	AccountNo   string
}
