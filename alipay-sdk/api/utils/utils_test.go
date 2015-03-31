package utils

import (
	"log"
	"testing"
)

func TestBuildQuery(t *testing.T) {

	url := "https://localhost:8080/"
	params := make(map[string]string)
	params["content"] = "<xml></xml>"
	params["biz"] = "biz"
	params["sign"] = "sign"
	params["sign_type"] = "rsa"

	content := BuildQuery(url, params)
	if content == "" {
		t.Error("content must not be empty")
		t.FailNow()
	}
	log.Printf("%s", content)
}
