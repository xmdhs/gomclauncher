package launcher

import (
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type Gameinfo struct {
	//D:/mc/.minecraft
	Minecraftpath string
	//4096
	RAM string
	//xmdhs
	Name string
	//9f51573a5ec545828c2b09f7f08497b1
	UUID string
	//eyJhbGciOiJIUzI1NiJ9
	AccessToken string
	//D:/mc/.minecraft/versions/1.15.2
	Gamedir string
	//1.15.2
	Version string
	//1.15.json []byte
	Jsonbyte []byte
	flag     map[string]string
	Flag     []string
	//"{\"preferredLanguage\":[\"zh-cn\"],\"registrationCountry\":[\"CN\"]}"
	Userproperties string
	Log            bool
	JavePath       string
	ApiAddress     string
	authlibpath    string

	inheritsFrom string

	jvmflagrelaceOnce sync.Once
	jvmflagrelaceMap  map[string]func() string
}

func (g *Gameinfo) Run115() (err error) {
	l, _, err := g.GenLauncherCmdArgs()
	if err != nil {
		return fmt.Errorf("Run115: %w", err)
	}
	err = l.Launcher115()
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("Run115: %w", ErrJavaPath)
		}
		return fmt.Errorf("Run115: %w", err)
	}
	return nil
}

var ErrJavaPath = errors.New("java path not exists")

func (g *Gameinfo) GenLauncherCmdArgs() (l *launcher1155, args []string, err error) {
	defer func() {
		e := recover()
		if e != nil {
			err = e.(error)
		}
	}()
	err = creatlauncherprofiles(g)
	if err != nil {
		return nil, nil, fmt.Errorf("Gameinfo.GenLauncherCmdArgs: %w", err)
	}
	l, err = g.modjson()
	if err != nil {
		return nil, nil, fmt.Errorf("Gameinfo.GenLauncherCmdArgs: %w", err)
	}
	l.flag = append(l.flag, `-Dminecraft.client.jar=`+g.Minecraftpath+`/versions/`+l.json.ID+`/`+l.json.ID+`.jar`)
	l.flag = append(l.flag, `-Xmx`+g.RAM+`m`)
	l.flag = append(l.flag, `-Xms`+g.RAM+`m`)
	if g.ApiAddress != "https://authserver.mojang.com" {
		l.flag = append(l.flag, `-Dauthlibinjector.side=client`)
		l.flag = append(l.flag, `-javaagent:`+g.authlibpath+`=`+g.ApiAddress)
	}
	if g.Flag != nil {
		l.flag = append(l.flag, g.Flag...)
	}
	err = g.argumentsjvm(l)
	if err != nil {
		return nil, nil, fmt.Errorf("Gameinfo.GenLauncherCmdArgs: %w", err)
	}
	if l.fixlog4j {
		log4j(l)
	}
	l.flag = append(l.flag, l.json.MainClass)
	err = g.argumentsGame(l)
	if err != nil {
		return nil, nil, fmt.Errorf("Gameinfo.GenLauncherCmdArgs: %w", err)
	}
	return l, l.flag, nil
}

//go:embed fixlog4j.jar
var fixlog4jJar []byte

// log4j to fix the CVE-2021-44228
func log4j(l *launcher1155) {
	path := filepath.Join(l.Minecraftpath, "fixlog4j-0.0.1.jar")
	_, err := os.Stat(path)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			err := os.WriteFile(path, fixlog4jJar, 0777)
			if err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	}
	l.flag = append(l.flag, `-javaagent:`+path)
}

func creatlauncherprofiles(g *Gameinfo) error {
	g.authlibpath = g.Minecraftpath + `libraries/moe/yushi/authlibinjector/authlib-injector/authlib-injector.jar`
	path := g.Minecraftpath + "/launcher_profiles.json"
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		f, err := os.Create(path)
		if err != nil {
			return fmt.Errorf("creatlauncherprofiles: %w", err)
		}
		defer f.Close()
		_, err = f.WriteString(`{"selectedProfile": "(Default)","profiles": {"(Default)": {"name": "(Default)"}},"clientToken": "88888888-8888-8888-8888-888888888888"}`)
		if err != nil {
			return fmt.Errorf("creatlauncherprofiles: %w", err)
		}
	}
	return nil
}

var JsonErr = errors.New("json err")

func (g *Gameinfo) modjson() (*launcher1155, error) {
	g.flag = make(map[string]string)
	j := LauncherjsonX115{}
	mod := Modsjson{}
	var err error
	err = json.Unmarshal(g.Jsonbyte, &mod)
	if err != nil {
		return nil, JsonErr
	}
	g.inheritsFrom = mod.InheritsFrom
	if mod.InheritsFrom != "" {
		b, err := ioutil.ReadFile(g.Minecraftpath + `/versions/` + mod.InheritsFrom + "/" + mod.InheritsFrom + ".json")
		if err != nil {
			return nil, fmt.Errorf("gameinfo.modjson: %w", err)
		}
		err = json.Unmarshal(b, &j)
		if err != nil {
			return nil, JsonErr
		}
		for _, v := range mod.Libraries {
			l := g.libraries2LibraryX115(v)
			j.Libraries = append(j.Libraries, l)
		}
		g.Version = mod.ID
		j.MainClass = mod.MainClass
		if len(mod.Arguments.Game) != 0 {
			j.Arguments.Game = append(j.Arguments.Game, mod.Arguments.Game...)
		}
		if len(mod.Arguments.Jvm) != 0 {
			j.Arguments.Jvm = append(j.Arguments.Jvm, mod.Arguments.Jvm...)
		}
		if mod.MinecraftArguments != "" {
			j.Arguments.Game = minecraftArguments2jvm(mod.MinecraftArguments)
			j.Arguments.Jvm = getjvm()
		}
		if mod.Logging != nil {
			j.Logging = *mod.Logging
		}
	} else {
		err = json.Unmarshal(g.Jsonbyte, &j)
		if err != nil {
			return nil, JsonErr
		}
		g.Version = j.ID
	}
	if j.MinecraftArguments != "" && (len(j.Arguments.Jvm) == 0 || len(j.Arguments.Game) == 0) {
		j.Arguments.Jvm = getjvm()
		j.Arguments.Game = minecraftArguments2jvm(j.MinecraftArguments)
	}
	l := newlauncher1155(j)
	l.Gameinfo = g
	return l, nil
}

func getjvm() []interface{} {
	return []interface{}{
		"-Djava.library.path=${natives_directory}",
		"-Dminecraft.launcher.brand=${launcher_name}",
		"-Dminecraft.launcher.version=${launcher_version}",
		"-cp",
		"${classpath}",
	}
}

func minecraftArguments2jvm(m string) []interface{} {
	l := strings.Split(m, " ")
	i := make([]interface{}, 0)
	for _, v := range l {
		i = append(i, v)
	}
	return i
}

func (g *Gameinfo) libraries2LibraryX115(l Librarie) LibraryX115 {
	p := Name2path(l.Name)
	g.flag[p[0]+p[1]] = p[2]
	FullLibraryX115(&l.LibraryX115, l.Url)
	return l.LibraryX115
}

// Name2path return [<group>,<name>,<version>]
func Name2path(name string) [3]string {
	l := strings.Split(name, ":")
	if len(l) != 3 {
		return [3]string{name, name, "1.0"}
	}
	return *(*[3]string)(l)

}

type Modsjson struct {
	//1.15.2
	InheritsFrom string `json:"inheritsFrom"`
	patchX115
	Libraries []Librarie   `json:"libraries"`
	Logging   *loggingX115 `json:"logging"`
}

type Librarie struct {
	Url       string `json:"url"`
	Clientreq bool   `json:"clientreq"`
	Serverreq bool   `json:"serverreq"`
	LibraryX115
}
