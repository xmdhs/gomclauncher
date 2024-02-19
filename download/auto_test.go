package download

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_randurls_auto(t *testing.T) {
	r := newrandurls("bmclapi")
	_, m := r.auto()
	if m != "bmclapi" {
		t.Fail()
	}
	r = newrandurls("")
	for i := 0; i < 20; i++ {
		t.Log(r.auto())
	}
}

func TestFail(t *testing.T) {
	r := newrandurls("")
	for i := 0; i < 50; i++ {
		r.fail("bmclapi")
	}
	require.Equal(t, "vanilla", r.fail("bmclapi"))
}
