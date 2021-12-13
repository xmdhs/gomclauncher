package internal

import (
	"runtime"
	"strings"

	"github.com/xmdhs/gomclauncher/launcher"
)

func Swichnatives(l launcher.LibraryX115) (path, sha1, url string) {
	Os := runtime.GOOS
	key := ""
	switch Os {
	case "windows":
		key = l.Natives.Windows
	case "darwin":
		key = l.Natives.Osx
	case "linux":
		key = l.Natives.Linux
	default:
		panic("???")
	}
	Arch := runtime.GOARCH
	arch := ""
	switch Arch {
	case "amd64":
		arch = "64"
	case "386":
		arch = "32"
	}
	key = strings.ReplaceAll(key, "${arch}", arch)
	a, ok := l.Downloads.Classifiers[key]
	if !ok {
		return "", "", ""
	}
	return a.Path, a.Sha1, a.URL
}

func Librarie2LibraryX115(l *launcher.Librarie) *launcher.LibraryX115 {
	Librarie := l.LibraryX115
	if l.Downloads.Artifact.URL != "" {
		return &Librarie
	} else {
		fullLibraryX115(&Librarie, l.Url)
		return &Librarie
	}
}

func fullLibraryX115(l *launcher.LibraryX115, url string) {
	s := launcher.Name2path(l.Name)
	if l.Downloads.Artifact.URL == "" || l.Downloads.Artifact.Path == "" {
		p := strings.ReplaceAll(s[0], ".", "/") + "/" + s[1] + "/" + s[2] + "/" + s[1] + "-" + s[2] + ".jar"
		l.Downloads.Artifact.Path = p
		if url != "" {
			l.Downloads.Artifact.URL = url + p
		} else {
			l.Downloads.Artifact.URL = `https://libraries.minecraft.net/` + p
		}
	}
}
