package auth

import (
	"strings"
)

func Getauthlibapi(api string) string {
	if strings.Contains(api, "http") {
		api = `https://` + api
	}
	return ""
}
