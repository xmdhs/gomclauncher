package download

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/xmdhs/gomclauncher/launcher"
)

func (l Libraries) Downassets(typee string, i int, c chan int) error {
	e, done, ch := creatch(len(l.assetIndex.Objects), i)
	go func() {
		for _, v := range l.assetIndex.Objects {
			v := v
			ok := ver(launcher.Minecraft+`/assets/objects/`+v.Hash[:2]+`/`+v.Hash, v.Hash)
			if !ok {
				d := downinfo{
					typee: typee,
					url:   `https://resources.download.minecraft.net/` + v.Hash[:2] + `/` + v.Hash,
					path:  launcher.Minecraft + `/assets/objects/` + v.Hash[:2] + `/` + v.Hash,
					e:     e,
					Sha1:  v.Hash,
					done:  done,
				}
				ch <- true
				go d.down()
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
			c <- len(l.assetIndex.Objects) - n
			if n == len(l.assetIndex.Objects) {
				close(c)
				return nil
			}
		case err := <-e:
			return err
		}
	}
}

func ver(path, hash string) bool {
	if hash != "" {
		file, err := os.Open(path)
		defer file.Close()
		if err != nil {
			return false
		}
		m := sha1.New()
		if _, err := io.Copy(m, file); err != nil {
			return false
		}
		h := hex.EncodeToString(m.Sum(nil))
		if h == hash {
			return true
		}
		return false
	}
	_, err := os.Stat(path)
	if err != nil {
		return false
	}
	return true

}

func (l Libraries) Downlibrarie(typee string, i int, c chan int) error {
	e, done, ch := creatch(len(l.librarie.Libraries), i)
	go func() {
		for _, v := range l.librarie.Libraries {
			v := v
			path := launcher.Minecraft + `/libraries/` + v.Downloads.Artifact.Path
			if v.Downloads.Artifact.URL == "" {
				done <- true
				continue
			}
			if !ver(path, v.Downloads.Artifact.Sha1) {
				d := downinfo{
					typee: typee,
					url:   v.Downloads.Artifact.URL,
					path:  path,
					e:     e,
					Sha1:  v.Downloads.Artifact.Sha1,
					done:  done,
				}
				ch <- true
				go d.down()
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
			c <- len(l.librarie.Libraries) - n
			if n == len(l.librarie.Libraries) {
				close(c)
				return nil
			}
		case err := <-e:
			return err
		}
	}
}

func (l Libraries) Downjar(typee, version string) error {
	path := launcher.Minecraft + `/versions/` + version + "/" + version + ".jar"
	if ver(path, l.librarie.Downloads.Client.Sha1) {
		return nil
	}
	t := typee
	for i := 0; i < 4; i++ {
		if i == 3 {
			return errors.New("file download fail")
		}
		t = auto(t)
		err := get(source(l.librarie.Downloads.Client.URL, t), path)
		if err != nil {
			fmt.Println("似乎是网络问题，重试", source(l.librarie.Downloads.Client.URL, t), err)
			t = fail(t)
			continue
		}
		if !ver(path, l.librarie.Downloads.Client.Sha1) {
			fmt.Println("文件效验失败，重新下载", source(l.librarie.Downloads.Client.URL, t), err)
			t = fail(t)
			continue
		}
		break
	}
	return nil
}

type downinfo struct {
	typee string
	url   string
	path  string
	e     chan error
	Sha1  string
	done  chan bool
	ch    chan bool
}

func (d downinfo) down() {
	f := d.typee
	for i := 0; i < 7; i++ {
		if i == 6 {
			d.e <- errors.New("file download fail")
			break
		}
		f = auto(f)
		err := get(source(d.url, f), d.path)
		if err != nil {
			fmt.Println("似乎是网络问题，重试", source(d.url, f), err)
			f = fail(f)
			continue
		}
		if !ver(d.path, d.Sha1) {
			fmt.Println("文件效验失败，重新下载", source(d.url, f))
			f = fail(f)
			continue
		}
		d.done <- true
		<-d.ch
		add(f)
		break
	}

}

func creatch(a, i int) (e chan error, done, ch chan bool) {
	e = make(chan error, a)
	done = make(chan bool, a)
	ch = make(chan bool, i)
	return
}
