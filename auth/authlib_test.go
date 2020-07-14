package auth

import (
	"fmt"
	"testing"
)

func TestGetauthlibapi(t *testing.T) {
	api, err := Getauthlibapi(`littleskin.cn`)
	if err != nil {
		t.Fatal(err)
	}
	if api != `https://littleskin.cn/api/yggdrasil` {
		t.Fail()
	}
	fmt.Println(api)
	api, err = Getauthlibapi(`baidu.com`)
	if err == nil {
		t.Fatal(err)
	}
	fmt.Println(api)
}
