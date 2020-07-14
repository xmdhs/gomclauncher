package flag

import (
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
	err := gmlconfig[f.Email].setonline(f.Email, f.Password)
	if err != nil {
		if err.Error() == "have" {
			a := auth.Auth{
				AccessToken: gmlconfig[f.Email].AccessToken,
				ClientToken: gmlconfig[f.Email].ClientToken,
			}
			if gmlconfig[f.Email].Authlib != "" {
				auth.ApiAddress = gmlconfig[f.Email].Authlib
			}
			atime := time.Now().Unix()
			if atime-gmlconfig[f.Email].Time > 1200 {
				if err := auth.Validate(a); err != nil {
					err = auth.Refresh(&a)
					if err != nil {
						if err.Error() == "not ok" {
							fmt.Println("请尝试重新登录帐号")
							os.Exit(0)
						} else {
							fmt.Println("登录失败，可能是网络问题，可再次尝试")
							panic(err)
						}
					}
					aconfig := gmlconfig[f.Email]
					aconfig.Name = a.Username
					aconfig.UUID = a.ID
					aconfig.AccessToken = a.AccessToken
					aconfig.Time = time.Now().Unix()
					aconfig.ClientToken = a.ClientToken
					aconfig.Authlib = auth.ApiAddress
					gmlconfig[f.Email] = aconfig
					saveconfig()
				}
			}
		} else if err.Error() == "not ok" {
			fmt.Println("账户名或密码错误")
			os.Exit(0)
		} else {
			panic(err)
		}
	}
	if gmlconfig[f.Email].Name == "" {
		panic("请创建或者选择角色")
	}
	f.Userproperties = gmlconfig[f.Email].Userproperties
	f.AccessToken = gmlconfig[f.Email].AccessToken
	f.Name = gmlconfig[f.Email].Name
	f.UUID = gmlconfig[f.Email].UUID
}
