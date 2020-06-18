package download

import (
	"io/ioutil"
	"testing"
)

func TestNewlibraries(t *testing.T) {
	b, err := ioutil.ReadFile("1.15.2.json")
	if err != nil {
		panic(err)
	}
	l, err := Newlibraries(b)
	if err != nil {
		panic(err)
	}
	err = l.Downlibrarie("")
	if err != nil {
		panic(err)
	}
}
