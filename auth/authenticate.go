package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var (
	NotOk      = errors.New("not ok")
	NoProfiles = errors.New("无可用角色")
)

//Authenticate return accessToken, err
func Authenticate(ApiAddress, username, email, password, clientToken string) (Auth, error) {
	if ApiAddress == "" {
		ApiAddress = "https://authserver.mojang.com"
	}
	a := authenticatePayload{
		Agent: authenticateAgent{
			Name:    "Minecraft",
			Version: 1,
		},
		Username:    email,
		Password:    password,
		RequestUser: true,
		ClientToken: clientToken,
	}
	b, err := json.Marshal(a)
	Auth := Auth{}
	Auth.ApiAddress = ApiAddress
	if err != nil {
		panic(err)
	}
	b, err, i := post(Auth.ApiAddress, "authenticate", b)
	if err != nil {
		return Auth, fmt.Errorf("Authenticate: %w", err)
	}
	if i != http.StatusOK {
		return Auth, NotOk
	}
	auth := &authenticateResponse{}
	if err = json.Unmarshal(b, auth); err != nil {
		panic(err)
	}
	Auth.AccessToken = auth.AccessToken
	Auth.ClientToken = auth.ClientToken
	Auth.availableProfiles = auth.AvailableProfiles
	if len(Auth.availableProfiles) == 0 {
		return Auth, NoProfiles
	}
	if auth.SelectedProfile.Name == "" {
		a, err := selectProfile(Auth.availableProfiles, username)
		if err != nil {
			return Auth, fmt.Errorf("Authenticate: %w", err)
		}
		Auth.selectedProfile = a
		err = Refresh(&Auth)
		if err != nil {
			return Auth, fmt.Errorf("Authenticate: %w", err)
		}
	} else {
		Auth.ID = auth.SelectedProfile.ID
		Auth.Username = auth.SelectedProfile.Name
	}
	return Auth, nil
}

func selectProfile(a []authenticateResponseAvailableProfile, username string) (sElectedProfile, error) {
	if username == "" {
		return sElectedProfile{}, ErrNotSelctProFile
	}
	var selectedProfile authenticateResponseAvailableProfile
	for _, p := range a {
		if p.Name == username {
			selectedProfile = p
		}
	}
	if selectedProfile.Name == "" {
		return sElectedProfile{}, ErrProFileNoExist
	}
	s := sElectedProfile{
		Name: selectedProfile.Name,
		ID:   selectedProfile.ID,
	}
	return s, nil
}

func ListAvailableProfileName(a Auth) []string {
	list := []string{}
	for _, v := range a.availableProfiles {
		list = append(list, v.Name)
	}
	return list
}

var (
	ErrNotSelctProFile = errors.New("ErrNotSelctProFile")
	ErrProFileNoExist  = errors.New("ErrProFileNoExist")
)

type Auth struct {
	Username          string
	ClientToken       string
	ID                string
	AccessToken       string
	selectedProfile   sElectedProfile
	ApiAddress        string
	availableProfiles []authenticateResponseAvailableProfile
}

type authenticatePayload struct {
	Agent       authenticateAgent `json:"agent"`
	ClientToken string            `json:"clientToken"`
	Password    string            `json:"password"`
	RequestUser bool              `json:"requestUser"`
	Username    string            `json:"username"`
}

type authenticateAgent struct {
	Name    string `json:"name"`
	Version int    `json:"version"`
}

type authenticateResponse struct {
	AccessToken       string                                 `json:"accessToken"`
	AvailableProfiles []authenticateResponseAvailableProfile `json:"availableProfiles"`
	ClientToken       string                                 `json:"clientToken"`
	SelectedProfile   authenticateResponseSelectedProfile    `json:"selectedProfile"`
	User              authenticateResponseUser               `json:"user"`
}

type authenticateResponseAvailableProfile struct {
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

type authenticateResponseSelectedProfile struct {
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

type authenticateResponseUser struct {
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
	Properties        []authenticateResponseUserProperty `json:"properties"`
	RegisterIP        string                             `json:"registerIp"`
	RegisteredAt      float64                            `json:"registeredAt"`
	Secured           bool                               `json:"secured"`
	Suspended         bool                               `json:"suspended"`
	Username          string                             `json:"username"`
	VerifiedByParent  bool                               `json:"verifiedByParent"`
}

type authenticateResponseUserProperty struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
