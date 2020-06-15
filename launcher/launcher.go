package launcher

import (
	"bytes"
	"gomclauncher/launcher/launcherjson"
	"os/exec"
	"runtime"
	"strings"
)

type Launcher1155 struct {
	json launcherjson.LauncherjsonX115
	flag []string
}

func NewLauncher1155(json launcherjson.LauncherjsonX115) Launcher1155 {
	flag := make([]string, 0)
	return Launcher1155{json: json, flag: flag}
}

func (l Launcher1155) Launcher115() {
	l.flag = append(l.flag, `-Dminecraft.client.jar=D:\mc\.minecraft\versions\1.15.2\1.15.2.jar`)
	l.flag = append(l.flag, `-XX:+UseG1GC`)
	l.flag = append(l.flag, `-Xmx4096m`)
	l.flag = append(l.flag, `-XX:-UseAdaptiveSizePolicy`)
	l.flag = append(l.flag, `-XX:-OmitStackTraceInFastThrow`)
	l.flag = append(l.flag, `-Dfml.ignoreInvalidMinecraftCertificates=true`)
	l.flag = append(l.flag, `-Dfml.ignorePatchDiscrepancies=true`)
	l.flag = append(l.flag, `-XX:HeapDumpPath=MojangTricksIntelDriversForPerformance_javaw.exe_minecraft.exe.heapdump`)
	l.flag = append(l.flag, `-Djava.library.path=D:\mc\.minecraft\versions\1.15.2\natives`)
	l.flag = append(l.flag, `-Dminecraft.launcher.brand=GML`)
	l.flag = append(l.flag, `-Dminecraft.launcher.version=beta0.1`)
	l.cp()
	cmd := exec.Command("java", l.flag...)
	cmd.Run()
}

func (l *Launcher1155) cp() {
	path := `D:\mc\.minecraft\libraries\`
	l.flag = append(l.flag, `-cp`)
	b := bytes.NewBuffer(nil)
	for _, p := range l.json.Patches[0].Libraries {
		if paths(p) != "" {
			b.WriteString(path)
			b.WriteString(paths(p))
			b.WriteString(";")
		}
	}
	b.WriteString(`D:\mc\.minecraft\versions\1.15.2\1.15.2.jar`)
	l.flag = append(l.flag, b.String())
	l.flag = append(l.flag, l.json.Patches[0].MainClass)
	l.flag = append(l.flag, `--username`)
	l.flag = append(l.flag, `xmdhs`)
	l.flag = append(l.flag, `--version`)
	l.flag = append(l.flag, `"GML beta0.1"`)
	l.flag = append(l.flag, `--gameDir`)
	l.flag = append(l.flag, `D:\mc\.minecraft`)
	l.flag = append(l.flag, `--assetsDir`)
	l.flag = append(l.flag, `D:\mc\.minecraft\assets`)
	l.flag = append(l.flag, `--assetIndex`)
	l.flag = append(l.flag, `1.15`)
	l.flag = append(l.flag, `--uuid 9f51573a5ec545828c2b09f7f08497b1`)
	l.flag = append(l.flag, `9f51573a5ec545828c2b09f7f08497b1`)
	l.flag = append(l.flag, `--accessToken`)
	l.flag = append(l.flag, `eyJhbGciOiJIUzI1NiJ9.eyJzdWIiOiI5MjljOTM1ZGIxODI0Y2I5ZmJlMTY2ZDM2NDIxZDU3ZSIsInlnZ3QiOiIyNmY0OGY0OGEwOWE0NTYxYmQxNDcyYjM3ZjE5ODgxZSIsInNwciI6IjlmNTE1NzNhNWVjNTQ1ODI4YzJiMDlmN2YwODQ5N2IxIiwiaXNzIjoiWWdnZHJhc2lsLUF1dGgiLCJleHAiOjE1OTIzNjExNzIsImlhdCI6MTU5MjE4ODM3Mn0.n29V8_McWOpe_9p2GObOEUwxTCZF89hiN5YjOlRN0D8`)
	l.flag = append(l.flag, `--userType`)
	l.flag = append(l.flag, ` mojang`)
	l.flag = append(l.flag, `--versionType`)
	l.flag = append(l.flag, `"GML beta0.1"`)

}

func paths(l launcherjson.LibraryX115) string {
	var allow bool
	if l.Rules != nil {
		for _, r := range l.Rules {
			if r.Action == "disallow" && osbool(r.Os.Name) {
				return ""
			}
			if r.Action == "allow" && r.Os.Name == "" {
				allow = true
			}
		}
		if !allow {
			return ""
		}
	}
	if runtime.GOOS == "windows" {
		a := strings.ReplaceAll(l.Downloads.Artifact.Path, `/`, `\`)
		return a
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
