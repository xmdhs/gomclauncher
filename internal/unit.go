package internal

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"path/filepath"
	"runtime"
	"runtime/debug"
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

var ua string

func init() {
	b, _ := debug.ReadBuildInfo()
	var hash string
	for _, v := range b.Settings {
		if v.Key == "vcs.revision" {
			hash = v.Value
		}
	}
	ua = fmt.Sprintf("gomclauncher/%s (%v)", Launcherversion, hash)
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
	rep.Header.Set("User-Agent", ua)
	c := http.Client{
		Transport: t,
	}
	reps, err := c.Do(rep)
	if err != nil {
		return reps, nil, fmt.Errorf("HttpGet: %w", err)
	}
	return reps, timer, nil
}

var ErrPathInvalid = errors.New("path invalid")

func SafePathJoin(base string, path ...string) (string, error) {
	p := filepath.Join(append([]string{base}, path...)...)
	a, err := filepath.Rel(base, p)
	if err != nil {
		return "", fmt.Errorf("SafePathJoin: %w", err)
	}
	if strings.HasPrefix(a, ".") {
		return "", ErrPathInvalid
	}
	return p, nil
}
