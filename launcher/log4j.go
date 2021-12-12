package launcher

import (
	"github.com/Masterminds/semver/v3"
)

func needFixlog4j(ver string) bool {
	v, err := semver.NewVersion(ver)
	if err != nil {
		return true
	}
	if v.Major() >= 2 && v.LessThan(semver.MustParse("2.15.0")) {
		return true
	}
	return false
}

func Fixlog4j(v *LibraryX115) {
	l := Name2path(v.Name)
	if _, ok := log4j[l[1]]; !ok {
		return
	}
	if !needFixlog4j(l[2]) {
		return
	}
	log := log4j[l[1]]
	v.Name = log.Name
	v.Downloads = log.Downloads
}

var log4j map[string]LibraryX115 = map[string]LibraryX115{
	"log4j-api": {
		Name: "org.apache.logging.log4j:log4j-api:2.15.0",
		Downloads: downloadsX115{
			Artifact: artifactX115{
				URL:  "https://maven.minecraftforge.net/org/apache/logging/log4j/log4j-api/2.15.0/log4j-api-2.15.0.jar",
				Sha1: "4a5aa7e55a29391c6f66e0b259d5189aa11e45d0",
				Path: "org/apache/logging/log4j/log4j-api/2.15.0/log4j-api-2.15.0.jar",
				Size: 301804,
			},
		},
	},
	"log4j-core": {
		Name: "org.apache.logging.log4j:log4j-core:2.15.0",
		Downloads: downloadsX115{
			Artifact: artifactX115{
				URL:  "https://maven.minecraftforge.net/org/apache/logging/log4j/log4j-core/2.15.0/log4j-core-2.15.0.jar",
				Sha1: "ba55c13d7ac2fd44df9cc8074455719a33f375b9",
				Path: "org/apache/logging/log4j/log4j-core/2.15.0/log4j-core-2.15.0.jar",
				Size: 1789769,
			},
		},
	},
	"log4j-slf4j18-impl": {
		Name: "org.apache.logging.log4j:log4j-slf4j18-impl:2.15.0",
		Downloads: downloadsX115{
			Artifact: artifactX115{
				URL:  "https://maven.minecraftforge.net/org/apache/logging/log4j/log4j-slf4j18-impl/2.15.0/log4j-slf4j18-impl-2.15.0.jar",
				Sha1: "88f72ad364bfc3a7cf43186fc17212f2b4bb8d97",
				Path: "org/apache/logging/log4j/log4j-slf4j18-impl/2.15.0/log4j-slf4j18-impl-2.15.0.jar",
				Size: 21223,
			},
		},
	},
}
