package download

import (
	"archive/zip"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
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
			path = l.path + `/libraries/` + path
			if url == "" {
				done <- struct{}{}
				continue
			}
			allow := launcher.Ifallow(v)
			if allow {
				m.Lock()
				natives = append(natives, path)
				m.Unlock()
			}
			if allow && !ver(path, sha1) {
				d := downinfo{
					url:      url,
					path:     path,
					e:        e,
					Sha1:     sha1,
					done:     done,
					ch:       ch,
					cxt:      cxt,
					randurls: l.randurls,
					print:    l.print,
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
	p := l.path + `/versions/` + l.librarie.ID + `/natives/`
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
					err := deCompress(v, p)
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

var needSuffix = map[string]struct{}{
	".dll":   {},
	".so":    {},
	".dylib": {},
}

func deCompress(zipFile, dest string) error {
	reader, err := zip.OpenReader(zipFile)
	if err != nil {
		return fmt.Errorf("DeCompress: %w", err)
	}
	defer reader.Close()
	for _, file := range reader.File {
		if file.FileInfo().IsDir() {
			continue
		}
		err := func() error {
			name := file.FileInfo().Name()
			ext := filepath.Ext(name)
			if _, ok := needSuffix[ext]; ok {
				rc, err := file.Open()
				if err != nil {
					return fmt.Errorf("DeCompress: %w", err)
				}
				defer rc.Close()
				filename := filepath.Join(dest, name)
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
			}
			return nil
		}()
		if err != nil {
			return err
		}
	}
	return nil
}
