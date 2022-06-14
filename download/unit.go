package download

import (
	"runtime"
	"strings"

	"github.com/xmdhs/gomclauncher/internal"
	"github.com/xmdhs/gomclauncher/launcher"
)

func swichnatives(l launcher.LibraryX115) (path, sha1, url string) {
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
	if key == "" && len(l.Rules) > 0 && launcher.Ifallow(l) && isArch1_19(l.Name) {
		return l.Downloads.Artifact.Path, l.Downloads.Artifact.Sha1, l.Downloads.Artifact.URL
	}

	key = strings.ReplaceAll(key, "${arch}", internal.Getarch())
	a, ok := l.Downloads.Classifiers[key]
	if !ok {
		return "", "", ""
	}
	return a.Path, a.Sha1, a.URL
}

func librarie2LibraryX115(l *launcher.Librarie) *launcher.LibraryX115 {
	Librarie := l.LibraryX115
	if l.Downloads.Artifact.URL != "" {
		return &Librarie
	} else {
		launcher.FullLibraryX115(&Librarie, l.Url)
		return &Librarie
	}
}

func isArch1_19(s string) bool {
	l := strings.Split(s, "-")
	if len(l) != 2 {
		return true
	}
	arch := l[len(l)-1]

	switch arch {
	case "x86":
		return runtime.GOARCH == "386"
	case "aarch_64", "arm64":
		return runtime.GOARCH == "arm64"
	default:
		return runtime.GOARCH == "amd64"
	}
}
