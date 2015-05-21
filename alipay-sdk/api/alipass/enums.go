package alipass

// alipass用户识别方式枚举类型
const (
	TRADE  = "1" // 支付宝交易
	USERID = "2" // 支付宝用户ID
	MOBILE = "3" // 支付宝用户绑定手机号
	OPENID = "4" // 支付宝公众号开放ID
)

// alipass 状态枚举
const (
	PASS_STATUS_CLOSED = "CLOSED"
	PASS_STATUS_USED   = "USED"
)
