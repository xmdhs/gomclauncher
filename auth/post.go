package auth

import (
	"bytes"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

var Transport = http.DefaultTransport.(*http.Transport).Clone()

var HttpClient *http.Client

func init() {
	Transport.DialContext = (&net.Dialer{
		Timeout:   5 * time.Second,
		KeepAlive: 30 * time.Second,
		DualStack: true,
	}).DialContext
	HttpClient = &http.Client{
		Transport: Transport,
	}
}

var ApiAddress string = "https://authserver.mojang.com"

func post(endpoint string, Payload []byte) ([]byte, error, int) {
	var api string
	if ApiAddress != "https://authserver.mojang.com" {
		api = ApiAddress + "/authserver"
	}
	if api == "" {
		api = "https://authserver.mojang.com"
	}
	h, err := http.NewRequest("POST", api+"/"+endpoint, bytes.NewReader(Payload))
	if err != nil {
		return nil, err, 0
	}
	h.Header.Set("Content-Type", "application/json")
	h.Header.Set("Accept", "*/*")
	h.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")
	rep, err := HttpClient.Do(h)
	if rep != nil {
		defer rep.Body.Close()
	}
	if err != nil {
		return nil, err, 0
	}
	b, err := ioutil.ReadAll(rep.Body)
	if err != nil {
		return nil, err, 0
	}
	return b, nil, rep.StatusCode
}
