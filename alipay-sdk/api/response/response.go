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
	Code    string `align:"code"`
	Msg     string `align:"msg"`
	SubCode string `align:"sub_code"`
	SubMsg  string `align:"sub_msg"`
}

func (r *AlipayMobilePublicMessageCustomSendResponse) IsSuccess() bool {
	// sub_code如果为空，表明执行成功
	return strings.EqualFold("", r.SubCode)
}
