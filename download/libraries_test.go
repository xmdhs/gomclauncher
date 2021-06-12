package download

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/xmdhs/gomclauncher/auth"
	"github.com/xmdhs/gomclauncher/launcher"
)

func TestNewlibrarie(t *testing.T) {
	b, err := ioutil.ReadFile("1.15.2.json")
	if err != nil {
		t.Fatal(err)
	}
	l, err := Newlibraries(context.Background(), b, "", func(s string) { fmt.Println(s) }, launcher.Minecraft)
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
	code := m.Run()
	os.Exit(code)
}

func Test_get(t *testing.T) {
	err := get(context.Background(), "https://launchermeta.mojang.com/mc/game/version_manifest.json", "test/test/a.json")
	if err != nil {
		t.Fatal(err)
		return
	}
	_, err = os.Stat("test/test/a.json")
	if err != nil {
		t.Fatal(err)
		return
	}
	os.RemoveAll("test")
}
