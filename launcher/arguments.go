package launcher

import "strings"

func (g *Gameinfo) argumentsjvm(l *launcher1155) {
	j := l.json.Patches[0].Arguments.Jvm
	for _, v := range j {
		switch v := v.(type) {
		case map[string]interface{}:
			Jvm := Rule(v)
			flags := Jvmarguments(Jvm)
			if flags != nil {
				for _, v := range flags {
					if v != "" {
						l.flag = append(l.flag, g.Jvmflagrelace(v, l))
					}
				}
			}
		case string:
			if v != "" {
				l.flag = append(l.flag, g.Jvmflagrelace(v, l))
			}
		}
	}
}

func (g *Gameinfo) Jvmflagrelace(s string, l *launcher1155) string {
	s = strings.ReplaceAll(s, "${natives_directory}", g.Minecraftpath+`versions/`+g.Version+`/natives`)
	s = strings.ReplaceAll(s, "${launcher_name}", Launcherbrand)
	s = strings.ReplaceAll(s, "${launcher_version}", Launcherversion)
	s = strings.ReplaceAll(s, "${classpath}", "")
	s = strings.ReplaceAll(s, "-cp", "")
	return s
}

func Rule(v map[string]interface{}) Jvm {
	jvm := Jvm{}
	value := v["rules"].([]string)
	jvm.Value = value
	rules := v["rules"]
	r := rules.([]interface{})
	rule := make([]JvmRule, 0)
	for _, rr := range r {
		jvmrule := JvmRule{}
		r := rr.(map[string]interface{})
		action := r["action"].(string)
		jvmrule.Action = action
		jvmrule.Os = r["os"].(map[string]string)["name"]
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
