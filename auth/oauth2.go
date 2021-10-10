package auth

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"time"
)

type msToken struct {
	AccessToken  string `json:"access_token,omitempty"`
	ExpiresIn    int    `json:"expires_in,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	Scope        string `json:"scope,omitempty"`
	TokenType    string `json:"token_type,omitempty"`
}

type MsToken struct {
	msToken
	ExpiresIn time.Time
}

func (m *MsToken) Expires() bool {
	return time.Now().After(m.ExpiresIn)
}

func (m *MsToken) Refresh() error {
	v := &url.Values{}
	v.Add("client_id", "a48a9fad-1702-46d7-8ee9-42b857ad292d")
	v.Add("scope", "XboxLive.signin offline_access")
	v.Add("refresh_token", m.RefreshToken)
	v.Add("grant_type", "refresh_token")
	b, err := httPost(oauth20Token, v.Encode(), "application/x-www-form-urlencoded")
	if err != nil {
		return fmt.Errorf("MsToken.Refresh: %w", err)
	}
	tm := msToken{}
	err = json.Unmarshal(b, &tm)
	if err != nil {
		return fmt.Errorf("MsToken.Refresh: %w", err)
	}
	m.parse(tm)
	return nil
}

func (m *MsToken) parse(mm msToken) {
	m.msToken = mm
	m.ExpiresIn = time.Now().Add(time.Duration(mm.ExpiresIn) * time.Second)
}

type ErrHttpCode struct {
	code int
	msg  string
}

func (e ErrHttpCode) Error() string {
	return "http code: " + strconv.Itoa(e.code) + " msg: " + e.msg
}
