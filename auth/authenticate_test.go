package auth

import (
	"errors"
	"testing"
)

func TestAuthenticate(t *testing.T) {
	a, err := Authenticate("xmdhss@gmail.com", "K8JxiNtCFhG6R2n", "")
	if !errors.Is(err, NotOk) {
		t.Fatal(err)
	}
	t.Log(a)
}
