package auth

import (
	"testing"
)

func TestRefresh(t *testing.T) {
	a, err := Authenticate("xmdhss@gmail.com", "K8JxiNtCFhG6R2n", "test")
	if err.Error() != "not ok" {
		t.Fatal(err)
	}
	err = Refresh(&a)
	if err.Error() != "not ok" {
		t.Fatal(err)
	}
}

func TestValidate(t *testing.T) {
	a, err := Authenticate("xmdhss@gmail.com", "K8JxiNtCFhG6R2n", "test")
	if err.Error() != "not ok" {
		t.Fatal(err)
	}
	err = Validate(a)
	if err.Error() != "accessToken is can not use" {
		t.Fatal(err)
	}
}
