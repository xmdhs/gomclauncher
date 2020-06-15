package launcher

import (
	"encoding/json"
	"gomclauncher/launcher/launcherjson"
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

func (g Gameinfo) Run115() {
	j := launcherjson.LauncherjsonX115{}
	json.Unmarshal(g.Jsonbyte, &j)
	l := NewLauncher1155(j)
	l.Gameinfo = g
	l.flag = append(l.flag, `-Dminecraft.client.jar=`+g.Minecraftpath+`\versions\`+g.Version+`\`+g.Version+`.jar`)
	l.flag = append(l.flag, `-XX:+UseG1GC`)
	l.flag = append(l.flag, `-Xmx`+g.RAM+`m`)
	l.flag = append(l.flag, `-XX:-UseAdaptiveSizePolicy`)
	l.flag = append(l.flag, `-XX:-OmitStackTraceInFastThrow`)
	l.flag = append(l.flag, `-Dfml.ignoreInvalidMinecraftCertificates=true`)
	l.flag = append(l.flag, `-Dfml.ignorePatchDiscrepancies=true`)
	l.flag = append(l.flag, `-Djava.library.path=`+g.Minecraftpath+`\versions\`+g.Version+`\natives`)
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
	l.flag = append(l.flag, j.Patches[0].AssetIndex.ID)
	l.flag = append(l.flag, `--uuid`)
	l.flag = append(l.flag, g.UUID)
	l.flag = append(l.flag, `--accessToken`)
	l.flag = append(l.flag, g.AccessToken)
	l.flag = append(l.flag, `--userType`)
	l.flag = append(l.flag, ` mojang`)
	l.flag = append(l.flag, `--versionType`)
	l.flag = append(l.flag, Launcherbrand+" "+Launcherversion)
	l.Launcher115()
}
