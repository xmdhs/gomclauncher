package launcher

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/xmdhs/gomclauncher/auth"
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
}

func (g *Gameinfo) Run115() error {
	creatlauncherprofiles(g)
	l, err := g.Modjson()
	if err != nil {
		return fmt.Errorf("Run115: %w", err)
	}
	l.flag = append(l.flag, `-Dminecraft.client.jar=`+g.Minecraftpath+`/versions/`+l.json.ID+`/`+l.json.ID+`.jar`)
	l.flag = append(l.flag, `-XX:+UseG1GC`)
	l.flag = append(l.flag, `-Xmx`+g.RAM+`m`)
	l.flag = append(l.flag, `-XX:-UseAdaptiveSizePolicy`)
	l.flag = append(l.flag, `-XX:-OmitStackTraceInFastThrow`)
	l.flag = append(l.flag, `-Dfml.ignoreInvalidMinecraftCertificates=true`)
	l.flag = append(l.flag, `-Dfml.ignorePatchDiscrepancies=true`)
	if auth.ApiAddress != "https://authserver.mojang.com" {
		l.flag = append(l.flag, `-Dauthlibinjector.side=client`)
		l.flag = append(l.flag, `-javaagent:`+authlibpath+`=`+auth.ApiAddress)
	}
	if g.Flag != nil {
		l.flag = append(l.flag, g.Flag...)
	}
	err = g.argumentsjvm(l)
	if err != nil {
		return fmt.Errorf("Run115: %w", err)
	}
	l.flag = append(l.flag, l.json.MainClass)
	g.argumentsGame(l)
	l.Launcher115()
	return nil
}

var authlibpath string

func creatlauncherprofiles(g *Gameinfo) {
	authlibpath = g.Minecraftpath + `/libraries/` + `moe/yushi/authlibinjector/` + "authlib-injector/" + auth.Authlibversion + "/authlib-injector-" + auth.Authlibversion + ".jar"
	path := g.Minecraftpath + "/launcher_profiles.json"
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		f, err := os.Create(path)
		defer f.Close()
		if err != nil {
			panic(err)
		}
		f.WriteString(`{"selectedProfile": "(Default)","profiles": {"(Default)": {"name": "(Default)"}},"clientToken": "88888888-8888-8888-8888-888888888888"}`)
	}
}

var JsonErr = errors.New("json err")

func (g *Gameinfo) Modjson() (*launcher1155, error) {
	g.flag = make(map[string]string)
	j := LauncherjsonX115{}
	mod := Modsjson{}
	var err error
	err = json.Unmarshal(g.Jsonbyte, &mod)
	if err != nil {
		return nil, JsonErr
	}
	if mod.InheritsFrom != "" {
		b, err := ioutil.ReadFile(g.Minecraftpath + `/versions/` + mod.InheritsFrom + "/" + mod.InheritsFrom + ".json")
		if err != nil {
			if os.IsNotExist(err) {
				return nil, err
			}
			panic(err)
		}
		err = json.Unmarshal(b, &j)
		if err != nil {
			return nil, JsonErr
		}
		for _, v := range mod.Libraries {
			l := g.Libraries2LibraryX115(v)
			j.Libraries = append(j.Libraries, l)
		}
		g.Version = mod.ID
		j.MainClass = mod.MainClass
		if len(mod.Arguments.Game) != 0 {
			j.Arguments.Game = append(j.Arguments.Game, mod.Arguments.Game...)
		}
		if mod.MinecraftArguments != "" {
			j.Arguments.Game = append(j.Arguments.Game, MinecraftArguments2jvm(mod.MinecraftArguments)...)
			j.Arguments.Jvm = append(j.Arguments.Jvm, getjvm()...)
		}

	} else {
		err = json.Unmarshal(g.Jsonbyte, &j)
		if err != nil {
			return nil, JsonErr
		}
		if j.MinecraftArguments != "" {
			j.Arguments.Game = append(j.Arguments.Game, MinecraftArguments2jvm(j.MinecraftArguments)...)
			j.Arguments.Jvm = append(j.Arguments.Jvm, getjvm()...)
		}
		g.Version = j.ID
	}
	l := NewLauncher1155(j)
	l.Gameinfo = g
	return l, nil
}

func getjvm() []interface{} {
	cp := []string{"-Djava.library.path\u003d${natives_directory}", "-Dminecraft.launcher.brand\u003d${launcher_name}", "-Dminecraft.launcher.version\u003d${launcher_version}", "-cp", "${classpath}"}
	i := make([]interface{}, 0)
	for _, v := range cp {
		i = append(i, v)
	}
	return i
}

func MinecraftArguments2jvm(m string) []interface{} {
	l := strings.Split(m, " ")
	i := make([]interface{}, 0)
	for _, v := range l {
		i = append(i, v)
	}
	return i
}

func (g *Gameinfo) Libraries2LibraryX115(l Librarie) LibraryX115 {
	p := Name2path(l.Name)
	g.flag[p[0]] = p[2]
	return LibraryX115{
		Downloads: DownloadsX115{
			Artifact: ArtifactX115{
				//<package>/<name>/<version>/<name>-<version>.jar
				Path: strings.ReplaceAll(p[0], ".", "/") + "/" + p[1] + "/" + p[2] + "/" + p[1] + "-" + p[2] + ".jar",
			},
		},
	}

}

func Name2path(name string) []string {
	return strings.Split(name, ":")
}

type Modsjson struct {
	//1.15.2
	MinecraftArguments string        `json:"minecraftArguments"`
	InheritsFrom       string        `json:"inheritsFrom"`
	MainClass          string        `json:"mainClass"`
	ID                 string        `json:"id"`
	Libraries          []Librarie    `json:"libraries"`
	Arguments          ArgumentsX115 `json:"arguments"`
}

type Librarie struct {
	Name      string        `json:"name"`
	Url       string        `json:"url"`
	Clientreq bool          `json:"clientreq"`
	Serverreq bool          `json:"serverreq"`
	Downloads DownloadsX115 `json:"downloads"`
}
