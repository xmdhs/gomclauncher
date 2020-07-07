package launcher

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"runtime"
	"strings"
	"sync"
)

func (g *Gameinfo) argumentsjvm(l *launcher1155) error {
	j := l.json.Arguments.Jvm
	for _, v := range j {
		switch v := v.(type) {
		case map[string]interface{}:
			Jvm := Rule(v)
			flags := Jvmarguments(Jvm)
			if flags != nil {
				for _, v := range flags {
					g.jvmflagadd(v, l)
				}
			}
		case string:
			g.jvmflagadd(v, l)
		default:
			return errors.New("json not true")
		}
	}
	return nil
}

func (g *Gameinfo) jvmflagadd(v string, l *launcher1155) {
	flag := g.Jvmflagrelace(v, l)
	if v != "" {
		l.flag = append(l.flag, flag)
	}
}

func (g *Gameinfo) Jvmflagrelace(s string, l *launcher1155) string {
	s = strings.ReplaceAll(s, "${natives_directory}", g.Minecraftpath+`/versions/`+g.Version+`/natives`)
	s = strings.ReplaceAll(s, "${launcher_name}", Launcherbrand)
	s = strings.ReplaceAll(s, "${launcher_version}", Launcherversion)
	s = strings.ReplaceAll(s, "${classpath}", l.cp())
	return s
}

func Rule(v map[string]interface{}) Jvm {
	jvm := Jvm{}
	var values []interface{}
	switch vv := v["value"].(type) {
	case []interface{}:
		values = append(values, vv...)
	case string:
		values = append(values, vv)
	}
	value := make([]string, 0)
	for _, v := range values {
		value = append(value, v.(string))
	}
	jvm.Value = value
	rules := v["rules"]
	r := rules.([]interface{})
	rule := make([]JvmRule, 0)
	for _, rr := range r {
		jvmrule := JvmRule{}
		r := rr.(map[string]interface{})
		action, ok := r["action"].(string)
		if ok {
			jvmrule.Action = action
		}
		name, ok := r["os"].(map[string]interface{})["name"]
		if ok {
			jvmrule.Os = name.(string)
		}
		arch, ok := r["os"].(map[string]interface{})["arch"]
		if ok {
			jvmrule.arch = arch.(string)
		}
		rule = append(rule, jvmrule)
	}
	jvm.Rules = rule
	return jvm
}

func Jvmarguments(j Jvm) []string {
	var allow bool
	for _, v := range j.Rules {
		if v.Action == "disallow" && osbool(v.Os) {
			return nil
		}
		if v.Action == "allow" && (v.Os == "" || osbool(v.Os)) && (v.arch == "" || archbool(v.arch)) {
			allow = true
		}
	}
	if allow {
		return j.Value
	}
	return nil
}

type Jvm struct {
	Rules []JvmRule
	Value []string
}

type JvmRule struct {
	Action string
	Os     string
	arch   string
}

func (g *Gameinfo) argumentsGame(l *launcher1155) {
	j := l.json.Arguments.Game
	for _, v := range j {
		argument, ok := v.(string)
		if ok {
			flag := g.argumentsrelace(argument, l)
			if flag != "" {
				l.flag = append(l.flag, flag)
			}
		}
	}
}

func (g *Gameinfo) argumentsrelace(s string, l *launcher1155) string {
	s = strings.ReplaceAll(s, "${auth_player_name}", g.Name)
	s = strings.ReplaceAll(s, "${version_name}", Launcherbrand+" "+Launcherversion)
	s = strings.ReplaceAll(s, "${game_directory}", g.Gamedir)
	s = strings.ReplaceAll(s, "${assets_root}", g.Minecraftpath+`/assets`)
	if strings.Contains(s, "${game_assets}") {
		g.legacy(l)
	}
	s = strings.ReplaceAll(s, "${game_assets}", g.Minecraftpath+`/assets/virtual/legacy`)
	s = strings.ReplaceAll(s, "${assets_index_name}", l.json.AssetIndex.ID)
	s = strings.ReplaceAll(s, "${auth_uuid}", g.UUID)
	s = strings.ReplaceAll(s, "${auth_access_token}", g.AccessToken)
	s = strings.ReplaceAll(s, "${user_type}", "mojang")
	s = strings.ReplaceAll(s, "${version_type}", Launcherbrand+" "+Launcherversion)
	if g.Userproperties == "" {
		g.Userproperties = "{}"
	}
	s = strings.ReplaceAll(s, "${user_properties}", g.Userproperties)
	return s
}

func archbool(arch string) bool {
	if arch == "x86" {
		if runtime.GOARCH == "386" {
			return true
		}
	} else {
		if runtime.GOARCH == "amd64" {
			return true
		}
	}
	return false
}

func (g *Gameinfo) legacy(l *launcher1155) {
	p := g.Minecraftpath + `/assets/virtual/legacy/`
	fileerr := func(err error) {
		if err != nil {
			fmt.Println("文件不存在，请开启文件效验")
			os.Exit(0)
		}
	}
	var f *os.File
	f, err := os.Open(g.Minecraftpath + "/assets/indexes/" + l.json.AssetIndex.ID + ".json")
	fileerr(err)
	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	a := assets{}
	json.Unmarshal(b, &a)
	var w sync.WaitGroup
	for path, v := range a.Objects {
		path, v := path, v
		w.Add(1)
		go func() {
			s := strings.Split(path, "/")
			ss := strings.ReplaceAll(path, s[len(s)-1], "")
			if l.json.AssetIndex.ID != "legacy" {
				err = os.MkdirAll(g.Gamedir+"/resources/"+ss, 0777)
			} else {
				err = os.MkdirAll(p+ss, 0777)
			}
			if err != nil {
				panic(err)
			}
			f, err := os.Open(g.Minecraftpath + "/assets/objects/" + v.Hash[0:2] + "/" + v.Hash)
			fileerr(err)
			if l.json.AssetIndex.ID != "legacy" {
				fff, err := os.Create(g.Gamedir + "/resources/" + path)
				if err != nil {
					panic(err)
				}
				_, err = io.Copy(fff, f)
			} else {
				ff, err := os.Create(p + path)
				if err != nil {
					panic(err)
				}
				_, err = io.Copy(ff, f)
			}
			w.Done()
		}()
	}
	w.Wait()
}

type assets struct {
	Objects map[string]asset `json:"objects"`
}

type asset struct {
	Hash string `json:"hash"`
}
