package flag

import (
	"fmt"
	"io/ioutil"
	"log"
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
		log.Fatalln(err)
	}
	f.Jsonbyte = b
	err = f.Run115()
	if err != nil {
		if err.Error() == "json not exist" {
			log.Fatalln("请先安装对应的原版")
		}
		if err.Error() == "json err" {
			log.Fatalln("json 错误，可尝试到 .minecraft/versions 中删除对应的 json 文件")
		}
	}
}
