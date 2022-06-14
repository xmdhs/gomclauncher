package download

import (
	"runtime"
	"testing"
)

func Test_isArch1_19(t *testing.T) {
	if !isArch1_19("org.lwjgl:lwjgl-jemalloc:3.3.1:natives-windows-x86") && runtime.GOARCH == "amd64" {
		t.Log("pass")
	} else {
		t.Error("pass")
	}
}
