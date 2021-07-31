package lang

import "testing"

func TestLang(t *testing.T) {
	err := Setlanguge("en")
	if err != nil {
		t.Fatal(err)
	}
	if Lang("nousername") != en["nousername"] {
		t.Fail()
	}
}
