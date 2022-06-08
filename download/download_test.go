package download

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/xmdhs/gomclauncher/launcher"
)

func TestNewlibraries(t *testing.T) {
	t.Cleanup(func() {
		os.RemoveAll(".minecraft")
	})
	b, err := ioutil.ReadFile("1.15.2.json")
	if err != nil {
		t.Fatal(err)
	}
	l, err := Newlibraries(context.Background(), b, "", func(s string) { fmt.Println(s) }, launcher.Minecraft)
	if err != nil {
		t.Fatal(err)
	}
	ch := make(chan int, 5)
	e := make(chan error)
	go func() {
		err = l.Downlibrarie(64, ch)
		if err != nil {
			e <- err
		}

	}()
b:
	for {
		select {
		case _, ok := <-ch:
			if !ok {
				break b
			}
			//fmt.Println(i)
		case err := <-e:
			t.Fatal(err)
			break b
		}
	}
	b, err = os.ReadFile(".minecraft/com/mojang/patchy/1.1/patchy-1.1.jar")
	if err != nil {
		t.Fatal(err)
	}
	h := sha1.New()
	h.Write(b)
	if hex.EncodeToString(h.Sum(nil)) != "aef610b34a1be37fa851825f12372b78424d8903" {
		t.Fatal("sha1 not match")
	}
}

func TestDownassets(t *testing.T) {
	b, err := ioutil.ReadFile("1.15.2.json")
	if err != nil {
		t.Fatal(err)
	}
	l, err := Newlibraries(context.Background(), b, "", func(s string) { fmt.Println(s) }, launcher.Minecraft)
	if err != nil {
		t.Fatal(err)
	}
	ch := make(chan int, 5)
	e := make(chan error)
	go func() {
		err = l.Downassets(64, ch)
		if err != nil {
			e <- err
		}

	}()
b:
	for {
		select {
		case _, ok := <-ch:
			if !ok {
				break b
			}
			//fmt.Println(i)
		case err := <-e:
			t.Fatal(err)
			break b
		}
	}
}
