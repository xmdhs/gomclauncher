package launcher

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

type launcher1155 struct {
	json LauncherjsonX115
	flag []string
	*Gameinfo
}

var Log bool

func NewLauncher1155(json LauncherjsonX115) *launcher1155 {
	flag := make([]string, 0)
	return &launcher1155{json: json, flag: flag}
}

func (l launcher1155) Launcher115() {
	fmt.Println(l.flag)
	var cmd *exec.Cmd
	if Log {
		cmd = exec.Command("java", l.flag...)
		cmd.Stdout = os.Stdout
		cmd.Dir = l.Gamedir
		err := cmd.Run()
		if err != nil {
			panic(err)
		}
	} else {
		if runtime.GOOS == "windows" {
			cmd = exec.Command("javaw", l.flag...)
		} else {
			l.flag = append(l.flag, "&")
			cmd = exec.Command("java", l.flag...)
		}
		cmd.Dir = l.Gamedir
		err := cmd.Start()
		if err != nil {
			panic(err)
		}
	}
}

func (l *launcher1155) cp() string {
	path := l.Minecraftpath + `/libraries/`
	b := bytes.NewBuffer(nil)
	for _, p := range l.json.Libraries {
		if Ifallow(p) {
			pack := Name2path(p.Name)
			v, ok := l.Gameinfo.flag[pack[0]]
			add := func() {
				b.WriteString(path)
				b.WriteString(p.Downloads.Artifact.Path)
				if runtime.GOOS == "windows" {
					b.WriteString(";")
				} else {
					b.WriteString(":")
				}
			}
			if ok {
				if v == pack[2] {
					add()
				}
			} else {
				add()
			}
		}
	}
	b.WriteString(l.Minecraftpath + `/versions/` + l.json.ID + `/` + l.json.ID + `.jar`)
	return b.String()
}

func Ifallow(l LibraryX115) bool {
	if l.Rules != nil {
		var allow bool
		for _, r := range l.Rules {
			if r.Action == "disallow" && osbool(r.Os.Name) {
				return false
			}
			if r.Action == "allow" && (r.Os.Name == "" || osbool(r.Os.Name)) {
				allow = true
			}
		}
		return allow
	}
	return true
}

func osbool(os string) bool {
	GOOS := runtime.GOOS
	if GOOS == "darwin" {
		GOOS = "osx"
	}
	return os == GOOS
}
