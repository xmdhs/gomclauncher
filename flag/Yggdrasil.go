package flag

import (
	"context"
	"errors"
	"fmt"

	"github.com/xmdhs/gomclauncher/auth"
	"github.com/xmdhs/gomclauncher/download"
)

func (f Flag) Authlib() {
	err := download.Downauthlib(context.Background())
	if err != nil {
		fmt.Println("authlib-injector 下载失败")
		panic(err)
	}
	api, err := auth.Getauthlibapi(f.ApiAddress)
	if err != nil {
		if errors.Is(err, auth.JsonNotTrue) {
			panic("外置登录地址错误")
		} else {
			fmt.Println("或许是网络问题")
			panic(err)
		}
	}
	f.ApiAddress = api
}
