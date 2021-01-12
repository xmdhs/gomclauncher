package download

import (
	"context"
	"io/ioutil"
	"os"
	"testing"
)

func TestNewlibraries(t *testing.T) {
	t.Cleanup(func() {
		os.RemoveAll(".minecraft")
	})
	b, err := ioutil.ReadFile("1.15.2.json")
	if err != nil {
		t.Fatal(err)
	}
	l, err := Newlibraries(context.Background(), b, "")
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
}

func TestDownassets(t *testing.T) {
	b, err := ioutil.ReadFile("1.15.2.json")
	if err != nil {
		t.Fatal(err)
	}
	l, err := Newlibraries(context.Background(), b, "")
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
