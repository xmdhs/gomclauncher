package auth

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

//Authenticate return accessToken, err
func Authenticate(username, password, clientToken string) (Auth, error) {
	a := AuthenticatePayload{
		Agent: AuthenticateAgent{
			Name:    "Minecraft",
			Version: 1,
		},
		Username:    username,
		Password:    password,
		RequestUser: true,
		ClientToken: clientToken,
	}
	b, err := json.Marshal(a)
	Auth := Auth{}
	if err != nil {
		panic(err)
	}
	b, err, i := post("authenticate", b)
	if err != nil {
		return Auth, err
	}
	if i != http.StatusOK {
		return Auth, errors.New("not ok")
	}
	auth := &AuthenticateResponse{}
	if err = json.Unmarshal(b, auth); err != nil {
		panic(err)
	}
	w := bytes.NewBuffer(nil)
	var preferredLanguage, registrationCountry string
	for _, v := range auth.User.Properties {
		if v.Name == "preferredLanguage" {
			preferredLanguage = v.Value
		}
		if v.Name == "registrationCountry" {
			registrationCountry = v.Value
		}
	}
	w.WriteString(`"{\"preferredLanguage\":[\"`)
	w.WriteString(preferredLanguage)
	w.WriteString(`\"],\"registrationCountry\":[\"`)
	w.WriteString(registrationCountry)
	w.WriteString(`\"]}"`)
	Auth.AccessToken = auth.AccessToken
	Auth.ID = auth.SelectedProfile.ID
	Auth.ClientToken = auth.ClientToken
	Auth.Username = auth.SelectedProfile.Name
	Auth.Userproperties = w.String()
	return Auth, nil
}

type Auth struct {
	Username       string
	ClientToken    string
	ID             string
	AccessToken    string
	Userproperties string
}

type AuthenticatePayload struct {
	Agent       AuthenticateAgent `json:"agent"`
	ClientToken string            `json:"clientToken"`
	Password    string            `json:"password"`
	RequestUser bool              `json:"requestUser"`
	Username    string            `json:"username"`
}

type AuthenticateAgent struct {
	Name    string `json:"name"`
	Version int    `json:"version"`
}

type AuthenticateResponse struct {
	AccessToken       string                                 `json:"accessToken"`
	AvailableProfiles []AuthenticateResponseAvailableProfile `json:"availableProfiles"`
	ClientToken       string                                 `json:"clientToken"`
	SelectedProfile   AuthenticateResponseSelectedProfile    `json:"selectedProfile"`
	User              AuthenticateResponseUser               `json:"user"`
}

type AuthenticateResponseAvailableProfile struct {
	Agent         string  `json:"agent"`
	CreatedAt     float64 `json:"createdAt"`
	ID            string  `json:"id"`
	Legacy        bool    `json:"legacy"`
	LegacyProfile bool    `json:"legacyProfile"`
	Migrated      bool    `json:"migrated"`
	Name          string  `json:"name"`
	Paid          bool    `json:"paid"`
	Suspended     bool    `json:"suspended"`
	UserID        string  `json:"userId"`
}

type AuthenticateResponseSelectedProfile struct {
	CreatedAt     float64 `json:"createdAt"`
	ID            string  `json:"id"`
	Legacy        bool    `json:"legacy"`
	LegacyProfile bool    `json:"legacyProfile"`
	Migrated      bool    `json:"migrated"`
	Name          string  `json:"name"`
	Paid          bool    `json:"paid"`
	Suspended     bool    `json:"suspended"`
	UserID        string  `json:"userId"`
}

type AuthenticateResponseUser struct {
	Blocked           bool                               `json:"blocked"`
	DateOfBirth       float64                            `json:"dateOfBirth"`
	Email             string                             `json:"email"`
	EmailVerified     bool                               `json:"emailVerified"`
	ID                string                             `json:"id"`
	LegacyUser        bool                               `json:"legacyUser"`
	Migrated          bool                               `json:"migrated"`
	MigratedAt        float64                            `json:"migratedAt"`
	MigratedFrom      string                             `json:"migratedFrom"`
	PasswordChangedAt float64                            `json:"passwordChangedAt"`
	Properties        []AuthenticateResponseUserProperty `json:"properties"`
	RegisterIP        string                             `json:"registerIp"`
	RegisteredAt      float64                            `json:"registeredAt"`
	Secured           bool                               `json:"secured"`
	Suspended         bool                               `json:"suspended"`
	Username          string                             `json:"username"`
	VerifiedByParent  bool                               `json:"verifiedByParent"`
}

type AuthenticateResponseUserProperty struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
