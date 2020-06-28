package download

import (
	"archive/zip"
	"errors"
	"fmt"
	"gomclauncher/launcher"
	"io"
	"os"
	"runtime"
	"strings"
	"sync"
)

func (l Libraries) Unzip(typee string, i int) error {
	ch := make(chan bool, i)
	e := make(chan error, len(l.librarie.Libraries))
	done := make(chan bool, len(l.librarie.Libraries))
	natives := make([]string, 0)
	m := sync.Mutex{}
	go func() {
		for _, v := range l.librarie.Libraries {
			v := v
			path, sha1, url := swichnatives(v)
			path = `.minecraft/libraries/` + path
			if url == "" {
				done <- true
				continue
			}
			m.Lock()
			natives = append(natives, path)
			m.Unlock()
			if ifallow(v) && !ver(path, sha1) {
				if path != "" {
					ch <- true
					go func() {
						defer func() {
							<-ch
							done <- true
						}()
						for i := 0; i < 4; i++ {
							if i == 3 {
								e <- errors.New("file download fail")
								break
							}
							err := get(source(url, typee), path)
							if err != nil {
								if err.Error() == "proxy err" {
									e <- errors.New("proxy err")
									break
								}
								fmt.Println("似乎是网络问题，重试", err)
								continue
							}
							if !ver(path, sha1) {
								fmt.Println("文件效验失败，重新下载", url)
								continue
							}
							break
						}
					}()
				}
			} else {
				done <- true
			}

		}
	}()
	n := 0
	for {
		select {
		case <-done:
			n++
			if n == len(l.librarie.Libraries) {
				m.Lock()
				defer m.Unlock()
				return l.unzipnative(natives)
			}
		case err := <-e:
			return err
		}
	}
}

func (l Libraries) unzipnative(n []string) error {
	e := make(chan error, len(n))
	done := make(chan bool, len(n))
	p := `.minecraft/versions/` + l.librarie.ID + `/natives/`
	err := os.MkdirAll(p, 777)
	if err != nil {
		return err
	}
	go func() {
		for _, v := range n {
			v := v
			go func() {
				err := DeCompress(v, p)
				if err != nil {
					e <- err
				}
				done <- true
			}()
		}
	}()
	i := 0
	for {
		select {
		case <-done:
			i++
			if i == len(n) {
				return nil
			}
		case err := <-e:
			return err
		}
	}
}

func ifallow(l launcher.LibraryX115) bool {
	if l.Rules != nil {
		var allow bool
		for _, r := range l.Rules {
			if r.Action == "disallow" && osbool(r.Os.Name) {
				return false
			}
			if r.Action == "allow" && (r.Os.Name == "" || osbool(r.Os.Name)) {
				allow = true
			}
		}
		return allow
	}

	return true

}

func osbool(os string) bool {
	GOOS := runtime.GOOS
	if GOOS == "darwin" {
		GOOS = "osx"
	}
	return os == GOOS
}

func swichnatives(l launcher.LibraryX115) (path, sha1, url string) {
	Os := runtime.GOOS
	switch Os {
	case "windows":
		path = l.Downloads.Classifiers.NativesWindows.Path
		sha1 = l.Downloads.Classifiers.NativesLinux.Sha1
		url = l.Downloads.Classifiers.NativesLinux.URL
	case "darwin":
		if l.Downloads.Classifiers.NativesOsx.Path != "" {
			path = l.Downloads.Classifiers.NativesOsx.Path
			sha1 = l.Downloads.Classifiers.NativesOsx.Sha1
			url = l.Downloads.Classifiers.NativesOsx.URL
		} else {
			path = l.Downloads.Classifiers.NativesMacos.Path
			sha1 = l.Downloads.Classifiers.NativesMacos.Sha1
			url = l.Downloads.Classifiers.NativesMacos.URL
		}
	case "linux":
		path = l.Downloads.Classifiers.NativesLinux.Path
		sha1 = l.Downloads.Classifiers.NativesLinux.Sha1
		url = l.Downloads.Classifiers.NativesLinux.URL
	default:
		panic("???")
	}
	return
}

func DeCompress(zipFile, dest string) error {
	reader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer reader.Close()
	for _, file := range reader.File {
		if !strings.Contains(strings.ToTitle(file.Name), strings.ToTitle("META-INF")) && (strings.HasSuffix(strings.ToTitle(file.Name), strings.ToTitle("dll")) || strings.HasSuffix(strings.ToTitle(file.Name), strings.ToTitle("dylib")) || strings.HasSuffix(strings.ToTitle(file.Name), strings.ToTitle("so"))) {
			rc, err := file.Open()
			if err != nil {
				return err
			}
			defer rc.Close()
			filename := dest + file.Name
			if err != nil {
				return err
			}
			w, err := os.Create(filename)
			if err != nil {
				return err
			}
			defer w.Close()
			_, err = io.Copy(w, rc)
			if err != nil {
				return err
			}
			w.Close()
			rc.Close()
		}
	}
	return nil
}
