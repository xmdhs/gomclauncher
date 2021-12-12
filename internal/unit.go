package internal

import (
	"runtime"

	"github.com/xmdhs/gomclauncher/launcher"
)

func Swichnatives(l launcher.LibraryX115) (path, sha1, url string) {
	Os := runtime.GOOS
	switch Os {
	case "windows":
		path = l.Downloads.Classifiers.NativesWindows.Path
		sha1 = l.Downloads.Classifiers.NativesWindows.Sha1
		url = l.Downloads.Classifiers.NativesWindows.URL
	case "darwin":
		if l.Downloads.Classifiers.NativesOsx.Path != "" {
			path = l.Downloads.Classifiers.NativesOsx.Path
			sha1 = l.Downloads.Classifiers.NativesOsx.Sha1
			url = l.Downloads.Classifiers.NativesOsx.URL
		} else {
			path = l.Downloads.Classifiers.NativesMacos.Path
			sha1 = l.Downloads.Classifiers.NativesMacos.Sha1
			url = l.Downloads.Classifiers.NativesMacos.URL
		}
	case "linux":
		path = l.Downloads.Classifiers.NativesLinux.Path
		sha1 = l.Downloads.Classifiers.NativesLinux.Sha1
		url = l.Downloads.Classifiers.NativesLinux.URL
	default:
		panic("???")
	}
	return
}
