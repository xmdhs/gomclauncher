package main

import (
	"bufio"
	"flag"
	"fmt"
	"gomclauncher/auth"
	down "gomclauncher/download"
	"gomclauncher/launcher"
	"log"
	"os"
)

func main() {
	if *runlist {
		arunlist()
	}

	if *online {

	}
}

var (
	g         launcher.Gameinfo
	online    *bool
	username  *string
	passworld *string
	email     *string
	download  *string
	verlist   *bool
	run       *string
	runlist   *bool
	runram    *string
	runflag   *string
	proxy     *string
	atype     *string
)

func init() {
	str, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	g.Minecraftpath = str
	online = flag.Bool("online", false, `是否启用正版登录，默认关闭`)
	username = flag.String("username", "", `用户名，只需第一次设置，第二次无需使用此参数。`)
	email = flag.String("email", "", `正版帐号邮箱`)
	passworld = flag.String("passworld", "", `正版帐号密码，只需第一次设置，第二次无需使用此参数。`)
	download = flag.String("downver", "", "尝试下载的版本")
	verlist = flag.Bool("verlist", false, "显示所有可下载的版本")
	run = flag.String("run", "", `尝试启动的版本`)
	runlist = flag.Bool("runlist", false, "显示所有可启动的版本")
	runram = flag.String("ram", "2048", `分配启动游戏的内存大小(mb)`)
	runflag = flag.String("flag", "", "自定的启动参数")
	proxy = flag.String(`proxy`, "", `设置下载用的代理(http)`)
	atype = flag.String("type", "", `设置下载源。目前只能使用官方下载源`)

	flag.Parse()
}

func arunlist() {
	l, err := down.Getversionlist(*atype)
	if err != nil {
		if err.Error() == "proxy err" {
			log.Fatalln("设置的代理有误")
		} else {
			log.Fatalln("可能是网络问题，可再次尝试")
		}
		log.Fatal(err)
	}
	w := bufio.NewScanner(os.Stdin)
	fmt.Print("输入想查看的类型，可选")
	m := make(map[string]bool)
	for _, v := range l.Versions {
		m[v.Type] = true
	}
	for k, _ := range m {
		fmt.Print(k)
		fmt.Print(", ")
	}
	w.Scan()
	t := w.Text()
	for _, v := range l.Versions {
		if v.Type == t {
			fmt.Println(v.ID)
		}
	}
	w.Scan()
}

func runmc() {
	if *email == "" {
		log.Println("请设置邮箱")
		log.Fatal("比如 -email xxx@xxx.xx")
	}
	err := aconfig.setonline(*email, *passworld)
	if err != nil {
		if err.Error() == "have" {
			a := auth.Auth{
				AccessToken: aconfig.EmailAccessToken[*email],
				ClientToken: aconfig.ClientToken,
			}
			err := auth.Refresh(&a)
			if err != nil {
				if err.Error() == "not ok" {
					log.Fatalln("请尝试重新登录")
				} else {
					panic(err)
				}
			}

		} else if err.Error() == "not ok" {
			log.Fatalln("账户名或密码错误")
		} else {
			log.Fatalln(err)
		}
	} else {
		g.AccessToken = aconfig.AccessToken
		g.Name = aconfig.Name
		g.UUID = aconfig.UUID
	}

}
