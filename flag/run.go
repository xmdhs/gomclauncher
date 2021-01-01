package flag

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/xmdhs/gomclauncher/lang"
	"github.com/xmdhs/gomclauncher/launcher"
)

func (f Flag) Arun() {
	f.Version = f.Run
	if f.Independent {
		f.Gamedir = f.Minecraftpath + "/versions/" + f.Version
	} else {
		f.Gamedir = f.Minecraftpath
	}
	b, err := ioutil.ReadFile(f.Minecraftpath + "/versions/" + f.Version + "/" + f.Version + ".json")
	if err != nil {
		fmt.Println(lang.Lang("nofindthisversion"))
		panic(err)
	}
	if f.Outmsg {
		t := test{}
		err := json.Unmarshal(b, &t)
		if err != nil {
			panic(err)
		}
		if t.ID != f.Version {
			b = bytes.ReplaceAll(b, []byte(t.ID), []byte(f.Version))
			err := ioutil.WriteFile(f.Minecraftpath+"/versions/"+f.Version+"/"+f.Version+".json", b, 0777)
			if err != nil {
				panic(err)
			}
		}
		if t.InheritsFrom != "" {
			f.Download = t.InheritsFrom
			f.D()
		} else {
			f.Download = f.Version
			f.D()
		}
	}
	f.Jsonbyte = b
	err = f.Run115()
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println(lang.Lang("flag.os.ErrNotExist"))
			os.Exit(0)
		}
		if errors.Is(err, launcher.JsonErr) {
			fmt.Printf(lang.Lang("launcher.JsonErr"), launcher.Minecraft)
			os.Exit(0)
		}
		if errors.Is(err, launcher.JsonNorTrue) {
			fmt.Println(lang.Lang("launcher.JsonNorTrue"))
			os.Exit(0)
		}
		panic(err)
	}
}

type test struct {
	ID           string `json:"id"`
	InheritsFrom string `json:"inheritsFrom"`
}
