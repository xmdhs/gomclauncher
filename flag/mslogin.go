package flag

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/xmdhs/gomclauncher/auth"
	"github.com/xmdhs/gomclauncher/lang"
	msauth "github.com/xmdhs/msauth/auth"
)

func (f *Flag) MsLogin() {
	if f.Email == "" {
		fmt.Println(lang.Lang("msemailnil"))
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
		fmt.Println(lang.Lang("msauth.ErrHostname"))
	case errors.Is(err, auth.ErrCode):
		fmt.Println(lang.Lang("auth.ErrCode"))
	case errors.Is(err, auth.ErrProfile):
		fmt.Println(lang.Lang("auth.ErrProfile"))
	case errors.Is(err, msauth.ErrNotInstallChrome):
		fmt.Println(lang.Lang("msauth.ErrNotInstallChrome"))
	default:
		panic(err)
	}
	os.Exit(0)
}
