package sign

import (
	"fmt"
	"testing"
)

func TestPrepareContent(t *testing.T) {
	params := make(map[string]string)
	params["content"] = "<xml></xml>"
	params["biz"] = "biz"
	params["sign"] = "sign"
	params["sign_type"] = "rsa"

	content := PrepareContent(params)
	fmt.Println(content)
}
