package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func Refresh(a *Auth) error {
	r := Refreshs{
		AccessToken: a.AccessToken,
		ClientToken: a.ClientToken,
	}
	if a.selectedProfile.Name != "" {
		r.SelectedProfile = a.selectedProfile
	}
	b, err := json.Marshal(r)
	if err != nil {
		return fmt.Errorf("Refresh: %w", err)
	}
	b, err, i := post("refresh", b)
	if err != nil {
		return fmt.Errorf("Refresh: %w", err)
	}
	if i != http.StatusOK {
		return NotOk
	}
	err = json.Unmarshal(b, &r)
	a.AccessToken = r.AccessToken
	a.Username = r.SelectedProfile.Name
	a.ID = r.SelectedProfile.ID
	if err != nil {
		return fmt.Errorf("Refresh: %w", err)
	}
	return nil
}

type Refreshs struct {
	AccessToken     string          `json:"accessToken"`
	ClientToken     string          `json:"clientToken"`
	SelectedProfile SElectedProfile `json:"selectedProfile"`
}

type SElectedProfile struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

func Validate(a Auth) error {
	r := Refreshs{
		AccessToken: a.AccessToken,
		ClientToken: a.ClientToken,
	}
	b, err := json.Marshal(r)
	if err != nil {
		return fmt.Errorf("Validate: %w", err)
	}
	_, err, i := post("validate", b)
	if err != nil {
		return fmt.Errorf("Validate: %w", err)
	}
	if i == 204 {
		return nil
	}
	return AccessTokenCanNotUse
}

var AccessTokenCanNotUse = errors.New("accessToken is can not use")
