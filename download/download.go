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
					t := typee
					for i := 0; i < 6; i++ {
						if i == 5 {
							e <- errors.New("file download fail")
							break
						}
						err := get(source(`https://resources.download.minecraft.net/`+v.Hash[:2]+`/`+v.Hash, t), launcher.Minecraft+`/assets/objects/`+v.Hash[:2]+`/`+v.Hash)
						if err != nil {
							fmt.Println("似乎是网络问题，重试", source(`https://resources.download.minecraft.net/`+v.Hash[:2]+`/`+v.Hash, t), err)
							t = fail(t)
							continue
						}
						ok := ver(launcher.Minecraft+`/assets/objects/`+v.Hash[:2]+`/`+v.Hash, v.Hash)
						if !ok {
							fmt.Println("文件效验失败，重新下载", source(`https://resources.download.minecraft.net/`+v.Hash[:2]+`/`+v.Hash, t))
							t = fail(t)
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
					t := typee
					for i := 0; i < 4; i++ {
						if i == 3 {
							e <- errors.New("file download fail")
							break
						}
						err := get(source(v.Downloads.Artifact.URL, t), path)
						if err != nil {
							fmt.Println("似乎是网络问题，重试", source(v.Downloads.Artifact.URL, t), err)
							t = fail(t)
							continue
						}
						if !librariesvar(v, path) {
							fmt.Println("文件效验失败，重新下载", source(v.Downloads.Artifact.URL, t))
							t = fail(t)
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
