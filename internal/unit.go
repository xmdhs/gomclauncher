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
