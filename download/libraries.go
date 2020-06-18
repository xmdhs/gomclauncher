package download

import (
	"encoding/json"
	"gomclauncher/auth"
	"gomclauncher/launcher"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type libraries struct {
	librarie   launcher.LauncherjsonX115
	assetIndex assets
}

func Newlibraries(b []byte) libraries {
	mod := launcher.Modsjson{}
	var url, id string
	l := launcher.LauncherjsonX115{}
	err := json.Unmarshal(b, &mod)
	if err != nil {
		log.Fatalln(err)
	}
	if mod.InheritsFrom != "" {
		b, err := ioutil.ReadFile(`.minecraft/versions/` + mod.InheritsFrom + "/" + mod.InheritsFrom + ".json")
		if err != nil {
			log.Fatal(err)
		}
		json.Unmarshal(b, &l)
		modlibraries2(mod.Libraries, l)
	} else {
		json.Unmarshal(b, &l)
	}
	url = l.Patches[0].AssetIndex.URL
	id = l.Patches[0].AssetIndex.ID
	path := ".minecraft/assets/indexes/" + id + ".json"
	_, err = os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			get(url, path)
		} else {
			panic(err)
		}
	}
	bb, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	a := assets{}
	json.Unmarshal(bb, &a)
	return libraries{
		librarie:   l,
		assetIndex: a,
	}
}

type assets struct {
	Objects map[string]asset `json:"objects"`
}

type asset struct {
	Hash string `json:"hash"`
}

func get(u, path string) error {
	var c http.Client
	if auth.Proxyaddr != "" {
		proxy, err := url.Parse(auth.Proxyaddr)
		if err != nil {
			return err
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
	rep, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return err
	}
	rep.Header.Set("Accept", "*/*")
	rep.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")
	reps, err := c.Do(rep)
	if err != nil {
		return err
	}
	b, err := ioutil.ReadAll(reps.Body)
	defer reps.Body.Close()
	if err != nil {
		return err
	}
	_, err = os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			s := strings.Split(path, "/")
			ss := strings.ReplaceAll(path, s[len(s)-1], "")
			os.MkdirAll(ss, 777)
		} else {
			panic(err)
		}
	}
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 777)
	defer f.Close()
	if err != nil {
		return err
	}
	f.Write(b)
	return nil
}

func modlibraries2(l []launcher.Librarie, Launcherjson launcher.LauncherjsonX115) {
	for _, v := range l {
		if v.Downloads.Artifact.URL != "" {
			Librarie := launcher.LibraryX115{}
			Librarie.Downloads.Artifact.Path = v.Downloads.Artifact.Path
			Librarie.Downloads.Artifact.URL = v.Downloads.Artifact.URL
			Librarie.Downloads.Artifact.Sha1 = v.Downloads.Artifact.Sha1
			Launcherjson.Patches[0].Libraries = append(Launcherjson.Patches[0].Libraries, Librarie)
		} else {
			Librarie := launcher.LibraryX115{}
			s := launcher.Name2path(v.Name)
			path := strings.ReplaceAll(s[0], ".", "/") + "/" + s[1] + "/" + s[2] + "/" + s[1] + "-" + s[2] + ".jar"
			Librarie.Downloads.Artifact.Path = path
			Librarie.Downloads.Artifact.URL = v.Url + path
			Launcherjson.Patches[0].Libraries = append(Launcherjson.Patches[0].Libraries, Librarie)
		}
	}
}

func source(url, types string) string {
	switch types {
	default:
		return url
	}
}
