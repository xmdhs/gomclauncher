package download

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"gomclauncher/launcher"
	"io"
	"os"
)

func (l libraries) Downassets(typee string, i int, c chan int) error {
	ch := make(chan bool, i)
	e := make(chan error, len(l.assetIndex.Objects))
	done := make(chan bool, len(l.assetIndex.Objects))
	for _, v := range l.assetIndex.Objects {
		v := v
		ok := ver(`.minecraft/assets/objects/`+v.Hash[:2]+`/`+v.Hash, v.Hash)
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
					err := get(source(`http://resources.download.minecraft.net/`+v.Hash[:2]+`/`+v.Hash, typee), `.minecraft/assets/objects/`+v.Hash[:2]+`/`+v.Hash)
					if err != nil {
						if err.Error() == "proxy err" {
							e <- errors.New("proxy err")
							break
						}
						continue
					}
					ok := ver(`.minecraft/assets/objects/`+v.Hash[:2]+`/`+v.Hash, v.Hash)
					if !ok {
						continue
					}
					break
				}
			}()
		} else {
			done <- true
		}
	}
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

func (l libraries) Downlibrarie(typee string, i int, c chan int) error {
	ch := make(chan bool, i)
	e := make(chan error, len(l.librarie.Patches[0].Libraries))
	done := make(chan bool, len(l.librarie.Patches[0].Libraries))
	for _, v := range l.librarie.Patches[0].Libraries {
		v := v
		path := `.minecraft/libraries/` + v.Downloads.Artifact.Path
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
						continue
					}
					if !librariesvar(v, path) {
						continue
					}
					break
				}
			}()
		} else {
			done <- true
		}
	}
	n := 0
	for {
		select {
		case <-done:
			n++
			c <- len(l.librarie.Patches[0].Libraries) - n
			if n == len(l.librarie.Patches[0].Libraries) {
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

func (l libraries) Downjar(typee string, i int) error {
	path := `.minecraft/versions` + l.librarie.ID + "/" + l.librarie.ID + ".jar"
	if ver(path, l.librarie.Patches[0].Downloads.Client.Sha1) {
		return nil
	}
	err := get(l.librarie.Patches[0].Downloads.Client.URL, path)
	if err != nil {
		return err
	}
	for i := 0; i < 4; i++ {
		if i == 3 {
			return errors.New("file download fail")
		}
		err := get(source(l.librarie.Patches[0].Downloads.Client.URL, typee), path)
		if err != nil {
			if err.Error() == "proxy err" {
				return errors.New("proxy err")
			}
			continue
		}
		if !ver(path, l.librarie.Patches[0].Downloads.Client.Sha1) {
			continue
		}
		break
	}
	return nil
}
