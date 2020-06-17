package auth

import (
	"testing"
)

func TestRefresh(t *testing.T) {
	a, err := Authenticate("xmdhss@gmail.com", "K8JxiNtCFhG6R2n", "test")
	if err != nil {
		t.Fatal(err)
	}
	err = Refresh(&a)
	if err != nil {
		t.Fatal(err)
	}
}

func TestValidate(t *testing.T) {
	a, err := Authenticate("xmdhss@gmail.com", "K8JxiNtCFhG6R2n", "test")
	if err != nil {
		t.Fatal(err)
	}
	err = Validate(a)
	if err != nil {
		t.Fatal(err)
	}
}
