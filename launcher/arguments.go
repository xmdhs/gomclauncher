package launcher

import (
	"errors"
	"strings"
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
		action := r["action"].(string)
		jvmrule.Action = action
		name, ok := r["os"].(map[string]interface{})["name"]
		if ok {
			jvmrule.Os = name.(string)
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
		if v.Action == "allow" && (v.Os == "" || osbool(v.Os)) {
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
	s = strings.ReplaceAll(s, "${game_assets}", g.Minecraftpath+`/assets`)
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
