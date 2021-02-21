package download

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/xmdhs/gomclauncher/auth"
	"github.com/xmdhs/gomclauncher/lang"
)

func Downauthlib(cxt context.Context) error {
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
			fmt.Println(lang.Lang("authlibdownloadfail"), fmt.Errorf("Downauthlib: %w", err), url)
			url = randurl(url)
			continue
		}
		if !ver(path, auth.Authlibsha1) {
			fmt.Println(lang.Lang("authlibcheckerr"), url)
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
