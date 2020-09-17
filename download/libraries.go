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
	"strings"
	"time"

	"github.com/xmdhs/gomclauncher/auth"
	"github.com/xmdhs/gomclauncher/launcher"
)

type Libraries struct {
	librarie   launcher.LauncherjsonX115
	assetIndex assets
	typee      string
}

func Newlibraries(b []byte, typee string) (Libraries, error) {
	mod := launcher.Modsjson{}
	var url, id string
	l := launcher.LauncherjsonX115{}
	err := json.Unmarshal(b, &mod)
	if err != nil {
		return Libraries{}, fmt.Errorf("Newlibraries: %w", err)
	}
	if mod.InheritsFrom != "" {
		b, err := ioutil.ReadFile(launcher.Minecraft + `/versions/` + mod.InheritsFrom + "/" + mod.InheritsFrom + ".json")
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
	path := launcher.Minecraft + "/assets/indexes/" + id + ".json"
	if !ver(path, l.AssetIndex.Sha1) {
		err := assetsjson(url, path, typee, l.AssetIndex.Sha1)
		if err != nil {
			return Libraries{}, fmt.Errorf("Newlibraries: %w", err)
		}
	}
	bb, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	a := assets{}
	err = json.Unmarshal(bb, &a)
	if err != nil {
		return Libraries{}, fmt.Errorf("Newlibraries: %w", err)
	}
	return Libraries{
		librarie:   l,
		assetIndex: a,
		typee:      typee,
	}, nil
}

type assets struct {
	Objects map[string]asset `json:"objects"`
}

type asset struct {
	Hash string `json:"hash"`
}

func get(u, path string) error {
	reps, timer, err := Aget(u)
	if reps != nil {
		defer reps.Body.Close()
	}
	if err != nil {
		return fmt.Errorf("get: %w", err)
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
	bw := bufio.NewWriter(f)
	defer f.Close()
	if err != nil {
		return fmt.Errorf("get: %w", err)
	}
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
	}
	return url
}

func Aget(aurl string) (*http.Response, *time.Timer, error) {
	ctx, cancel := context.WithCancel(context.TODO())
	rep, err := http.NewRequestWithContext(ctx, "GET", aurl, nil)
	timer := time.AfterFunc(5*time.Second, func() {
		cancel()
	})
	if err != nil {
		return nil, nil, fmt.Errorf("Aget: %w", err)
	}
	rep.Header.Set("Accept", "*/*")
	rep.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")
	reps, err := auth.HttpClient.Do(rep)
	if err != nil {
		return reps, nil, fmt.Errorf("Aget: %w", err)
	}
	return reps, timer, nil
}

func assetsjson(url, path, typee, sha1 string) error {
	var err error
	f := auto(typee)
	for i := 0; i < 4; i++ {
		if i == 3 {
			return err
		}
		err = get(source(url, f), path)
		if err != nil {
			f = fail(f)
			fmt.Println("下载失败，重试", fmt.Errorf("assetsjson: %w", err), url)
			continue
		}
		if !ver(path, sha1) {
			f = fail(f)
			fmt.Println("文件效验失败，重试", url)
			continue
		}
		break
	}
	return nil
}
