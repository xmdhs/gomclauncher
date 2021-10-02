package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
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
	v.Add("client_id", "00000000402b5328")
	v.Add("scope", "service::user.auth.xboxlive.com::MBI_SSL")
	v.Add("refresh_token", m.RefreshToken)
	v.Add("grant_type", "refresh_token")

	resp, err := c.Get(oauth20Token + "?" + v.Encode())
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return fmt.Errorf("MsToken.Refresh: %w", err)
	}
	b, err := io.ReadAll(resp.Body)
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
	m.ExpiresIn = time.Now().Add(time.Duration(mm.ExpiresIn) + time.Second)
}
