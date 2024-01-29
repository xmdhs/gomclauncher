package download

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"time"

	"github.com/avast/retry-go/v4"
	"github.com/xmdhs/gomclauncher/lang"
)

func Downauthlib(cxt context.Context, print func(string)) (err error) {
	var path = ".minecraft/libraries/moe/yushi/authlibinjector/authlib-injector/authlib-injector.jar"
	url := ""

	err = retry.Do(func() error {
		url = randAuthlibUrls(url)
		var d, h string
		d, h, err = getAuthlibLatestUrl(cxt, url)
		if err != nil {
			return fmt.Errorf("%v %w %v", lang.Lang("authlibdownloadfail"), err, url)
		}
		if ver(path, h) {
			return nil
		}
		err = get(cxt, d, path)
		if err != nil {
			return fmt.Errorf("%v %w %v", lang.Lang("authlibdownloadfail"), err, url)
		}
		if !ver(path, h) {
			return fmt.Errorf("%v %v", lang.Lang("authlibcheckerr"), url)
		}
		return nil
	}, append(retryOpts, retry.OnRetry(func(n uint, err error) {
		print(fmt.Sprintf("retry %d: %v", n, err))
	}))...)

	if err != nil {
		return fmt.Errorf("Downauthlib: %w", errors.Join(err, FileDownLoadFail))
	}
	return nil
}

var authlibUrls = []string{
	"https://authlib-injector.yushi.moe/artifact/latest.json",
	"https://bmclapi2.bangbang93.com/mirrors/authlib-injector/artifact/latest.json",
	"https://download.mcbbs.net/mirrors/authlib-injector/artifact/latest.json",
}

func randAuthlibUrls(url string) string {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	u := ""
	for {
		i := r.Intn(len(authlibUrls))
		u = authlibUrls[i]
		if url != u {
			break
		}
	}
	return u
}

func getAuthlibLatestUrl(cxt context.Context, url string) (downloadURL string, hash string, err error) {
	reps, _, err := Aget(cxt, url)
	if reps != nil {
		defer reps.Body.Close()
	}
	if err != nil {
		return "", "", fmt.Errorf("getAuthlibLatestUrl: %w", err)
	}
	b, err := io.ReadAll(reps.Body)
	if err != nil {
		return "", "", fmt.Errorf("getAuthlibLatestUrl: %w", err)
	}
	adata := authlibData{}
	err = json.Unmarshal(b, &adata)
	if err != nil {
		return "", "", fmt.Errorf("getAuthlibLatestUrl: %w", err)
	}
	return adata.DownloadURL, adata.Checksums.Sha256, nil
}

type authlibData struct {
	BuildNumber int                  `json:"build_number"`
	Checksums   authlibDataChecksums `json:"checksums"`
	DownloadURL string               `json:"download_url"`
	Version     string               `json:"version"`
}

type authlibDataChecksums struct {
	Sha256 string `json:"sha256"`
}
