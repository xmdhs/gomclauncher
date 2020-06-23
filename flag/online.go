package flag

import (
	"gomclauncher/auth"
	"log"
)

func (f Flag) Aonline() {
	if f.Email == "" {
		log.Println("请设置邮箱")
		log.Fatal("比如 -email xxx@xxx.xx")
	}
	err := aconfig.setonline(f.Email, f.Passworld)
	if err != nil {
		if err.Error() == "have" {
			a := auth.Auth{
				AccessToken: aconfig.EmailAccessToken[f.Email],
				ClientToken: aconfig.ClientToken,
			}
			err := auth.Refresh(&a)
			if err != nil {
				if err.Error() == "not ok" {
					log.Fatalln("请尝试重新登录")
				} else if err.Error() == "proxy err" {
					log.Fatalln("设置的代理有误")
				} else {
					log.Fatalln("可能是网络问题，可再次尝试")
				}
			}
			aconfig.Name = a.Username
			aconfig.EmailAccessToken[f.Email] = a.AccessToken
		} else if err.Error() == "not ok" {
			log.Fatalln("账户名或密码错误")
		} else {
			log.Fatalln(err)
		}
	}
	f.Userproperties = aconfig.Userproperties
	f.AccessToken = aconfig.AccessToken
	f.Name = aconfig.Name
	f.UUID = aconfig.UUID
}
