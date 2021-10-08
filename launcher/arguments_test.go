package launcher

import (
	"errors"
	"testing"
)

func TestGameinfo_legacy(t *testing.T) {
	err := ErrLegacyNoExit(errors.New(""))
	var e ErrLegacyNoExit
	if !errors.As(err, &e) {
		t.FailNow()
	}
}
