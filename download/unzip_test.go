package download

import (
	"context"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestUnzip(t *testing.T) {
	b, err := ioutil.ReadFile("1.15.2.json")
	if err != nil {
		t.Fatal(err)
	}
	l, err := Newlibraries(context.Background(), b, "", func(s string) { fmt.Println(s) })
	if err != nil {
		t.Fatal(err)
	}
	err = l.Unzip(20)
	if err != nil {
		t.Fatal(err)
	}
}
