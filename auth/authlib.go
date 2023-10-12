package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"

	"github.com/xmdhs/gomclauncher/internal"
)

const (
	// Deprecated: move to package download
	Authlibversion = "1.1.41"
	// Deprecated: move to package download
	Authlibsha1 = "64e96fc1e29e312a4d1e9d530f492cfa4b089cf1"
)

// Deprecated: move to package download
var Authliburls = []string{
	"https://download.mcbbs.net/mirrors/authlib-injector/artifact/41/authlib-injector-1.1.41.jar",
	"https://authlib-injector.yushi.moe/artifact/41/authlib-injector-1.1.41.jar",
	"https://bmclapi2.bangbang93.com/mirrors/authlib-injector/artifact/41/authlib-injector-1.1.41.jar",
}

func Getauthlibapi(api string) (apiaddress string, err error) {
	u, err := url.Parse(api)
	if err != nil {
		return "", fmt.Errorf("Getauthlibapi: %w", err)
	}
	if u.Scheme == "" {
		u.Scheme = "https"
		api = u.String()
	}
	reps, _, err := internal.HttpGet(context.TODO(), api, Transport, nil)
	if reps != nil {
		defer reps.Body.Close()
	}
	if err != nil {
		return "", fmt.Errorf("Getauthlibapi: %w", err)
	}
	defer func() {
		err = checkapi(apiaddress)
	}()
	header := reps.Header.Get("X-Authlib-Injector-API-Location")
	if header == "" {
		return api, nil
	}
	hurl, err := url.Parse(header)
	if err != nil {
		return "", fmt.Errorf("Getauthlibapi: %w", err)
	}
	if hurl.Scheme == "" {
		api, err = url.JoinPath("https://", reps.Request.URL.Host, header)
		if err != nil {
			return "", fmt.Errorf("Getauthlibapi: %w", err)
		}
	} else {
		return header, nil
	}
	return api, nil
}

type yggdrasil struct {
	SignaturePublickey string `json:"signaturePublickey"`
}

func checkapi(url string) error {
	reps, _, err := internal.HttpGet(context.TODO(), url, Transport, nil)
	if reps != nil {
		defer reps.Body.Close()
	}
	if err != nil {
		return fmt.Errorf("checkapi: %w", err)
	}
	b, err := ioutil.ReadAll(reps.Body)
	if err != nil {
		return fmt.Errorf("checkapi: %w", err)
	}
	var y yggdrasil
	err = json.Unmarshal(b, &y)
	if err != nil {
		return fmt.Errorf("checkapi: %w", err)
	}
	if y.SignaturePublickey == "" {
		return fmt.Errorf("checkapi: %w", JsonNotTrue)
	}
	return nil
}

var JsonNotTrue = errors.New("json not true")
