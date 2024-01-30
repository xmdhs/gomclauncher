package internal

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

func Getarch() string {
	Arch := runtime.GOARCH
	switch Arch {
	case "amd64":
		return "64"
	case "386":
		return "32"
	default:
		panic("???")
	}
}

func HttpGet(cxt context.Context, aurl string, t *http.Transport, header http.Header) (*http.Response, *time.Timer, error) {
	ctx, cancel := context.WithCancel(cxt)
	rep, err := http.NewRequestWithContext(ctx, "GET", aurl, nil)
	timer := time.AfterFunc(5*time.Second, func() {
		cancel()
	})
	if err != nil {
		return nil, nil, fmt.Errorf("HttpGet: %w", err)
	}
	if header != nil {
		rep.Header = header
	}
	rep.Header.Set("Accept", "*/*")
	rep.Header.Set("User-Agent", "Mozilla/5.0（Windows NT 10.0；Win64；x64；rv:115.0） Gecko/20100101 Firefox/115.0")
	c := http.Client{
		Transport: t,
	}
	reps, err := c.Do(rep)
	if err != nil {
		return reps, nil, fmt.Errorf("HttpGet: %w", err)
	}
	return reps, timer, nil
}

var ErrPathNotInBase = errors.New("path not in base")

func SafePathJoin(base string, path ...string) (string, error) {
	p := filepath.Join(append([]string{base}, path...)...)
	a, err := filepath.Rel(base, p)
	if err != nil {
		return "", fmt.Errorf("SafePathJoin: %w", err)
	}
	if strings.HasPrefix(a, ".") {
		return "", ErrPathNotInBase
	}
	return p, nil
}
