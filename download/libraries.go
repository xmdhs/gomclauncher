package download

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/xmdhs/gomclauncher/auth"
	"github.com/xmdhs/gomclauncher/lang"
	"github.com/xmdhs/gomclauncher/launcher"
)

type Libraries struct {
	librarie   launcher.LauncherjsonX115
	assetIndex assets
	typee      string
	cxt        context.Context
	print      func(string)
	path       string
	*randurls
}

func Newlibraries(cxt context.Context, b []byte, typee string, print func(string), apath string) (Libraries, error) {
	mod := launcher.Modsjson{}
	var url, id string
	l := launcher.LauncherjsonX115{}
	err := json.Unmarshal(b, &mod)
	r := newrandurls(typee)
	if err != nil {
		return Libraries{}, fmt.Errorf("Newlibraries: %w", err)
	}
	if mod.InheritsFrom != "" {
		b, err := ioutil.ReadFile(apath + `/versions/` + mod.InheritsFrom + "/" + mod.InheritsFrom + ".json")
		if err != nil {
			return Libraries{}, fmt.Errorf("Newlibraries: %w", err)
		}
		err = json.Unmarshal(b, &l)
		if err != nil {
			return Libraries{}, fmt.Errorf("Newlibraries: %w", err)
		}
		modlibraries2(mod.Libraries, &l)
		l.ID = mod.ID
	} else {
		err = json.Unmarshal(b, &l)
		if err != nil {
			return Libraries{}, fmt.Errorf("Newlibraries: %w", err)
		}
	}
	url = l.AssetIndex.URL
	id = l.AssetIndex.ID
	path := apath + "/assets/indexes/" + id + ".json"
	if !ver(path, l.AssetIndex.Sha1) {
		err := assetsjson(cxt, r, url, path, typee, l.AssetIndex.Sha1, print)
		if err != nil {
			return Libraries{}, fmt.Errorf("Newlibraries: %w", err)
		}
	}
	bb, err := ioutil.ReadFile(path)
	if err != nil {
		return Libraries{}, fmt.Errorf("Newlibraries: %w", err)
	}
	a := assets{}
	err = json.Unmarshal(bb, &a)
	if err != nil {
		return Libraries{}, fmt.Errorf("Newlibraries: %w", err)
	}
	return Libraries{
		print:      print,
		librarie:   l,
		assetIndex: a,
		typee:      typee,
		cxt:        cxt,
		randurls:   r,
		path:       apath,
	}, nil
}

type assets struct {
	Objects map[string]asset `json:"objects"`
}

type asset struct {
	Hash string `json:"hash"`
}

func get(cxt context.Context, u, path string) error {
	reps, timer, err := Aget(cxt, u)
	if reps != nil {
		defer reps.Body.Close()
	}
	if err != nil {
		return fmt.Errorf("get: %w", err)
	}
	_, err = os.Stat(path)

	if err != nil {
		dir, _ := filepath.Split(path)
		err := os.MkdirAll(dir, 0777)
		if err != nil {
			return fmt.Errorf("get: %w", err)
		}
	}
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("get: %w", err)
	}
	defer f.Close()
	bw := bufio.NewWriter(f)
	for {
		timer.Reset(5 * time.Second)
		i, err := io.CopyN(bw, reps.Body, 100000)
		if err != nil && err != io.EOF {
			return fmt.Errorf("get: %w", err)
		}
		if i == 0 {
			break
		}
	}
	err = bw.Flush()
	if err != nil {
		return fmt.Errorf("get: %w", err)
	}
	return nil
}

func modlibraries2(l []launcher.Librarie, Launcherjson *launcher.LauncherjsonX115) {
	for _, v := range l {
		if v.Downloads.Artifact.URL != "" {
			Librarie := launcher.LibraryX115{}
			Librarie.Downloads.Artifact.Path = v.Downloads.Artifact.Path
			Librarie.Downloads.Artifact.URL = v.Downloads.Artifact.URL
			Librarie.Downloads.Artifact.Sha1 = v.Downloads.Artifact.Sha1
			Launcherjson.Libraries = append(Launcherjson.Libraries, Librarie)
		} else {
			Librarie := launcher.LibraryX115{}
			s := launcher.Name2path(v.Name)
			path := strings.ReplaceAll(s[0], ".", "/") + "/" + s[1] + "/" + s[2] + "/" + s[1] + "-" + s[2] + ".jar"
			Librarie.Downloads.Artifact.Path = path
			if v.Url != "" {
				Librarie.Downloads.Artifact.URL = v.Url + path
			} else {
				Librarie.Downloads.Artifact.URL = `https://libraries.minecraft.net/` + path
			}
			Launcherjson.Libraries = append(Launcherjson.Libraries, Librarie)
		}
	}
}

var mirror = map[string]map[string]string{
	"bmclapi": {
		`launchermeta.mojang.com`:          `bmclapi2.bangbang93.com`,
		`launcher.mojang.com`:              `bmclapi2.bangbang93.com`,
		`resources.download.minecraft.net`: `bmclapi2.bangbang93.com/assets`,
		`libraries.minecraft.net`:          `bmclapi2.bangbang93.com/maven`,
		`files.minecraftforge.net/maven`:   `bmclapi2.bangbang93.com/maven`,
		"maven.minecraftforge.net":         "bmclapi2.bangbang93.com/maven",
	},
	"mcbbs": {
		`launchermeta.mojang.com`:          `download.mcbbs.net`,
		`launcher.mojang.com`:              `download.mcbbs.net`,
		`resources.download.minecraft.net`: `download.mcbbs.net/assets`,
		`libraries.minecraft.net`:          `download.mcbbs.net/maven`,
		`files.minecraftforge.net/maven`:   `download.mcbbs.net/maven`,
		"maven.minecraftforge.net":         "bmclapi2.bangbang93.com/maven",
	},
}

func source(url, types string) string {
	m, ok := mirror[types]
	if ok {
		for k, v := range m {
			if strings.Contains(url, k) {
				return strings.Replace(url, k, v, 1)
			}
		}
	}
	return url
}

func Aget(cxt context.Context, aurl string) (*http.Response, *time.Timer, error) {
	ctx, cancel := context.WithCancel(cxt)
	rep, err := http.NewRequestWithContext(ctx, "GET", aurl, nil)
	timer := time.AfterFunc(5*time.Second, func() {
		cancel()
	})
	if err != nil {
		return nil, nil, fmt.Errorf("Aget: %w", err)
	}
	rep.Header.Set("Accept", "*/*")
	rep.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")
	c := http.Client{
		Transport: auth.Transport,
	}
	reps, err := c.Do(rep)
	if err != nil {
		return reps, nil, fmt.Errorf("Aget: %w", err)
	}
	return reps, timer, nil
}

func assetsjson(cxt context.Context, r *randurls, url, path, typee, sha1 string, print func(string)) error {
	var err error
	_, f := r.auto()
	for i := 0; i < 4; i++ {
		if i == 3 {
			return err
		}
		err = get(cxt, source(url, f), path)
		if err != nil {
			f = r.fail(f)
			print(lang.Lang("weberr") + " " + fmt.Errorf("assetsjson: %w", err).Error() + " " + url)
			continue
		}
		if !ver(path, sha1) {
			f = r.fail(f)
			err = ErrFileChecker
			print(lang.Lang("filecheckerr") + " " + url)
			continue
		}
		break
	}
	return nil
}
