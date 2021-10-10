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

	mm := MsToken{
		msToken: msToken{
			AccessToken:  "",
			ExpiresIn:    0,
			RefreshToken: "M.R3_BAY.-CR7L!O4eEjAhTfifb51254mjMAQRjgxBGQpCA*tXwpZnkg4HvDCmjpA2spBmIRj8NF*bx8Io8sXFZDkNwbg*6m3el38zn2SmHSXbg6V32q1wXs*I9OuzBE!nHJR3!XNE7RZfIUIIteIo5ofKyaBF1MqbdqioGJJfhBdGmOeGE8rJQOYHr0I34cp4I!LfuH2GhV3AzBn9hzyD9eyHebooOKzRxSrdMxp7laYlOaqfDoxW0r*IH0dQyR5YazBWpq9!ymyDlm*jPKEBAqdQ0Rrvq8PAYPn4tEBx5cXhdLB2YKj3",
			Scope:        "XboxLive.signin offline_access",
			TokenType:    "",
		},
		ExpiresIn: time.Time{},
	}
	err = mm.Refresh()
	if err != nil {
		t.Fail()
	}
}
