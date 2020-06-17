package auth

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func post(endpoint, Payload, proxyaddr string) ([]byte, error, int) {
	var c http.Client
	if proxyaddr != "" {
		proxy, err := url.Parse(proxyaddr)
		if err != nil {
			return nil, err, 0
		}
		c = http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxy),
			},
			Timeout: 10 * time.Second,
		}
	} else {
		c = http.Client{
			Timeout: 10 * time.Second,
		}
	}
	h, err := http.NewRequest("POST", "https://authserver.mojang.com/"+endpoint, strings.NewReader(Payload))
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
