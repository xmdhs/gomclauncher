package launcher

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"
)

type Gameinfo struct {
	//D:/mc/.minecraft/
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
	GameDir string
	//1.15.2
	Version string
	//1.15.json []byte
	Jsonbyte []byte
	flag     map[string]string
}

func (g *Gameinfo) Run115() {
	l := g.modjson()
	l.flag = append(l.flag, `-Dminecraft.client.jar=`+g.Minecraftpath+`versions/`+l.json.ID+`/`+l.json.ID+`.jar`)
	l.flag = append(l.flag, `-XX:+UseG1GC`)
	l.flag = append(l.flag, `-Xmx`+g.RAM+`m`)
	l.flag = append(l.flag, `-XX:-UseAdaptiveSizePolicy`)
	l.flag = append(l.flag, `-XX:-OmitStackTraceInFastThrow`)
	l.flag = append(l.flag, `-Dfml.ignoreInvalidMinecraftCertificates=true`)
	l.flag = append(l.flag, `-Dfml.ignorePatchDiscrepancies=true`)
	g.argumentsjvm(l)
	l.flag = append(l.flag, l.json.Patches[0].MainClass)
	g.argumentsGame(l)
	l.Launcher115()
}

func (g *Gameinfo) modjson() *launcher1155 {
	g.flag = make(map[string]string)
	j := LauncherjsonX115{}
	mod := Modsjson{}
	var err error
	err = json.Unmarshal(g.Jsonbyte, &mod)
	if mod.InheritsFrom != "" {
		b, err := ioutil.ReadFile(g.Minecraftpath + `versions/` + mod.InheritsFrom + "/" + mod.InheritsFrom + ".json")
		if err != nil {
			log.Fatal(err)
		}
		err = json.Unmarshal(b, &j)
		for _, v := range mod.Libraries {
			l := g.Libraries2LibraryX115(v)
			j.Patches[0].Libraries = append(j.Patches[0].Libraries, l)
		}
		g.Version = mod.ID
		j.Patches[0].MainClass = mod.MainClass
		if len(mod.Arguments.Game) != 0 {
			j.Patches[0].Arguments.Game = append(j.Patches[0].Arguments.Game, mod.Arguments.Game...)
		}
	} else {
		err = json.Unmarshal(g.Jsonbyte, &j)
		g.Version = j.ID
	}
	if err != nil {
		log.Fatal(err)
	}

	l := NewLauncher1155(j)
	l.Gameinfo = g
	return l
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
	InheritsFrom string        `json:"inheritsFrom"`
	MainClass    string        `json:"mainClass"`
	ID           string        `json:"id"`
	Libraries    []Librarie    `json:"libraries"`
	Arguments    ArgumentsX115 `json:"arguments"`
}

type Librarie struct {
	Name      string `json:"name"`
	Url       string `json:"url"`
	Clientreq bool   `json:"clientreq"`
	Serverreq bool   `json:"serverreq"`
}
