package auth

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

var Proxyaddr string

func post(endpoint string, Payload []byte) ([]byte, error, int) {
	var c http.Client
	if Proxyaddr != "" {
		proxy, err := url.Parse(Proxyaddr)
		if err != nil {
			return nil, errors.New("proxy err"), 0
		}
		transport := http.DefaultTransport.(*http.Transport)
		transport.Proxy = http.ProxyURL(proxy)
		c = http.Client{
			Transport: transport,
			Timeout:   10 * time.Second,
		}
	} else {
		c = http.Client{
			Timeout: 10 * time.Second,
		}
	}
	h, err := http.NewRequest("POST", "https://authserver.mojang.com/"+endpoint, bytes.NewReader(Payload))
	if err != nil {
		return nil, err, 0
	}
	h.Header.Set("Content-Type", "application/json")
	h.Header.Set("Accept", "*/*")
	h.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")
	rep, err := c.Do(h)
	if err != nil {
		return nil, err, 0
	}
	b, err := ioutil.ReadAll(rep.Body)
	defer rep.Body.Close()
	if err != nil {
		return nil, err, 0
	}
	return b, nil, rep.StatusCode
}
