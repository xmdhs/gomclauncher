package auth

import (
	"testing"
)

func TestAuthenticate(t *testing.T) {
	a, err := Authenticate("xmdhss@gmail.com", "K8JxiNtCFhG6R2n", "")
	if err == nil {
		t.Fail()
	}
	t.Log(a)
}
