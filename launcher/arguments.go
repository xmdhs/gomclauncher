package launcher

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
)

func (g *Gameinfo) argumentsjvm(l *launcher1155) error {
	j := l.json.Arguments.Jvm
	for _, v := range j {
		switch v := v.(type) {
		case map[string]interface{}:
			Jvm := rule(v)
			flags := jvmarguments(Jvm)
			for _, v := range flags {
				g.jvmflagadd(v, l)
			}
		case string:
			g.jvmflagadd(v, l)
		default:
			return JsonNorTrue
		}
	}
	addloggingXml(&l.json.Logging, l)
	return nil
}

var JsonNorTrue = errors.New("json not true")

func (g *Gameinfo) jvmflagadd(v string, l *launcher1155) {
	flag := g.jvmflagrelace(v, l)
	if v != "" {
		l.flag = append(l.flag, flag)
	}
}

func addloggingXml(log *loggingX115, l *launcher1155) {
	filename := filepath.Join(l.Minecraftpath, "/assets/logging/", log.Client.File.ID)
	arg := log.Client.Argument
	if log.Client.File.ID != "" {
		s := strings.ReplaceAll(arg, "${path}", filename)
		l.flag = append(l.flag, s)
	}
}

func (g *Gameinfo) jvmflagrelace(s string, l *launcher1155) string {
	s = strings.ReplaceAll(s, "${natives_directory}", g.Minecraftpath+`/versions/`+g.Version+`/natives`)
	s = strings.ReplaceAll(s, "${launcher_name}", Launcherbrand)
	s = strings.ReplaceAll(s, "${launcher_version}", Launcherversion)
	s = strings.ReplaceAll(s, "${library_directory}", g.Minecraftpath+"/libraries")
	s = strings.ReplaceAll(s, "${classpath_separator}", string(os.PathListSeparator))
	s = strings.ReplaceAll(s, "${version_name}", g.inheritsFrom)
	s = strings.ReplaceAll(s, "${classpath}", l.cp())
	return s
}

func rule(v map[string]interface{}) ajvm {
	jvm := ajvm{}
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
	rule := make([]jvmRule, 0)
	for _, rr := range r {
		jvmrule := jvmRule{}
		r := rr.(map[string]interface{})
		action, ok := r["action"].(string)
		if ok {
			jvmrule.Action = action
		}
		os := r["os"].(map[string]interface{})
		name, ok := os["name"]
		if ok {
			jvmrule.Os = name.(string)
		}
		arch, ok := os["arch"]
		if ok {
			jvmrule.arch = arch.(string)
		}
		rule = append(rule, jvmrule)
	}
	jvm.Rules = rule
	return jvm
}

func jvmarguments(j ajvm) []string {
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

type ajvm struct {
	Rules []jvmRule
	Value []string
}

type jvmRule struct {
	Action string
	Os     string
	arch   string
}

func (g *Gameinfo) argumentsGame(l *launcher1155) error {
	j := l.json.Arguments.Game
	for _, v := range j {
		argument, ok := v.(string)
		if ok {
			flag, err := g.argumentsrelace(argument, l)
			if err != nil {
				return fmt.Errorf("g.Gameinfo: %w", err)
			}
			if flag != "" {
				l.flag = append(l.flag, flag)
			}
		}
	}
	return nil
}

func (g *Gameinfo) argumentsrelace(s string, l *launcher1155) (string, error) {
	s = strings.ReplaceAll(s, "${auth_player_name}", g.Name)
	s = strings.ReplaceAll(s, "${version_name}", Launcherbrand+" "+Launcherversion)
	s = strings.ReplaceAll(s, "${game_directory}", g.Gamedir)
	s = strings.ReplaceAll(s, "${assets_root}", g.Minecraftpath+`/assets`)
	if strings.Contains(s, "${game_assets}") {
		err := g.legacy(l)
		if err != nil {
			return "", fmt.Errorf("g.argumentsrelace: %w", err)
		}
	}
	s = strings.ReplaceAll(s, "${game_assets}", g.Minecraftpath+`/assets/virtual/legacy`)
	s = strings.ReplaceAll(s, "${assets_index_name}", l.json.AssetIndex.ID)
	s = strings.ReplaceAll(s, "${auth_uuid}", g.UUID)
	s = strings.ReplaceAll(s, "${auth_access_token}", g.AccessToken)
	s = strings.ReplaceAll(s, "${auth_session}", g.AccessToken)
	s = strings.ReplaceAll(s, "${user_type}", "mojang")
	s = strings.ReplaceAll(s, "${version_type}", Launcherbrand+" "+Launcherversion)
	if g.Userproperties == "" {
		g.Userproperties = "{}"
	}
	s = strings.ReplaceAll(s, "${user_properties}", g.Userproperties)
	return s, nil
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

func (g *Gameinfo) legacy(l *launcher1155) error {
	p := g.Minecraftpath + `/assets/virtual/legacy/`
	b, err := ioutil.ReadFile(g.Minecraftpath + "/assets/indexes/" + l.json.AssetIndex.ID + ".json")
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("gameinfo.legacy: %w", ErrLegacyNoExit{rawErr: err})
		} else {
			return fmt.Errorf("gameinfo.legacy: %w", err)
		}
	}
	a := assets{}
	err = json.Unmarshal(b, &a)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("gameinfo.legacy: %w", ErrLegacyNoExit{rawErr: err})
		} else {
			return fmt.Errorf("gameinfo.legacy: %w", err)
		}
	}
	eCh := make(chan error, 10)
	cxt, c := context.WithCancel(context.TODO())
	defer c()

	var w sync.WaitGroup
	for path, v := range a.Objects {
		path, v := path, v
		w.Add(1)
		go func() {
			s := strings.Split(path, "/")
			ss := strings.ReplaceAll(path, s[len(s)-1], "")
			if a.Virtual {
				err = os.MkdirAll(p+ss, 0777)
			} else {
				err = os.MkdirAll(g.Gamedir+"/resources/"+ss, 0777)
			}
			if err != nil {
				if errors.Is(err, fs.ErrNotExist) {
					select {
					case eCh <- fmt.Errorf("gameinfo.legacy: %w", ErrLegacyNoExit{rawErr: err}):
					case <-cxt.Done():
					}
				} else {
					select {
					case eCh <- fmt.Errorf("gameinfo.legacy: %w", err):
					case <-cxt.Done():
					}
				}
			}
			f, err := os.Open(g.Minecraftpath + "/assets/objects/" + v.Hash[0:2] + "/" + v.Hash)
			if err != nil {
				if errors.Is(err, fs.ErrNotExist) {
					select {
					case eCh <- fmt.Errorf("gameinfo.legacy: %w", ErrLegacyNoExit{rawErr: err}):
					case <-cxt.Done():
					}
				} else {
					select {
					case eCh <- fmt.Errorf("gameinfo.legacy: %w", err):
					case <-cxt.Done():
					}
				}
			}
			defer f.Close()
			if a.Virtual {
				ff, err := os.Create(p + path)
				if err != nil {
					select {
					case eCh <- fmt.Errorf("gameinfo.legacy: %w", err):
					case <-cxt.Done():
					}
				}
				defer ff.Close()
				_, err = io.Copy(ff, f)
				if err != nil {
					select {
					case eCh <- fmt.Errorf("gameinfo.legacy: %w", err):
					case <-cxt.Done():
					}
				}
			} else {
				fff, err := os.Create(g.Gamedir + "/resources/" + path)
				if err != nil {
					select {
					case eCh <- fmt.Errorf("gameinfo.legacy: %w", err):
					case <-cxt.Done():
					}
				}
				defer fff.Close()
				_, err = io.Copy(fff, f)
				if err != nil {
					select {
					case eCh <- fmt.Errorf("gameinfo.legacy: %w", err):
					case <-cxt.Done():
					}
				}
			}
			w.Done()
		}()
	}

	go func() {
		w.Wait()
		c()
	}()

	select {
	case <-cxt.Done():
		return nil
	case err := <-eCh:
		return err
	}
}

type ErrLegacyNoExit struct {
	rawErr error
}

func (e ErrLegacyNoExit) Error() string {
	return e.rawErr.Error()
}

func (e ErrLegacyNoExit) Unwrap() error {
	return e.rawErr
}

type assets struct {
	Objects map[string]asset `json:"objects"`
	Virtual bool             `json:"virtual"`
}

type asset struct {
	Hash string `json:"hash"`
}
