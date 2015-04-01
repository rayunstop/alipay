package response

import (
	"strings"
)

// AlipayResponse response接口
type AlipayResponse interface {

	// 判断是否成功
	IsSuccess() bool
}

// AlipayMobilePublicMessageCustomSendResponse
// 与AlipayMobilePublicMessageCustomSendRequest关联
type AlipayMobilePublicMessageCustomSendResponse struct {
	Code    string `json:"code"`
	Msg     string `json:"msg"`
	SubCode string `json:"sub_code"`
	SubMsg  string `json:"sub_msg"`
}

func (r *AlipayMobilePublicMessageCustomSendResponse) IsSuccess() bool {
	// sub_code如果为空，表明执行成功
	return strings.EqualFold("", r.SubCode)
}
