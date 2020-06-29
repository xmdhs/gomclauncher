package launcher

import (
	"bytes"
	"fmt"
	"os/exec"
	"runtime"
)

type launcher1155 struct {
	json LauncherjsonX115
	flag []string
	*Gameinfo
}

func NewLauncher1155(json LauncherjsonX115) *launcher1155 {
	flag := make([]string, 0)
	return &launcher1155{json: json, flag: flag}
}

func (l launcher1155) Launcher115() {
	fmt.Println(l.flag)
	cmd := exec.Command("java", l.flag...)
	cmd.Dir = l.Gamedir
	b, err := cmd.CombinedOutput()
	fmt.Println(string(b))
	if err != nil {
		panic(err)
	}
}

func (l *launcher1155) cp() string {
	path := l.Minecraftpath + `/libraries/`
	b := bytes.NewBuffer(nil)
	for _, p := range l.json.Libraries {
		if paths(p) != "" {
			pack := Name2path(p.Name)
			v, ok := l.Gameinfo.flag[pack[0]]
			if ok {
				if v == pack[2] {
					b.WriteString(path)
					b.WriteString(p.Downloads.Artifact.Path)
					b.WriteString(";")
				}
			} else {
				b.WriteString(path)
				b.WriteString(p.Downloads.Artifact.Path)
				b.WriteString(";")
			}
		}
	}
	b.WriteString(l.Minecraftpath + `/versions/` + l.json.ID + `/` + l.json.ID + `.jar`)
	return b.String()
}

func paths(l LibraryX115) string {
	if l.Rules != nil {
		var allow bool
		for _, r := range l.Rules {
			if r.Action == "disallow" && osbool(r.Os.Name) {
				return ""
			}
			if r.Action == "allow" && (r.Os.Name == "" || osbool(r.Os.Name)) {
				allow = true
			}
		}
		if !allow {
			return ""
		}
	}
	return l.Downloads.Artifact.Path
}

func osbool(os string) bool {
	GOOS := runtime.GOOS
	if GOOS == "darwin" {
		GOOS = "osx"
	}
	return os == GOOS
}
