package flag

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/xmdhs/gomclauncher/auth"
)

func (f *Flag) Aonline() {
	if f.Email == "" {
		fmt.Println("请设置邮箱")
		fmt.Println("比如 -email xxx@xxx.xx")
		os.Exit(0)
	}
	if gmlconfig[auth.ApiAddress] == nil {
		gmlconfig[auth.ApiAddress] = make(map[string]Config)
	}
	err := gmlconfig[auth.ApiAddress][f.Email].setonline(f.Email, f.Password)
	if err != nil {
		if errors.Is(err, HaveProfiles) {
			a := auth.Auth{
				AccessToken: gmlconfig[auth.ApiAddress][f.Email].AccessToken,
				ClientToken: gmlconfig[auth.ApiAddress][f.Email].ClientToken,
			}

			atime := time.Now().Unix()
			if atime-gmlconfig[auth.ApiAddress][f.Email].Time > 1200 {
				if err := auth.Validate(a); err != nil {
					err = auth.Refresh(&a)
					if err != nil {
						if errors.Is(err, auth.NotOk) {
							fmt.Println("请尝试重新登录帐号")
							os.Exit(0)
						} else {
							fmt.Println("登录失败，可能是网络问题，可再次尝试")
							panic(err)
						}
					}
					aconfig := gmlconfig[auth.ApiAddress][f.Email]
					aconfig.Name = a.Username
					aconfig.UUID = a.ID
					aconfig.AccessToken = a.AccessToken
					aconfig.Time = time.Now().Unix()
					aconfig.ClientToken = a.ClientToken
					gmlconfig[auth.ApiAddress][f.Email] = aconfig
					saveconfig()
				}
			}
		} else if errors.Is(err, auth.NotOk) {
			fmt.Println("账户名或密码错误")
			os.Exit(0)
		} else {
			panic(err)
		}
	}
	if gmlconfig[auth.ApiAddress][f.Email].Name == "" {
		panic("请创建角色")
	}
	f.AccessToken = gmlconfig[auth.ApiAddress][f.Email].AccessToken
	f.Name = gmlconfig[auth.ApiAddress][f.Email].Name
	f.UUID = gmlconfig[auth.ApiAddress][f.Email].UUID
}

func Listname() {
	fmt.Println("-----------------")
	for k, v := range gmlconfig {
		if k == "https://authserver.mojang.com" {
			fmt.Println("正版登录")
		} else {
			fmt.Println("外置登录，api 地址", k)
		}
		for k, v := range v {
			fmt.Println("邮箱:", k)
			fmt.Println("用户名:", v.Name)
		}
		fmt.Println("-----------------")
	}
}
