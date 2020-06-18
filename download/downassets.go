package download

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"io"
	"net/url"
	"os"
)

func Downassets(l libraries, typee string) error {
	for _, v := range l.assetIndex.Objects {
		ok := ver(`.minecraft/assets/objects/`+v.Hash[:2]+`/`+v.Hash, v.Hash)
		if ok {
			continue
		}
		for i := 0; i < 3; i++ {
			err := get(source(`http://resources.download.minecraft.net`+v.Hash[:2]+`/`+v.Hash, typee), `.minecraft/assets/objects/`+v.Hash[:2]+`/`+v.Hash)
			if err != nil {
				_, ok := err.(*url.Error)
				if ok {
					return errors.New("proxy err")
				}
				continue
			}
			ok := ver(`.minecraft/assets/objects/`+v.Hash[:2]+`/`+v.Hash, v.Hash)
			if !ok {
				continue
			}
			return errors.New("file download fail")
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
