package download

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/xmdhs/gomclauncher/internal"
	"github.com/xmdhs/gomclauncher/launcher"
	"golang.org/x/sync/errgroup"
)

func (l Libraries) Unzip(i int) error {
	natives := make([]string, 0)

	g, ctx := errgroup.WithContext(l.cxt)
	g.SetLimit(i)

	for _, v := range l.librarie.Libraries {
		v := v
		path, sha1, url := swichnatives(v)
		if url == "" {
			continue
		}
		path, err := internal.SafePathJoin(l.path, `/libraries/`, path)
		if err != nil {
			return fmt.Errorf("Unzip: %w", err)
		}
		allow := launcher.Ifallow(v)
		if allow {
			natives = append(natives, path)
		}
		g.Go(func() error {
			d := downinfo{
				url:      url,
				path:     path,
				Sha1:     sha1,
				randurls: l.randurls,
				print:    l.print,
			}
			return d.down(ctx)
		})
	}
	err := g.Wait()
	if err != nil {
		return fmt.Errorf("Unzip: %w", err)
	}
	return l.unzipnative(natives)
}

func (l Libraries) unzipnative(n []string) error {
	if len(n) == 0 {
		return nil
	}
	p := l.path + `/versions/` + l.librarie.ID + `/natives/`
	err := os.RemoveAll(p)
	if err != nil {
		return fmt.Errorf("unzipnative: %w", err)
	}
	err = os.MkdirAll(p, 0777)
	if err != nil {
		return fmt.Errorf("unzipnative: %w", err)
	}

	g, _ := errgroup.WithContext(l.cxt)
	g.SetLimit(8)

	for _, v := range n {
		v := v
		g.Go(func() error {
			return deCompress(v, p)
		})
	}
	return g.Wait()
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
			ext = strings.ToLower(ext)
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
