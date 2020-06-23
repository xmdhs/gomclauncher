package flag

import (
	"io/ioutil"
	"log"
)

func (f Flag) arun() {
	f.Version = f.Run
	b, err := ioutil.ReadFile(f.Minecraftpath + "/versions/" + f.Version + "/" + f.Version + ".json")
	if err != nil {
		print(err)
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
