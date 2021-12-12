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

func (l *launcher1155) GetLauncherjsonX115() LauncherjsonX115 {
	return l.json
}

func newlauncher1155(json LauncherjsonX115) *launcher1155 {
	flag := make([]string, 0)
	return &launcher1155{json: json, flag: flag}
}

func (l launcher1155) Launcher115() error {
	fmt.Println(l.flag)
	var cmd *exec.Cmd
	if l.JavePath == "" {
		l.JavePath = "java"
	}
	if l.Log {
		cmd = exec.Command(l.JavePath, l.flag...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Dir = l.Gamedir
		err := cmd.Run()
		if err != nil {
			if err != nil {
				return fmt.Errorf("launcher1155.Launcher115: %w", err)
			}
		}
	} else {
		if runtime.GOOS == "windows" && l.JavePath == "java" {
			cmd = exec.Command("javaw", l.flag...)
		} else {
			l.flag = append(l.flag, "&")
			cmd = exec.Command(l.JavePath, l.flag...)
		}
		cmd.Dir = l.Gamedir
		err := cmd.Start()
		if err != nil {
			return fmt.Errorf("launcher1155.Launcher115: %w", err)
		}
	}
	return nil
}

func (l *launcher1155) cp() string {
	path := l.Minecraftpath + `/libraries/`
	b := bytes.NewBuffer(nil)
	for _, p := range l.json.Libraries {
		if Ifallow(p) {
			pack := Name2path(p.Name)
			v, ok := l.Gameinfo.flag[pack[0]+pack[1]]
			add := func() {
				b.WriteString(path)
				b.WriteString(p.Downloads.Artifact.Path)
				b.WriteByte(os.PathListSeparator)
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

func (l *launcher1155) CP() []string {
	path := l.Minecraftpath + `/libraries/`
	list := make([]string, 0, len(l.json.Libraries))
	for _, p := range l.json.Libraries {
		pack := Name2path(p.Name)
		v, ok := l.Gameinfo.flag[pack[0]+pack[1]]
		add := func() {
			list = append(list, path+p.Downloads.Artifact.Path)
		}
		if ok {
			if v == pack[2] {
				add()
			}
		} else {
			add()
		}

	}
	return list
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
