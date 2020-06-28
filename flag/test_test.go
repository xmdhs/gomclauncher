package flag

import (
	"testing"
)

func TestTest(t *testing.T) {
	tt := Test("../download/1.15.2.json")
	if !tt {
		t.Fatal(tt)
	}
}
