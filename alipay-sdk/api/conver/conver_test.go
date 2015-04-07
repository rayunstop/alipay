package conver

import (
	"encoding/json"
	"log"
	"testing"
)

type People struct {
	Age     string `align:"age"`
	Name    string `align:"name"`
	Message Msg    `align:"msg"`
}

type Msg struct {
	Time    string `align:"time"`
	Address Addr   `align:"address"`
}

type Addr struct {
	Home string `align:"home"`
}

var rsp = `{"response":{"age":1800,"name":"test","msg":{"time":"20150331","address":{"home":"hello world"}}},"sign":"adsuhdawkjdiahandawdh"}`

func TestConver(t *testing.T) {

	params := make(map[string]interface{})
	json.Unmarshal([]byte(rsp), &params)

	o := &People{}
	err := Do(o, params)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	log.Printf("%+v", o)
}
