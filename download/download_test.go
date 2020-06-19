package download

import (
	"io/ioutil"
	"testing"
)

func TestNewlibraries(t *testing.T) {
	b, err := ioutil.ReadFile("1.15.2.json")
	if err != nil {
		t.Fatal(err)
	}
	l, err := Newlibraries(b)
	if err != nil {
		t.Fatal(err)
	}
	err = l.Downlibrarie("", 4)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDownassets(t *testing.T) {
	b, err := ioutil.ReadFile("1.15.2.json")
	if err != nil {
		t.Fatal(err)
	}
	l, err := Newlibraries(b)
	if err != nil {
		t.Fatal(err)
	}
	err = l.Downassets("", 100)
	if err != nil {
		t.Fatal(err)
	}
}
