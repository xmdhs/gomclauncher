package download

import (
	"errors"
	"fmt"
	"math/rand"
	"runtime"
	"time"

	"github.com/xmdhs/gomclauncher/auth"
)

var authliburls = []string{"https://authlib-injector.yushi.moe/artifact/27/authlib-injector-1.1.27-5ef5f8e.jar", "https://download.mcbbs.net/mirrors/authlib-injector/artifact/27/authlib-injector-1.1.27-5ef5f8e.jar"}
var minecraft string

const authlibsha1 = "EBE6CEFF486816E060356B9657A9263616AFB8C1"
const Authlibversion = "1.1.27-5ef5f8e"

func Downauthlib() error {
	if runtime.GOOS == "windows" {
		minecraft = `.minecraft`
	}
	url := randurl("")
	var path = minecraft + `/libraries/` + `moe/yushi/authlibinjector/` + "authlib-injector/" + Authlibversion + "/authlib-injector-" + Authlibversion + ".jar"
	if ver(path, authlibsha1) {
		return nil
	}
	for i := 0; i < 5; i++ {
		if i == 3 {
			return errors.New("download fail")
		}
		err := get(url, path)
		if err != nil {
			fmt.Println("authlib 下载失败，重试", err)
			url = randurl(url)
			continue
		}
		if !ver(path, authlibsha1) {
			fmt.Println("authlib 效验出错，重试")
			url = randurl(url)
			continue
		}
		break
	}
	auth.Authlibpath = path
	return nil
}

func randurl(aurl string) string {
	var url string
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for {
		i := r.Intn(len(authliburls) - 1)
		url = authliburls[i]
		if url != aurl {
			break
		}
	}
	return url
}
