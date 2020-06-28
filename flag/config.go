package flag

import (
	"encoding/json"
	"errors"
	"gomclauncher/auth"
	"io/ioutil"
	"log"
	"os"
)

var aconfig Config

func init() {
	aconfig.EmailAccessToken = make(map[string]string)
	_, err := os.Stat("gml.json")
	if err != nil {
		return
	}
	f, err := os.Open("gml.json")
	defer f.Close()
	if err != nil {
		panic(err)
	}
	b, err := ioutil.ReadAll(f)
	err = json.Unmarshal(b, &aconfig)
	if err != nil {
		log.Fatalln("json 损坏，可尝试删除 gml.json")
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
	b, err = json.Marshal(aconfig)
	aerr(err)
	_, err = f.Write(b)
	aerr(err)
}

func (c Config) setonline(email, pass string) error {
	if _, ok := aconfig.EmailAccessToken[email]; ok {
		return errors.New("have")
	}
	defer saveconfig()
	if c.ClientToken == "" {
		c.ClientToken = UUIDgen(email)
	}
	a, err := auth.Authenticate(email, pass, c.ClientToken)
	if err != nil {
		return err
	}
	aconfig.ClientToken = c.ClientToken
	aconfig.Name = a.Username
	aconfig.UUID = a.ID
	aconfig.AccessToken = a.AccessToken
	aconfig.EmailAccessToken[email] = a.AccessToken
	aconfig.Userproperties = a.Userproperties
	return nil
}

type Config struct {
	Name             string
	UUID             string
	ClientToken      string
	Userproperties   string
	AccessToken      string
	EmailAccessToken map[string]string
}
