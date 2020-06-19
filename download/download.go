package download

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"gomclauncher/launcher"
	"io"
	"net/url"
	"os"
)

func (l libraries) Downassets(typee string, i int) error {
	ch := make(chan bool, i)
	e := make(chan error, 10)
	defer close(ch)
	for _, v := range l.assetIndex.Objects {
		v := v
		ok := ver(`./.minecraft/assets/objects/`+v.Hash[:2]+`/`+v.Hash, v.Hash)
		if ok {
			continue
		}
		select {
		case ch <- true:
			go func() {
				for i := 0; i < 4; i++ {
					if i == 3 {
						e <- errors.New("file download fail")
					}
					err := get(source(`http://resources.download.minecraft.net`+v.Hash[:2]+`/`+v.Hash, typee), `./.minecraft/assets/objects/`+v.Hash[:2]+`/`+v.Hash)
					if err != nil {
						_, ok := err.(*url.Error)
						if ok {
							e <- errors.New("proxy err")
						}
						continue
					}
					ok := ver(`.minecraft/assets/objects/`+v.Hash[:2]+`/`+v.Hash, v.Hash)
					if !ok {
						continue
					}
					<-ch
					break
				}
			}()
		case err := <-e:
			return err
		}
	}
	return nil
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

func (l libraries) Downlibrarie(typee string, i int) error {
	ch := make(chan bool, i)
	e := make(chan error, 10)
	defer close(ch)
	for _, v := range l.librarie.Patches[0].Libraries {
		v := v
		path := `./.minecraft/libraries/` + v.Downloads.Artifact.Path
		if !librariesvar(v, path) {
			select {
			case ch <- true:
				go func() {
					for i := 0; i < 4; i++ {
						if i == 3 {
							e <- errors.New("file download fail")
						}
						err := get(source(v.Downloads.Artifact.URL, typee), path)
						if err != nil {
							_, ok := err.(*url.Error)
							if ok {
								e <- errors.New("proxy err")
							}
							continue
						}
						if !librariesvar(v, path) {
							continue
						}
						<-ch
						break
					}
				}()
			case err := <-e:
				return err
			}

		}
	}
	return nil
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
	path := `./.minecraft/versions` + l.librarie.ID + "/" + l.librarie.ID + ".jar"
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
			_, ok := err.(*url.Error)
			if ok {
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
