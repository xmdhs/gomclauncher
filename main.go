package main

import (
	"flag"
	"fmt"
	"gomclauncher/auth"
	aflag "gomclauncher/flag"
	"os"
	"strings"
)

func main() {
	auth.Proxyaddr = f.Proxy
	if f.Verlist {
		f.Arunlist()
	}
	if f.Runlist {
		s := aflag.Find(`.minecraft/versions`)
		for _, v := range s {
			if aflag.Test(`.minecraft/versions/` + v + `/` + v + ".json") {
				fmt.Println(v)
			}
		}
	}
	if f.Download != "" {
		f.D()
	}
	if f.Online {
		f.Aonline()
	} else {
		if f.Name == "" {
			fmt.Println("需要设置 username")
			os.Exit(0)
		}
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
		f.Arun()
	}
}

var f aflag.Flag

func init() {
	str, err := os.Getwd()
	str = strings.ReplaceAll(str, `\`, `/`)
	if err != nil {
		panic(err)
	}
	f.Minecraftpath = str + `/.minecraft`
	flag.BoolVar(&f.Online, "online", false, `是否启用正版登录，默认关闭`)
	flag.StringVar(&f.Name, "username", "", `用户名`)
	flag.StringVar(&f.Email, "email", "", `正版帐号邮箱`)
	flag.StringVar(&f.Passworld, "passworld", "", `正版帐号密码，只需第一次设置，第二次无需使用此参数。`)
	flag.StringVar(&f.Download, "downver", "", "尝试下载的版本")
	flag.BoolVar(&f.Verlist, "verlist", false, "显示所有可下载的版本")
	flag.IntVar(&f.Downint, "int", 32, "下载文件时使用的协程数。默认为 32")
	flag.StringVar(&f.Run, "run", "", `尝试启动的版本`)
	flag.BoolVar(&f.Runlist, "runlist", false, "显示所有可启动的版本")
	flag.StringVar(&f.RAM, "ram", "2048", `分配启动游戏的内存大小(mb)`)
	flag.StringVar(&f.Aflag, "flag", "", "自定的启动参数，比如 -XX:+AggressiveOpts -XX:+UseCompressedOops")
	flag.StringVar(&f.Proxy, `proxy`, "", `设置下载用的代理(http)`)
	flag.StringVar(&f.Atype, "type", "", `设置下载源。目前只能使用官方下载源`)
	flag.BoolVar(&f.Independent, "independent", false, "是否开启版本隔离")
	flag.BoolVar(&f.Outmsg, "test", true, "启动游戏前是否效验文件的完整和正确性")
	flag.Parse()
}
