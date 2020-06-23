package download

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestNewlibrarie(t *testing.T) {
	b, err := ioutil.ReadFile("1.15.32.json")
	if err != nil {
		t.Fatal(err)
	}
	l, err := Newlibraries(b)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(l)
}

func TestAssets(t *testing.T) {
	bb, err := ioutil.ReadFile("1.16.json")
	if err != nil {
		t.Fatal(err)
	}
	a := assets{}
	json.Unmarshal(bb, &a)
	for s, v := range a.Objects {
		t.Log(s)
		t.Log(v)
	}
}
