package flag

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/xmdhs/gomclauncher/auth"
	"github.com/xmdhs/gomclauncher/lang"
)

func (f *Flag) MsLogin() {
	if f.Email == "" {
		fmt.Println(lang.Lang("msemailnil"))
		os.Exit(0)
	}
	if f.Gmlconfig["ms"] == nil {
		f.Gmlconfig["ms"] = make(map[string]Config)
	}
	token := &auth.MsToken{}
	c, ok := f.Gmlconfig["ms"][f.Email]
	if ok && c.ExtData != nil {
		err := json.Unmarshal(c.ExtData, token)
		if err != nil {
			token = nil
		}
	}
	p, err := auth.GetProfile(c.AccessToken)
	if err != nil {
		p, err = auth.MsLoginRefresh(token)
		if err != nil {
			msLogincheakErr(err)
		}
	} else {
		p.MsToken = *token
	}
	aconfig := f.Gmlconfig["ms"][f.Email]
	aconfig.Name = p.Name
	aconfig.UUID = p.ID
	aconfig.AccessToken = p.AccessToken
	aconfig.Time = time.Now().Unix()
	aconfig.ClientToken = ""
	b, err := json.Marshal(p.MsToken)
	if err != nil {
		panic(err)
	}
	aconfig.ExtData = b
	f.Gmlconfig["ms"][f.Email] = aconfig
	saveconfig(f.Gmlconfig)

	f.Name = p.Name
	f.UUID = p.ID
	f.AccessToken = p.AccessToken

	f.Email = ""
}

func msLogincheakErr(err error) {
	switch {
	case errors.Is(err, auth.ErrChild):
		fmt.Println(lang.Lang("ErrChild"))
	case errors.Is(err, auth.ErrCode):
		fmt.Println(lang.Lang("auth.ErrCode"))
	case errors.Is(err, auth.ErrProfile):
		fmt.Println(lang.Lang("auth.ErrProfile"))
	default:
		panic(err)
	}
	os.Exit(0)
}
