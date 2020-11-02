package download

import (
	"context"
	"testing"
)

func TestDownauthlib(t *testing.T) {
	err := Downauthlib(context.Background())
	if err != nil {
		t.Fatal(err)
	}
}
