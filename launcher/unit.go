package launcher

import (
	"strings"

	"github.com/xmdhs/gomclauncher/internal"
)

func FullLibraryX115(l *LibraryX115, url string) [3]string {
	s := Name2path(l.Name)
	if url == "" {
		url = `https://libraries.minecraft.net/`
	}
	if (l.Downloads.Artifact.URL == "" || l.Downloads.Artifact.Path == "") && len(l.Extract.Exclude) == 0 {
		p := strings.ReplaceAll(s[0], ".", "/") + "/" + s[1] + "/" + s[2] + "/" + s[1] + "-" + s[2] + ".jar"
		l.Downloads.Artifact.Path = p
		l.Downloads.Artifact.URL = url + p
	}
	if len(l.Extract.Exclude) != 0 && len(l.Downloads.Classifiers) == 0 {
		l.Downloads.Classifiers = make(map[string]artifactX115, 3)
		do := func(os string) {
			if os == "" {
				return
			}
			l.Downloads.Classifiers[os] = artifactX115{
				Path: strings.ReplaceAll(s[0], ".", "/") + "/" + s[1] + "/" + s[2] + "/" + s[1] + "-" + s[2] + "-" + os + ".jar",
				URL:  url + s[0] + "/" + s[1] + "/" + s[2] + "/" + s[1] + "-" + s[2] + "-" + os + ".jar",
			}
		}
		arch := internal.Getarch()
		do(strings.ReplaceAll(l.Natives.Windows, "${arch}", arch))
		do(strings.ReplaceAll(l.Natives.Osx, "${arch}", arch))
		do(strings.ReplaceAll(l.Natives.Linux, "${arch}", arch))
	}
	return s
}
