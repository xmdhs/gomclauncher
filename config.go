package main

import (
	"encoding/json"
	"errors"
	"gomclauncher/auth"
	"io/ioutil"
	"os"

	"github.com/google/uuid"
)

type config struct {
	Name             string
	UUID             string
	ClientToken      string
	Userproperties   string
	AccessToken      string
	Nameuuid         map[string]string
	EmailAccessToken map[string]string
}

func (c config) setonline(email, pass string) error {
	if _, ok := aconfig.EmailAccessToken[email]; ok {
		return errors.New("have")
	}
	defer saveconfig()
	if c.ClientToken == "" {
		c.ClientToken = uuid.New().String()
	}
	a, err := auth.Authenticate(email, pass, c.ClientToken)
	if err != nil {
		return err
	}
	aconfig.Name = a.Username
	aconfig.UUID = a.ID
	aconfig.Nameuuid[a.Username] = a.ID
	aconfig.AccessToken = a.AccessToken
	aconfig.EmailAccessToken[email] = a.AccessToken
	aconfig.Userproperties = a.Userproperties
	return nil
}

var aconfig config

func init() {
	_, err := os.Stat("gml.json")
	if err != nil {
		if os.IsNotExist(err) {
			f, err := os.Create("gml.json")
			defer f.Close()
			if err != nil {
				panic(err)
			}
			aconfig.Nameuuid = make(map[string]string)
			aconfig.EmailAccessToken = make(map[string]string)
		} else {
			panic(err)
		}
	} else {
		f, err := os.Open("gml.json")
		defer f.Close()
		if err != nil {
			panic(err)
		}
		b, err := ioutil.ReadAll(f)
		json.Unmarshal(b, &aconfig)
	}

}

func saveconfig() {
	b, err := ioutil.ReadFile("gml.json")
	aerr(err)
	ff, err := os.Create("gml.json.bak")
	defer ff.Close()
	aerr(err)
	_, err = ff.Write(b)
	aerr(err)
	f, err := os.Create("gml.json")
	defer f.Close()
	aerr(err)
	b, err = json.Marshal(aconfig)
	aerr(err)
	_, err = f.Write(b)
	aerr(err)
}

func aerr(err error) {
	if err != nil {
		panic(err)
	}
}
