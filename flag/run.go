package flag

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

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
		fmt.Println("没有这个版本或者其他问题")
		fmt.Println(err)
		os.Exit(0)
	}
	if f.Outmsg {
		t := test{}
		json.Unmarshal(b, &t)
		if t.InheritsFrom != "" {
			f.Download = t.InheritsFrom
			f.D()
		} else {
			f.Download = t.ID
			f.D()
		}
	}
	f.Jsonbyte = b
	err = f.Run115()
	if err != nil {
		if err.Error() == "json not exist" {
			fmt.Println("请先安装对应的原版")
			os.Exit(0)
		}
		if err.Error() == "json err" {
			fmt.Println("json 错误，可尝试到 " + launcher.Minecraft + "/versions 中删除对应的 json 文件")
			os.Exit(0)
		}
	}
}

type test struct {
	ID           string `json:"id"`
	InheritsFrom string `json:"inheritsFrom"`
}
