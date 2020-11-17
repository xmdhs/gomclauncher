package download

import (
	"testing"
)

func Test_randurls_auto(t *testing.T) {
	r := newrandurls("mcbbs")
	m := r.auto()
	if m != "mcbbs" {
		t.Fail()
	}
	r = newrandurls("")
	for i := 0; i < 20; i++ {
		t.Log(r.auto())
	}
}
