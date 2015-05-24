package logger

import (
	// "bytes"
	"fmt"
	"github.com/z-ray/alipay/api/response"
	// "os"
	"runtime"
	"time"
)

// SecureError 接入支付宝时安全机制错误
func SecureError(params map[string]string, resp response.AlipayResponse) {

	now := time.Now()
	// osName := os.
	version := runtime.Version()
	fmt.Printf("[Alipay] %v | %3s:%s:%s | %s | %+v",
		now.Format("2006/01/02-15:04:05"),
		resp.GetCode(), resp.GetSubCode(), resp.GetMsg(),
		version,
		params,
	)

}
