package download

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/avast/retry-go/v4"
	"github.com/xmdhs/gomclauncher/lang"
)

func Getversionlist(cxt context.Context, atype string, print func(string)) (*version, error) {
	var b []byte
	r := newrandurls(atype)
	_, f := r.auto()

	err := retry.Do(func() error {
		url := source(`https://piston-meta.mojang.com/mc/game/version_manifest.json`, f)
		rep, _, err := Aget(cxt, url)
		if rep != nil {
			defer rep.Body.Close()
		}
		if err != nil {
			f = r.fail(f)
			return fmt.Errorf("%v %w %v", lang.Lang("getversionlistfail"), err, url)
		}
		b, err = io.ReadAll(rep.Body)
		if err != nil {
			f = r.fail(f)
			return fmt.Errorf("%v %w %v", lang.Lang("getversionlistfail"), err, url)
		}
		return nil
	}, append(retryOpts, retry.OnRetry(func(n uint, err error) {
		print(fmt.Sprintf("retry %d: %v", n, err))
	}))...)
	if err != nil {
		return nil, fmt.Errorf("Getversionlist: %w %w", err, FileDownLoadFail)
	}
	v := version{}
	err = json.Unmarshal(b, &v)
	v.atype = atype
	if err != nil {
		return nil, fmt.Errorf("Getversionlist: %w", err)
	}
	return &v, nil
}

type version struct {
	Latest   versionLatest    `json:"latest"`
	Versions []versionVersion `json:"versions"`
	atype    string
}

type versionLatest struct {
	Release  string `json:"release"`
	Snapshot string `json:"snapshot"`
}

type versionVersion struct {
	ID          string `json:"id"`
	ReleaseTime string `json:"releaseTime"`
	Time        string `json:"time"`
	Type        string `json:"type"`
	URL         string `json:"url"`
}

func (v version) Downjson(cxt context.Context, version, apath string, print func(string)) error {
	r := newrandurls(v.atype)
	_, f := r.auto()
	for _, vv := range v.Versions {
		if vv.ID == version {
			s := strings.Split(vv.URL, "/")
			path := apath + `/versions/` + vv.ID + `/` + vv.ID + `.json`
			if ver(path, s[len(s)-2]) {
				return nil
			}

			err := retry.Do(func() error {
				url := source(vv.URL, f)
				err := get(cxt, url, path)
				if err != nil {
					f = r.fail(f)
					return fmt.Errorf("%v %v %w", lang.Lang("weberr"), url, err)
				}
				if !ver(path, s[len(s)-2]) {
					f = r.fail(f)
					return fmt.Errorf("%v %v", lang.Lang("filecheckerr"), url)
				}
				return nil
			}, append(retryOpts, retry.OnRetry(func(n uint, err error) {
				print(fmt.Sprintf("retry %d: %v", n, err))
			}))...)
			if err != nil {
				return fmt.Errorf("Downjson: %w %w", err, FileDownLoadFail)
			}
			return nil
		}
	}
	return NoSuch
}

var (
	NoSuch         = errors.New("no such")
	ErrFileChecker = errors.New("file checker")
)
