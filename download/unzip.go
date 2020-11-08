package download

import (
	"archive/zip"
	"context"
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"sync"

	"github.com/xmdhs/gomclauncher/launcher"
)

func (l Libraries) Unzip(i int) error {
	e, done, ch := creatch(len(l.librarie.Libraries), i)
	natives := make([]string, 0)
	m := sync.Mutex{}
	cxt, cancel := context.WithCancel(l.cxt)
	defer cancel()
	go func() {
		for _, v := range l.librarie.Libraries {
			v := v
			path, sha1, url := swichnatives(v)
			path = launcher.Minecraft + `/libraries/` + path
			if url == "" {
				done <- struct{}{}
				continue
			}
			if launcher.Ifallow(v) {
				m.Lock()
				natives = append(natives, path)
				m.Unlock()
			}
			if launcher.Ifallow(v) && !ver(path, sha1) {
				d := downinfo{
					typee:    l.typee,
					url:      url,
					path:     path,
					e:        e,
					Sha1:     sha1,
					done:     done,
					ch:       ch,
					cxt:      cxt,
					randurls: l.randurls,
				}
				select {
				case ch <- struct{}{}:
					go d.down()
				case <-cxt.Done():
					return
				}
			} else {
				select {
				case done <- struct{}{}:
				case <-cxt.Done():
					return
				}
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
			return fmt.Errorf("Unzip: %w", err)
		}
	}
}

func (l Libraries) unzipnative(n []string) error {
	e := make(chan error, len(n))
	done := make(chan bool, len(n))
	cxt, cancel := context.WithCancel(l.cxt)
	defer cancel()
	p := launcher.Minecraft + `/versions/` + l.librarie.ID + `/natives/`
	err := os.MkdirAll(p, 0777)
	if err != nil {
		return fmt.Errorf("unzipnative: %w", err)
	}
	go func() {
		for _, v := range n {
			v := v
			select {
			case <-cxt.Done():
				return
			default:
				go func() {
					err := DeCompress(v, p)
					if err != nil {
						select {
						case e <- err:
						case <-cxt.Done():
						}
						return
					}
					select {
					case done <- true:
					case <-cxt.Done():
					}
				}()
			}
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
			return fmt.Errorf("unzipnative: %w", err)
		}
	}
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
		sha1 = l.Downloads.Classifiers.NativesWindows.Sha1
		url = l.Downloads.Classifiers.NativesWindows.URL
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
		return fmt.Errorf("DeCompress: %w", err)
	}
	defer reader.Close()
	for _, file := range reader.File {
		if !strings.Contains(strings.ToTitle(file.Name), strings.ToTitle("META-INF")) && (strings.HasSuffix(strings.ToTitle(file.Name), strings.ToTitle("dll")) || strings.HasSuffix(strings.ToTitle(file.Name), strings.ToTitle("dylib")) || strings.HasSuffix(strings.ToTitle(file.Name), strings.ToTitle("so"))) {
			rc, err := file.Open()
			if err != nil {
				return fmt.Errorf("DeCompress: %w", err)
			}
			defer rc.Close()
			filename := dest + file.Name
			if err != nil {
				return fmt.Errorf("DeCompress: %w", err)
			}
			w, err := os.Create(filename)
			if err != nil {
				return fmt.Errorf("DeCompress: %w", err)
			}
			defer w.Close()
			_, err = io.Copy(w, rc)
			if err != nil {
				return fmt.Errorf("DeCompress: %w", err)
			}
			w.Close()
			rc.Close()
		}
	}
	return nil
}
