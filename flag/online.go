package flag

import (
	"gomclauncher/auth"
	"log"
	"time"
)

func (f *Flag) Aonline() {
	if f.Email == "" {
		log.Println("请设置邮箱")
		log.Fatal("比如 -email xxx@xxx.xx")
	}
	err := gmlconfig[f.Email].setonline(f.Email, f.Passworld)
	if err != nil {
		if err.Error() == "have" {
			a := auth.Auth{
				AccessToken: gmlconfig[f.Email].AccessToken,
				ClientToken: gmlconfig[f.Email].ClientToken,
			}
			atime := time.Now().Unix()
			if atime-gmlconfig[f.Email].Time > 1200 {
				if err := auth.Validate(a); err != nil {
					err = auth.Refresh(&a)
					if err != nil {
						if err.Error() == "not ok" {
							log.Fatalln("请尝试重新登录")
						} else if err.Error() == "proxy err" {
							log.Fatalln("设置的代理有误")
						} else {
							log.Fatalln("可能是网络问题，可再次尝试")
						}
					}
					var aconfig Config
					aconfig.Name = a.Username
					aconfig.UUID = a.ID
					aconfig.AccessToken = a.AccessToken
					aconfig.Time = time.Now().Unix()
					gmlconfig[f.Email] = aconfig
					saveconfig()
				}
			}
		} else if err.Error() == "not ok" {
			log.Fatalln("账户名或密码错误")
		} else {
			log.Fatalln(err)
		}
	}
	f.Userproperties = gmlconfig[f.Email].Userproperties
	f.AccessToken = gmlconfig[f.Email].AccessToken
	f.Name = gmlconfig[f.Email].Name
	f.UUID = gmlconfig[f.Email].UUID
}
