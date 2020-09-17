package download

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"

	"github.com/xmdhs/gomclauncher/auth"
)

func Downauthlib() error {
	minecraft := "minecraft"
	if runtime.GOOS == "windows" {
		minecraft = `.minecraft`
	}
	url := randurl("")
	var path = minecraft + `/libraries/` + `moe/yushi/authlibinjector/` + "authlib-injector/" + auth.Authlibversion + "/authlib-injector-" + auth.Authlibversion + ".jar"
	if ver(path, auth.Authlibsha1) {
		return nil
	}
	for i := 0; i < 5; i++ {
		if i == 3 {
			return FileDownLoadFail
		}
		err := get(url, path)
		if err != nil {
			fmt.Println("authlib 下载失败，重试", fmt.Errorf("Downauthlib: %w", err), url)
			url = randurl(url)
			continue
		}
		if !ver(path, auth.Authlibsha1) {
			fmt.Println("authlib 效验出错，重试", url)
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
