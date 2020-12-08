package flag

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/xmdhs/gomclauncher/auth"
	msauth "github.com/xmdhs/msauth/auth"
)

func (f *Flag) MsLogin() {
	if f.Email == "" {
		fmt.Println("虽然实际上不需要你输入邮箱，但是需要用邮箱来标记账号，以保存 accesstoken，当然，你也可以完全输入一个不是邮箱的字符串")
		os.Exit(0)
	}
	if f.Gmlconfig["ms"] == nil {
		f.Gmlconfig["ms"] = make(map[string]Config)
	}
	var p *auth.Profile
	c, ok := f.Gmlconfig["ms"][f.Email]
	if ok {
		var err error
		p, err = auth.GetProfile(c.AccessToken)
		if err != nil {
			p, err = auth.MsLogin()
			if err != nil {
				msLogincheakErr(err)
			}
		}
	} else {
		var err error
		p, err = auth.MsLogin()
		if err != nil {
			msLogincheakErr(err)
		}
	}
	aconfig := f.Gmlconfig["ms"][f.Email]
	aconfig.Name = p.Name
	aconfig.UUID = p.ID
	aconfig.AccessToken = p.AccessToken
	aconfig.Time = time.Now().Unix()
	aconfig.ClientToken = ""
	f.Gmlconfig["ms"][f.Email] = aconfig
	saveconfig(f.Gmlconfig)

	f.Name = p.Name
	f.UUID = p.ID
	f.AccessToken = p.AccessToken

	f.Email = ""
}

func msLogincheakErr(err error) {
	switch {
	case errors.Is(err, msauth.ErrHostname):
		fmt.Println("忘记密码什么的就别在这操作啦。")
	case errors.Is(err, auth.ErrCode):
		fmt.Println("尝试重新登录微软的账号")
	case errors.Is(err, auth.ErrProfile):
		fmt.Println("你好像还没买游戏，或者还没迁移账号呢")
	default:
		panic(err)
	}
	os.Exit(0)
}
