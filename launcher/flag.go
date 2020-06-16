package launcher

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"
)

type Gameinfo struct {
	//D:\mc\.minecraft\
	Minecraftpath string
	//4096
	RAM string
	//xmdhs
	Name string
	//9f51573a5ec545828c2b09f7f08497b1
	UUID string
	//eyJhbGciOiJIUzI1NiJ9
	AccessToken string
	//D:\mc\.minecraft\versions\1.15.2
	GameDir string
	//1.15.2
	Version string
	//1.15
	Jsonbyte []byte
}

func (g *Gameinfo) Run115() {
	l := g.modjson()
	g.argumentsjvm(l)
	l.flag = append(l.flag, `-Dminecraft.client.jar=`+g.Minecraftpath+`versions/`+l.json.ID+`/`+l.json.ID+`.jar`)
	l.flag = append(l.flag, `-XX:+UseG1GC`)
	l.flag = append(l.flag, `-Xmx`+g.RAM+`m`)
	l.flag = append(l.flag, `-XX:-UseAdaptiveSizePolicy`)
	l.flag = append(l.flag, `-XX:-OmitStackTraceInFastThrow`)
	l.flag = append(l.flag, `-Dfml.ignoreInvalidMinecraftCertificates=true`)
	l.flag = append(l.flag, `-Dfml.ignorePatchDiscrepancies=true`)
	l.flag = append(l.flag, `-Djava.library.path=`+g.Minecraftpath+`versions/`+g.Version+`/natives`)
	l.flag = append(l.flag, `-Dminecraft.launcher.brand=`+Launcherbrand)
	l.flag = append(l.flag, `-Dminecraft.launcher.version=`+Launcherversion)
	l.flag = append(l.flag, `-cp`)
	l.flag = append(l.flag, l.cp())
	l.flag = append(l.flag, l.json.Patches[0].MainClass)
	l.flag = append(l.flag, `--username`)
	l.flag = append(l.flag, g.Name)
	l.flag = append(l.flag, `--version`)
	l.flag = append(l.flag, Launcherbrand+" "+Launcherversion)
	l.flag = append(l.flag, `--gameDir`)
	l.flag = append(l.flag, g.GameDir)
	l.flag = append(l.flag, `--assetsDir`)
	l.flag = append(l.flag, g.Minecraftpath+`assets`)
	l.flag = append(l.flag, `--assetIndex`)
	l.flag = append(l.flag, l.json.Patches[0].AssetIndex.ID)
	l.flag = append(l.flag, `--uuid`)
	l.flag = append(l.flag, g.UUID)
	l.flag = append(l.flag, `--accessToken`)
	l.flag = append(l.flag, g.AccessToken)
	l.flag = append(l.flag, `--userType`)
	l.flag = append(l.flag, `mojang`)
	l.flag = append(l.flag, `--versionType`)
	l.flag = append(l.flag, Launcherbrand+" "+Launcherversion)
	l.Launcher115()
}

func (g *Gameinfo) modjson() *launcher1155 {
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
			l := Libraries2LibraryX115(v)
			j.Patches[0].Libraries = append(j.Patches[0].Libraries, l)
		}
		g.Version = mod.ID
		j.Patches[0].MainClass = mod.MainClass
		if len(mod.Arguments.Game) != 0 {
			j.Patches[0].Arguments.Game = mod.Arguments.Game
		}
		if len(mod.Arguments.Jvm) != 0 {
			j.Patches[0].Arguments.Jvm = mod.Arguments.Jvm
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

func Libraries2LibraryX115(l Librarie) LibraryX115 {
	p := Name2path(l.Name)
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
