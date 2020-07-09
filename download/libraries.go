package download

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/xmdhs/gomclauncher/auth"
	"github.com/xmdhs/gomclauncher/launcher"
)

type Libraries struct {
	librarie   launcher.LauncherjsonX115
	assetIndex assets
}

func Newlibraries(b []byte) (Libraries, error) {
	mod := launcher.Modsjson{}
	var url, id string
	l := launcher.LauncherjsonX115{}
	err := json.Unmarshal(b, &mod)
	if err != nil {
		return Libraries{}, err
	}
	if mod.InheritsFrom != "" {
		b, err := ioutil.ReadFile(launcher.Minecraft + `/versions/` + mod.InheritsFrom + "/" + mod.InheritsFrom + ".json")
		if err != nil {
			return Libraries{}, err
		}
		json.Unmarshal(b, &l)
		modlibraries2(mod.Libraries, &l)
		l.ID = mod.ID
	} else {
		json.Unmarshal(b, &l)
	}
	url = l.AssetIndex.URL
	id = l.AssetIndex.ID
	path := launcher.Minecraft + "/assets/indexes/" + id + ".json"
	_, err = os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			err := get(url, path)
			if err != nil {
				return Libraries{}, err
			}
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
	return Libraries{
		librarie:   l,
		assetIndex: a,
	}, nil
}

type assets struct {
	Objects map[string]asset `json:"objects"`
}

type asset struct {
	Hash string `json:"hash"`
}

func get(u, path string) error {
	reps, err := Aget(u)
	if reps != nil {
		defer reps.Body.Close()
	}
	if err != nil {
		return err
	}
	_, err = os.Stat(path)
	if err != nil {
		s := strings.Split(path, "/")
		ss := strings.ReplaceAll(path, s[len(s)-1], "")
		err := os.MkdirAll(ss, 0777)
		if err != nil {
			panic(err)
		}
	}
	f, err := os.Create(path)
	defer f.Close()
	if err != nil {
		return err
	}
	_, err = io.Copy(f, reps.Body)
	if err != nil {
		return err
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

func source(url, types string) string {
	switch types {
	case "bmclapi":
		url = strings.ReplaceAll(url, `launchermeta.mojang.com`, `bmclapi2.bangbang93.com`)
		url = strings.ReplaceAll(url, `launcher.mojang.com`, `bmclapi2.bangbang93.com`)
		url = strings.ReplaceAll(url, `resources.download.minecraft.net`, `bmclapi2.bangbang93.com/assets`)
		url = strings.ReplaceAll(url, `libraries.minecraft.net`, `bmclapi2.bangbang93.com/maven`)
		url = strings.ReplaceAll(url, `files.minecraftforge.net/maven`, `bmclapi2.bangbang93.com/maven`)
	case "mcbbs":
		url = strings.ReplaceAll(url, `launchermeta.mojang.com`, `download.mcbbs.net`)
		url = strings.ReplaceAll(url, `launcher.mojang.com`, `download.mcbbs.net`)
		url = strings.ReplaceAll(url, `resources.download.minecraft.net`, `download.mcbbs.net/assets`)
		url = strings.ReplaceAll(url, `libraries.minecraft.net`, `download.mcbbs.net/maven`)
		url = strings.ReplaceAll(url, `files.minecraftforge.net/maven`, `download.mcbbs.net/maven`)
	case "tss":
		url = strings.ReplaceAll(url, `launcher.mojang.com`, `mc.mirrors.tmysam.top`)
		url = strings.ReplaceAll(url, `resources.download.minecraft.net`, `mcres.mirrors.tmysam.top`)
		url = strings.ReplaceAll(url, `libraries.minecraft.net`, `mclib.mirrors.tmysam.top`)
	}
	if strings.Contains(types, "|") {
		sou := rand.NewSource(time.Now().UnixNano())
		r := rand.New(sou)
		i := r.Intn(len(ttypes))
		return source(url, ttypes[i])
	}
	return url
}

func Aget(aurl string) (*http.Response, error) {
	c := http.Client{
		Transport: auth.Transport,
		Timeout:   10 * time.Second,
	}
	rep, err := http.NewRequest("GET", aurl, nil)
	if err != nil {
		return nil, err
	}
	rep.Header.Set("Accept", "*/*")
	rep.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")
	reps, err := c.Do(rep)
	if err != nil {
		return reps, err
	}
	return reps, nil
}
