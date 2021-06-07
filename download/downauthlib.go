package download

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/xmdhs/gomclauncher/auth"
	"github.com/xmdhs/gomclauncher/lang"
)

func Downauthlib(cxt context.Context, print func(string)) error {
	url := randurl("")
	var path = ".minecraft" + `/libraries/` + `moe/yushi/authlibinjector/` + "authlib-injector/" + auth.Authlibversion + "/authlib-injector-" + auth.Authlibversion + ".jar"
	if ver(path, auth.Authlibsha1) {
		return nil
	}
	for i := 0; i < 5; i++ {
		if i == 3 {
			return FileDownLoadFail
		}
		err := get(cxt, url, path)
		if err != nil {
			print(lang.Lang("authlibdownloadfail") + " " + fmt.Errorf("Downauthlib: %w", err).Error() + " " + url)
			url = randurl(url)
			continue
		}
		if !ver(path, auth.Authlibsha1) {
			print(lang.Lang("authlibcheckerr") + " " + url)
			url = randurl(url)
			continue
		}
		break
	}
	return nil
}

func randurl(aurl string) string {
	var url string
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for {
		i := r.Intn(len(auth.Authliburls))
		url = auth.Authliburls[i]
		if url != aurl {
			break
		}
	}
	return url
}
