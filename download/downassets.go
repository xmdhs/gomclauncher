package download

import (
	"errors"
	"net/url"
)

func Downassets(l libraries, typee string) error {
	for _, v := range l.assetIndex.Objects {
		for i := 0; i < 3; i++ {
			err := get(source(`http://resources.download.minecraft.net`+v.Hash[:2]+`/`+v.Hash, typee), `.minecraft/assets/objects/`+v.Hash[:2]+`/`+v.Hash)
			if err != nil {
				_, ok := err.(*url.Error)
				if ok {
					return errors.New("proxy err")
				}
				continue
			}
			break
		}
	}
	return nil
}
