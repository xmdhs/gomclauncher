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
