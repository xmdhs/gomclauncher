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
	ch := make(chan bool, i)
	e := make(chan error, len(l.assetIndex.Objects))
	done := make(chan bool, len(l.assetIndex.Objects))
	go func() {
		for _, v := range l.assetIndex.Objects {
			v := v
			ok := ver(launcher.Minecraft+`/assets/objects/`+v.Hash[:2]+`/`+v.Hash, v.Hash)
			if !ok {
				ch <- true
				go func() {
					defer func() {
						<-ch
						done <- true
					}()
					for i := 0; i < 6; i++ {
						if i == 5 {
							e <- errors.New("file download fail")
							break
						}
						err := get(source(`https://resources.download.minecraft.net/`+v.Hash[:2]+`/`+v.Hash, typee), launcher.Minecraft+`/assets/objects/`+v.Hash[:2]+`/`+v.Hash)
						if err != nil {
							if err.Error() == "proxy err" {
								e <- errors.New("proxy err")
								break
							}
							fmt.Println("似乎是网络问题，重试", err)
							continue
						}
						ok := ver(launcher.Minecraft+`/assets/objects/`+v.Hash[:2]+`/`+v.Hash, v.Hash)
						if !ok {
							fmt.Println("文件效验失败，重新下载", v.Hash)
							continue
						}
						break
					}
				}()
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

func (l Libraries) Downlibrarie(typee string, i int, c chan int) error {
	ch := make(chan bool, i)
	e := make(chan error, len(l.librarie.Libraries))
	done := make(chan bool, len(l.librarie.Libraries))
	go func() {
		for _, v := range l.librarie.Libraries {
			v := v
			path := launcher.Minecraft + `/libraries/` + v.Downloads.Artifact.Path
			if v.Downloads.Artifact.URL == "" {
				done <- true
				continue
			}
			if !librariesvar(v, path) {
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
						err := get(source(v.Downloads.Artifact.URL, typee), path)
						if err != nil {
							if err.Error() == "proxy err" {
								e <- errors.New("proxy err")
								break
							}
							fmt.Println("似乎是网络问题，重试", err)
							continue
						}
						if !librariesvar(v, path) {
							fmt.Println("文件效验失败，重新下载", path)
							continue
						}
						break
					}
				}()
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

func librariesvar(v launcher.LibraryX115, path string) bool {
	if v.Downloads.Artifact.Sha1 != "" {
		return ver(path, v.Downloads.Artifact.Sha1)
	}
	_, err := os.Stat(path)
	if err != nil {
		return false
	}
	return true

}

func (l Libraries) Downjar(typee string) error {
	path := launcher.Minecraft + `/versions/` + l.librarie.ID + "/" + l.librarie.ID + ".jar"
	if ver(path, l.librarie.Downloads.Client.Sha1) {
		return nil
	}
	err := get(l.librarie.Downloads.Client.URL, path)
	if err != nil {
		return err
	}
	for i := 0; i < 4; i++ {
		if i == 3 {
			return errors.New("file download fail")
		}
		err := get(source(l.librarie.Downloads.Client.URL, typee), path)
		if err != nil {
			if err.Error() == "proxy err" {
				return errors.New("proxy err")
			}
			continue
		}
		if !ver(path, l.librarie.Downloads.Client.Sha1) {
			continue
		}
		break
	}
	return nil
}
