package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strings"

	"github.com/xmdhs/gomclauncher/auth"
	"github.com/xmdhs/gomclauncher/download"
	aflag "github.com/xmdhs/gomclauncher/flag"
	"github.com/xmdhs/gomclauncher/launcher"
)

func main() {
	auth.Proxyaddr = f.Proxy
	if updata {
		check()
	}
	if credit {
		credits()
	}
	if strings.Contains(f.Atype, "|") {
		s := strings.Split(f.Atype, "|")
		download.Fail = true
		download.Types = s
	}
	if f.Verlist {
		f.Arunlist()
	}
	if f.Runlist {
		s := aflag.Find(launcher.Minecraft + `/versions`)
		for _, v := range s {
			if aflag.Test(launcher.Minecraft + `/versions/` + v + `/` + v + ".json") {
				fmt.Println(v)
			}
		}
	}
	if f.Download != "" {
		f.Outmsg = false
		f.D()
	}
	if f.Email != "" {
		f.Aonline()
	} else {
		f.Username = f.Name
		f.UUID = aflag.UUIDgen(f.Username)
		f.AccessToken = f.UUID
	}
	if f.Runflag != "" {
		s := strings.Split(f.Runflag, " ")
		f.Flag = s
	}
	f.Gameinfo.RAM = f.RAM
	if f.Run != "" {
		if f.Name == "" && f.Email == "" {
			fmt.Println("需要设置 username")
			os.Exit(0)
		}
		f.Arun()
	}
}

var f aflag.Flag

var (
	credit bool
	updata bool
)

func init() {
	str, err := os.Getwd()
	if runtime.GOOS == "windows" {
		launcher.Minecraft = `.minecraft`
	}
	str = strings.ReplaceAll(str, `\`, `/`)
	if err != nil {
		panic(err)
	}
	f.Minecraftpath = str + "/" + launcher.Minecraft
	flag.StringVar(&f.Name, "username", "", `用户名`)
	flag.StringVar(&f.Email, "email", "", `正版帐号邮箱，需要正版登录时设置此项，然后无需设置 username`)
	flag.StringVar(&f.Passworld, "passworld", "", `正版帐号密码，只需第一次设置，第二次无需使用此参数。`)
	flag.StringVar(&f.Download, "downver", "", "尝试下载的版本")
	flag.BoolVar(&f.Verlist, "verlist", false, "显示所有可下载的版本")
	flag.IntVar(&f.Downint, "int", 64, "下载文件时使用的协程数。")
	flag.StringVar(&f.Run, "run", "", `尝试启动的版本`)
	flag.BoolVar(&f.Runlist, "runlist", false, "显示所有可启动的版本")
	flag.StringVar(&f.RAM, "ram", "2048", `分配启动游戏的内存大小(mb)`)
	flag.StringVar(&f.Runflag, "flag", "", "自定的启动参数，比如 -XX:+AggressiveOpts -XX:+UseCompressedOops")
	flag.StringVar(&f.Proxy, `proxy`, "", `设置下载用的代理(http)`)
	flag.StringVar(&f.Atype, "type", "", `设置下载源。可选 vanilla bmclapi tss 和 mcbbs，不设置此项则使用将自动的为每一个文件选择下载源。可以使用 "bmclapi|vanilla" 的形式来负载均衡的使用两个下载源。`)
	flag.BoolVar(&f.Independent, "independent", true, "是否开启版本隔离")
	flag.BoolVar(&f.Outmsg, "test", true, "启动游戏前是否效验文件的完整和正确性")
	flag.BoolVar(&credit, "credits", false, "使用项目")
	flag.BoolVar(&updata, "updata", true, "是否检测更新")
	flag.Parse()
}

func credits() {
	fmt.Println(`使用了 bmclapi 作为镜像下载源，地址 https://bmclapidoc.bangbang93.com/`)
	fmt.Println(`使用了 TSS  作为镜像下载源，地址 https://www.mcbbs.net/thread-932755-1-1.html`)
}

type up struct {
	Tag  string `json:"tag_name"`
	Body string `json:"body"`
}

func check() {
	reps, err := download.Aget(`https://api.github.com/repos/xmdhs/gomclauncher/releases/latest`)
	if reps != nil {
		defer reps.Body.Close()
	}
	if err != nil {
		fmt.Println("检测更新失败")
		fmt.Println(err)
		return
	}
	b, err := ioutil.ReadAll(reps.Body)
	if err != nil {
		fmt.Println("检测更新失败")
		fmt.Println(err)
		return
	}
	u := up{}
	err = json.Unmarshal(b, &u)
	if err != nil {
		fmt.Println("检测更新失败")
		fmt.Println(err)
		return
	}
	if u.Tag != "v"+launcher.Launcherversion {
		fmt.Println("检测到更新,新版本为", u.Tag)
		fmt.Println("当前版本为", "v"+launcher.Launcherversion)
		fmt.Println("更新内容：")
		fmt.Println(u.Body)
	}
}
