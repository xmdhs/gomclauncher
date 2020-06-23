package auth

import (
	"encoding/json"
	"errors"
	"net/http"
)

func Refresh(a *Auth) error {
	r := Refreshs{
		AccessToken: a.AccessToken,
		ClientToken: a.ClientToken,
	}
	b, err := json.Marshal(r)
	if err != nil {
		return err
	}
	b, err, i := post("refresh", b)
	if err != nil {
		return err
	}
	if i != http.StatusOK {
		return errors.New("not ok")
	}
	json.Unmarshal(b, &r)
	a.AccessToken = r.AccessToken
	a.Username = r.SelectedProfile.Name
	return nil
}

type Refreshs struct {
	AccessToken     string          `json:"accessToken"`
	ClientToken     string          `json:"clientToken"`
	SelectedProfile SElectedProfile `json:"selectedProfile"`
}

type SElectedProfile struct {
	Name string `json:"name"`
}

func Validate(a Auth) error {
	r := Refreshs{
		AccessToken: a.AccessToken,
		ClientToken: a.ClientToken,
	}
	b, err := json.Marshal(r)
	if err != nil {
		return err
	}
	_, err, i := post("validate", b)
	if i == 204 {
		return nil
	}
	return errors.New("accessToken is can not use")
}
