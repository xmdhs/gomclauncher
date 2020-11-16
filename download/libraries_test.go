package download

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/xmdhs/gomclauncher/auth"
)

func TestNewlibrarie(t *testing.T) {
	b, err := ioutil.ReadFile("1.15.2.json")
	if err != nil {
		t.Fatal(err)
	}
	l, err := Newlibraries(context.Background(), b, "")
	if err != nil {
		t.Fatal(err)
	}
	if l.librarie.ID != "1.15.2" {
		t.Fatal(l.librarie.ID)
	}
}

func TestAssets(t *testing.T) {
	bb, err := ioutil.ReadFile("1.16.json")
	if err != nil {
		t.Fatal(err)
	}
	a := assets{}
	err = json.Unmarshal(bb, &a)
	if err != nil {
		t.Fatal(err)
	}
}

func TestMain(m *testing.M) {
	auth.Transport = http.DefaultTransport.(*http.Transport).Clone()
}
