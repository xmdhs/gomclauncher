package launcher

import (
	"strings"
)

func Librarie2LibraryX115(l *Librarie) *LibraryX115 {
	Librarie := l.LibraryX115
	if l.Downloads.Artifact.URL != "" {
		return &Librarie
	} else {
		fullLibraryX115(&Librarie, l.Url)
		return &Librarie
	}
}

func fullLibraryX115(l *LibraryX115, url string) {
	s := Name2path(l.Name)
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
