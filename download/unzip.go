package download

import (
	"archive/zip"
	"errors"
	"gomclauncher/launcher"
	"io"
	"os"
	"runtime"
	"strings"
)

func (l Libraries) Unzip(typee string, i int) error {
	ch := make(chan bool, i)
	e := make(chan error, len(l.librarie.Libraries))
	done := make(chan bool, len(l.librarie.Libraries))
	natives := make([]string, 0)
	go func() {
		for _, v := range l.librarie.Libraries {
			v := v
			h := swichnatives(v)
			if ifallow(v) && !ver(h[0], h[1]) {
				if h[0] != "" {
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
							err := get(source(h[2], typee), h[0])
							if err != nil {
								if err.Error() == "proxy err" {
									e <- errors.New("proxy err")
									break
								}
								continue
							}
							if !ver(h[0], h[1]) {
								continue
							}
							natives = append(natives, h[0])
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
	go func() {
		for _, v := range n {
			v := v
			go func() {
				err := DeCompress(v, `.minecraft/versions/`+l.librarie.ID+`natives`)
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

func swichnatives(l launcher.LibraryX115) [3]string {
	os := runtime.GOOS
	h := [3]string{}
	switch os {
	case "windows":
		h[0] = l.Downloads.Classifiers.NativesWindows.Path
		h[1] = l.Downloads.Classifiers.NativesLinux.Sha1
		h[2] = l.Downloads.Classifiers.NativesLinux.URL
	case "darwin":
		if l.Downloads.Classifiers.NativesOsx.Path != "" {
			h[0] = l.Downloads.Classifiers.NativesOsx.Path
			h[1] = l.Downloads.Classifiers.NativesOsx.Sha1
			h[2] = l.Downloads.Classifiers.NativesOsx.URL
		} else {
			h[0] = l.Downloads.Classifiers.NativesMacos.Path
			h[1] = l.Downloads.Classifiers.NativesMacos.Sha1
			h[2] = l.Downloads.Classifiers.NativesMacos.URL
		}
	case "linux":
		h[0] = l.Downloads.Classifiers.NativesLinux.Path
		h[1] = l.Downloads.Classifiers.NativesLinux.Sha1
		h[2] = l.Downloads.Classifiers.NativesLinux.URL
	}
	return h
}

func DeCompress(zipFile, dest string) error {
	reader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer reader.Close()
	for _, file := range reader.File {
		if strings.Contains(strings.ToTitle(file.Name), strings.ToTitle("META-INF")) && (strings.HasSuffix(strings.ToTitle(file.Name), strings.ToTitle("dll")) || strings.HasSuffix(strings.ToTitle(file.Name), strings.ToTitle("dylib")) || strings.HasSuffix(strings.ToTitle(file.Name), strings.ToTitle("so"))) {
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
