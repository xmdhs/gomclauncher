package download

import (
	"context"
	"fmt"
	"testing"
)

func TestDownauthlib(t *testing.T) {
	err := Downauthlib(context.Background(), func(s string) { fmt.Println(s) })
	if err != nil {
		t.Fatal(err)
	}
}
