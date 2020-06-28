package flag

import (
	"encoding/json"
	"io/ioutil"
	"strings"
)

func Test(path string) bool {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return false
	}
	t := t{}
	json.Unmarshal(b, &t)
	if t.Libraries == nil {
		return false
	}
	if t.AssetIndex == nil {
		return false
	}
	p := strings.Split(path, "/")
	if t.ID != p[len(p)-1] {
		return false
	}
	return true
}

type t struct {
	Libraries  []interface{} `json:"Libraries"`
	AssetIndex []interface{} `json:"assetIndex"`
	ID         string        `json:"id"`
}

func Find(path string) []string {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	s := make([]string, 0)
	for _, f := range files {
		if f.IsDir() {
			s = append(s, f.Name())
		}
	}
	return s
}
