package download

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_randurls_auto(t *testing.T) {
	r := newrandurls("mcbbs")
	_, m := r.auto()
	if m != "mcbbs" {
		t.Fail()
	}
	r = newrandurls("")
	for i := 0; i < 20; i++ {
		t.Log(r.auto())
	}
}

func TestFail(t *testing.T) {
	r := newrandurls("")
	for i := 0; i < 20; i++ {
		r.fail("mcbbs")
		r.fail("bmclapi")
	}
	require.Equal(t, "vanilla", r.fail("mcbbs"))
}
