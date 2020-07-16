package flag

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/xmdhs/gomclauncher/auth"
)

var gmlconfig Gmlconfig

func init() {
	gmlconfig = make(Gmlconfig)
	b, err := ioutil.ReadFile("gml.json")
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		panic(err)
	}
	err = json.Unmarshal(b, &gmlconfig)
	if err != nil {
		fmt.Println("json 损坏，可尝试删除 gml.json")
		panic(err)
	}
}

func saveconfig() {
	b, err := ioutil.ReadFile("gml.json")
	if err != nil && !os.IsNotExist(err) {
		panic(err)
	} else {
		ff, err := os.Create("gml.json.bak")
		defer ff.Close()
		aerr(err)
		_, err = ff.Write(b)
		aerr(err)
	}
	f, err := os.Create("gml.json")
	defer f.Close()
	aerr(err)
	b, err = json.Marshal(gmlconfig)
	aerr(err)
	_, err = f.Write(b)
	aerr(err)
}

func (c Config) setonline(email, pass string) error {
	if _, ok := gmlconfig[auth.ApiAddress][email]; ok && pass == "" {
		return errors.New("have")
	}
	if c.ClientToken == "" {
		c.ClientToken = UUIDgen(email)
	}
	a, err := auth.Authenticate(email, pass, c.ClientToken)
	if err != nil {
		return err
	}
	var aconfig Config
	aconfig.ClientToken = c.ClientToken
	aconfig.Name = a.Username
	aconfig.UUID = a.ID
	aconfig.AccessToken = a.AccessToken
	gmlconfig[auth.ApiAddress][email] = aconfig
	saveconfig()
	return nil
}

type Gmlconfig map[string]map[string]Config

type Config struct {
	Name        string
	UUID        string
	ClientToken string
	AccessToken string
	Time        int64
}
