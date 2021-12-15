package launcher

import (
	"testing"
)

func TestFullLibraryX115(t *testing.T) {
	l := &LibraryX115{
		Name: "com.mojang:patchy:1.1",
	}
	FullLibraryX115(l, "")
	if l.Downloads.Artifact.URL != "https://libraries.minecraft.net/com/mojang/patchy/1.1/patchy-1.1.jar" {
		t.Fatal("Artifact URL is not correct", l.Downloads.Artifact.URL)
	}
	l = &LibraryX115{
		Name: "net.java.jinput:jinput-platform:2.0.5",
		Extract: extractX115{
			Exclude: []string{"META-INF/"},
		},
		Natives: nativesX115{
			Linux:   "natives-linux",
			Osx:     "natives-osx",
			Windows: "natives-windows",
		},
	}
	FullLibraryX115(l, "")
	if l.Downloads.Classifiers["natives-linux"].URL != "https://libraries.minecraft.net/net/java/jinput/jinput-platform/2.0.5/jinput-platform-2.0.5-natives-linux.jar" {
		t.Fatal("Natives Linux URL is not correct", l.Downloads.Classifiers["natives-linux"].URL)
	}

	l = &LibraryX115{
		Name: "net.java.jinput:jinput-platform:2.0.5",
		Downloads: downloadsX115{
			Artifact: artifactX115{
				URL:  "aaa",
				Path: "bbb",
			},
		},
	}
	FullLibraryX115(l, "")
	if l.Downloads.Artifact.URL != "aaa" || l.Downloads.Artifact.Path != "bbb" {
		t.Fatal("Artifact URL is not correct", l.Downloads.Artifact.URL)
	}
}
