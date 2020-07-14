package auth

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const Authlibversion = "1.1.27-5ef5f8e"

func Getauthlibapi(api string) (apiaddress string, err error) {
	if !strings.Contains(strings.ToTitle(api), strings.ToTitle("http")) {
		api = `https://` + api
	}
	c := &http.Client{
		Timeout: 5 * time.Second,
	}
	req, err := http.NewRequest("GET", api, nil)
	if err != nil {
		return "", err
	}
	reps, err := c.Do(req)
	if reps != nil {
		defer reps.Body.Close()
	}
	if err != nil {
		return "", err
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
		Timeout: 5 * time.Second,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	reps, err := c.Do(req)
	if reps != nil {
		defer reps.Body.Close()
	}
	if err != nil {
		return err
	}
	b, err := ioutil.ReadAll(reps.Body)
	if err != nil {
		return err
	}
	var y Yggdrasil
	err = json.Unmarshal(b, &y)
	if y.SignaturePublickey == "" {
		return errors.New("json not true")
	}
	return err
}
