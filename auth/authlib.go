package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const (
	Authlibversion = "1.1.34"
	Authlibsha1    = "4a75f3324aa5056e2d050680e2ceda51af8c7b2f"
)

var Authliburls = []string{
	"https://download.mcbbs.net/mirrors/authlib-injector/artifact/34/authlib-injector-1.1.34.jar",
	"https://authlib-injector.yushi.moe/artifact/34/authlib-injector-1.1.34.jar",
}

func Getauthlibapi(api string) (apiaddress string, err error) {
	if !strings.Contains(strings.ToTitle(api), strings.ToTitle("http")) {
		api = `https://` + api
	}
	c := &http.Client{
		Timeout:   5 * time.Second,
		Transport: Transport,
	}
	req, err := http.NewRequest("GET", api, nil)
	if err != nil {
		return "", fmt.Errorf("Getauthlibapi: %w", err)
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")
	reps, err := c.Do(req)
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
	if !strings.Contains(strings.ToTitle(header), strings.ToTitle("http")) {
		api = "https://" + reps.Request.URL.Host + "/" + header
	} else {
		return header, nil
	}
	return api, nil
}

type Yggdrasil struct {
	SignaturePublickey string `json:"signaturePublickey"`
}

func checkapi(url string) error {
	c := &http.Client{
		Timeout:   5 * time.Second,
		Transport: Transport,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("checkapi: %w", err)
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")
	reps, err := c.Do(req)
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
	var y Yggdrasil
	err = json.Unmarshal(b, &y)
	if err != nil {
		return fmt.Errorf("checkapi: %w", err)
	}
	if y.SignaturePublickey == "" {
		return JsonNotTrue
	}
	return nil
}

var JsonNotTrue = errors.New("json not true")
