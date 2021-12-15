package internal

import "runtime"

func Getarch() string {
	Arch := runtime.GOARCH
	switch Arch {
	case "amd64":
		return "64"
	case "386":
		return "32"
	default:
		panic("???")
	}
}
