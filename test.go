/*package main

import (
	"fmt"
	"gomclauncher/auth"
	"gomclauncher/download"
	"gomclauncher/launcher"
	"io/ioutil"
)

func main() {
	b, err := ioutil.ReadFile("./1.15.2.json")
	if err != nil {
		fmt.Println(err)
	}
	j := launcher.Gameinfo{
		Jsonbyte:      b,
		Minecraftpath: `D:/mc/.minecraft`,
		RAM:           `4096`,
		Name:          `Name`,
		UUID:          `9f51573a5ec545828c2b09f7f08497b1`,
		AccessToken:   "nil",
		GameDir:       `D:/mc/.minecraft`,
	}
	j.Run115()
	auth.Authenticate("a", "a", "a")
	download.Newlibraries(nil)
}
*/