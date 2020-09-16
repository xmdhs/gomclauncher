package auth

import (
	"errors"
	"testing"
)

func TestRefresh(t *testing.T) {
	a, err := Authenticate("xmdhss@gmail.com", "K8JxiNtCFhG6R2n", "test")
	if !errors.Is(err, NotOk) {
		t.Fatal(err)
	}
	err = Refresh(&a)
	if !errors.Is(err, NotOk) {
		t.Fatal(err)
	}
}

func TestValidate(t *testing.T) {
	a, err := Authenticate("xmdhss@gmail.com", "K8JxiNtCFhG6R2n", "test")
	if !errors.Is(err, NotOk) {
		t.Fatal(err)
	}
	err = Validate(a)
	if !errors.Is(err, AccessTokenCanNotUse) {
		t.Fatal(err)
	}
}
