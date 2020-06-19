package download

import (
	"fmt"
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
	ch := make(chan int, 5)
	e := make(chan error)
	go func() {
		err = l.Downlibrarie("", 4, ch)
		if err != nil {
			e <- err
		}

	}()
b:
	for {
		select {
		case i, ok := <-ch:
			if !ok {
				break b
			}
			fmt.Println(i)
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
	l, err := Newlibraries(b)
	if err != nil {
		t.Fatal(err)
	}
	ch := make(chan int, 5)
	e := make(chan error)
	go func() {
		err = l.Downassets("", 4, ch)
		if err != nil {
			e <- err
		}

	}()
b:
	for {
		select {
		case i, ok := <-ch:
			if !ok {
				break b
			}
			fmt.Println(i)
		case err := <-e:
			t.Fatal(err)
			break b
		}
	}
}
