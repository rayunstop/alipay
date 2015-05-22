package conver

import (
	"encoding/json"
	"log"
	"testing"
)

type Anyom1 struct {
	Single string `align:"single"`
}

type Anyom struct {
	Sex string `align:"sex"`
	Anyom1
}

type People struct {
	Age     string `align:"age"`
	Name    string `align:"name"`
	Message Msg    `align:"msg"`
	Anyom
}

type Msg struct {
	Time    string `align:"time"`
	Address Addr   `align:"address"`
}

type Addr struct {
	Home string `align:"home"`
}

var rsp = `{"age":1800,"name":"test","sex":"nan","single":"yes","msg":{"time":"20150331","address":{"home":"hello world"}},"sign":"adsuhdawkjdiahandawdh"}`

func TestConver(t *testing.T) {

	params := make(map[string]interface{})
	err := json.Unmarshal([]byte(rsp), &params)
	if err != nil {
		log.Println(err)
	}

	o := &People{}
	err = Do(o, params)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	log.Printf("%+v", o)
}
