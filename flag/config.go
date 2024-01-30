package flag

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/xmdhs/gomclauncher/auth"
	"github.com/xmdhs/gomclauncher/lang"
)

func saveconfig(gmlconfig Gmlconfig) {
	b, err := ioutil.ReadFile("gml.json")
	if err != nil && !os.IsNotExist(err) {
		panic(err)
	} else {
		ff, err := os.Create("gml.json.bak")
		aerr(err)
		defer ff.Close()
		_, err = ff.Write(b)
		aerr(err)
	}
	f, err := os.Create("gml.json")
	aerr(err)
	defer f.Close()
	b, err = json.Marshal(gmlconfig)
	aerr(err)
	_, err = f.Write(b)
	aerr(err)
}

//lint:ignore ST1012 导出字段
var HaveProfiles = errors.New("have")

func (c Config) setonline(gmlconfig *Gmlconfig, f *Flag) error {
	if _, ok := (*gmlconfig)[f.ApiAddress][f.Email]; ok && f.Password == "" {
		return HaveProfiles
	}
	if c.ClientToken == "" {
		c.ClientToken = UUIDgen(f.Email)
	}
	a, err := auth.Authenticate(f.ApiAddress, f.Name, f.Email, f.Password, c.ClientToken)
	if err != nil {
		if errors.Is(err, auth.ErrNotSelctProFile) {
			fmt.Println(lang.Lang("ErrNotSelctProFile"))
			list := auth.ListAvailableProfileName(a)
			for _, p := range list {
				fmt.Println(p)
			}
			os.Exit(0)
		} else if errors.Is(err, auth.ErrProFileNoExist) {
			fmt.Println(lang.Lang("ErrProFileNoExist"))
			os.Exit(0)
		}
		return fmt.Errorf("setonline: %w", err)
	}
	var aconfig Config
	aconfig.ClientToken = c.ClientToken
	aconfig.Name = a.Username
	aconfig.UUID = a.ID
	aconfig.AccessToken = a.AccessToken
	(*gmlconfig)[f.ApiAddress][f.Email] = aconfig
	saveconfig(*gmlconfig)
	return nil
}

type Gmlconfig map[string]map[string]Config

type Config struct {
	Name        string
	UUID        string
	ClientToken string
	AccessToken string
	Time        int64
	ExtData     json.RawMessage
}
