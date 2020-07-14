package download

import (
	"testing"
)

func TestDownauthlib(t *testing.T) {
	err := Downauthlib()
	if err != nil {
		t.Fatal(err)
	}
}
