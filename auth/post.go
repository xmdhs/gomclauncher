package auth

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

var Transport = http.DefaultTransport.(*http.Transport).Clone()

func init() {
	Transport.TLSClientConfig = &tls.Config{
		Renegotiation: tls.RenegotiateOnceAsClient,
	}
	//microsoft auth neet this
}

func post(ApiAddress, endpoint string, Payload []byte) ([]byte, error, int) {
	var api string
	if ApiAddress != "https://authserver.mojang.com" {
		var err error
		api, err = url.JoinPath(ApiAddress, "/authserver")
		if err != nil {
			return nil, fmt.Errorf("post: %w", err), 0
		}
	}
	if api == "" {
		api = "https://authserver.mojang.com"
	}
	h, err := http.NewRequest("POST", api+"/"+endpoint, bytes.NewReader(Payload))
	if err != nil {
		return nil, fmt.Errorf("post: %w", err), 0
	}
	h.Header.Set("Content-Type", "application/json")
	h.Header.Set("Accept", "*/*")
	h.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")
	c := &http.Client{
		Timeout:   5 * time.Second,
		Transport: Transport,
	}
	rep, err := c.Do(h)
	if rep != nil {
		defer rep.Body.Close()
	}
	if err != nil {
		return nil, fmt.Errorf("post: %w", err), 0
	}
	b, err := ioutil.ReadAll(rep.Body)
	if err != nil {
		return nil, fmt.Errorf("post: %w", err), 0
	}
	return b, nil, rep.StatusCode
}
