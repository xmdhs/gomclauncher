package flag

import (
	"encoding/json"
	"io/ioutil"
)

func Test(path string) bool {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return false
	}
	t := t{}
	err = json.Unmarshal(b, &t)
	if err != nil {
		return false
	}
	if len(t.Libraries) == 0 {
		return false
	}
	if t.MainClass == "" {
		return false
	}
	return true
}

type t struct {
	Libraries []interface{} `json:"Libraries"`
	MainClass string        `json:"mainClass"`
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
