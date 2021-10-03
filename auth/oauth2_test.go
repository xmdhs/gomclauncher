package auth

import (
	"testing"
	"time"
)

func TestMsToken_Refresh(t *testing.T) {
	m := MsToken{
		msToken:   msToken{},
		ExpiresIn: time.Time{},
	}
	err := m.Refresh()
	if err == nil {
		t.Fail()
	}
}
