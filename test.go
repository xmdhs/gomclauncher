package main

import (
	"encoding/json"
	"fmt"
	"gomclauncher/launcher"
	"gomclauncher/launcher/launcherjson"
	"io/ioutil"
)

func main() {
	b, err := ioutil.ReadFile("./1.15.2.json")
	if err != nil {
		fmt.Println(err)
	}
	j := launcherjson.LauncherjsonX115{}
	json.Unmarshal(b, &j)
	lau := launcher.NewLauncher1155(j)
	lau.Launcher115()
}
