package main

import (
	"bufio"
	"flag"
	"fmt"
	"gomclauncher/auth"
	down "gomclauncher/download"
	"gomclauncher/launcher"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	auth.Proxyaddr = *proxy
	if *runlist {
		arunlist()
	}
	if *download != "" {

	}
	if *online {
		aonline()
	} else {
		g.Name = *username
		g.UUID = uuidgen(*username)
	}
	if *runflag != "" {
		s := strings.Split(*runflag, " ")
		g.Flag = s
	}
	g.RAM = *runram
	if *run != "" {
		arun()
	}
}

func d(ver string) {
	l, err := down.Getversionlist(*atype)
	errr(err)
	err = l.Downjson(*download)
	errr(err)
	b, err := ioutil.ReadFile(".minecraft/versions/" + g.Version + "/" + g.Version + ".json")
	if err != nil {
		panic(err)
	}
	dl, err := down.Newlibraries(b)
	errr(err)
	fmt.Println("正在下载游戏核心")
	err = dl.Downjar(*atype)
	errr(err)
	fmt.Println("完成")
	fmt.Println("正在下载库文件")
	dd(dl, false)
	fmt.Println("完成")
	fmt.Println("正在下载资源文件")
	dd(dl, true)
	fmt.Println("完成")
}

func dd(l down.Libraries, a bool) {
	ch := make(chan int, 5)
	e := make(chan error)
	var err error
	go func() {
		if a {
			err = l.Downassets(*atype, *downint, ch)
		} else {
			err = l.Downlibrarie(*atype, *downint, ch)
		}
		if err != nil {
			e <- err
		}
	}()
b:
	for {
		select {
		case i, ok := <-ch:
			if !ok {
				break b
			}
			fmt.Println(i)
		case err := <-e:
			log.Fatal(err)
			break b
		}
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
	downint   *int
)

func init() {
	str, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	g.Minecraftpath = str
	online = flag.Bool("online", false, `是否启用正版登录，默认关闭`)
	username = flag.String("username", "", `用户名`)
	email = flag.String("email", "", `正版帐号邮箱`)
	passworld = flag.String("passworld", "", `正版帐号密码，只需第一次设置，第二次无需使用此参数。`)
	download = flag.String("downver", "", "尝试下载的版本")
	verlist = flag.Bool("verlist", false, "显示所有可下载的版本")
	downint = flag.Int("int", 32, "下载文件时使用的协程数。默认为 32")
	run = flag.String("run", "", `尝试启动的版本`)
	runlist = flag.Bool("runlist", false, "显示所有可启动的版本")
	runram = flag.String("ram", "2048", `分配启动游戏的内存大小(mb)`)
	runflag = flag.String("flag", "", "自定的启动参数，比如 -XX:+AggressiveOpts -XX:+UseCompressedOops")
	proxy = flag.String(`proxy`, "", `设置下载用的代理(http)`)
	atype = flag.String("type", "", `设置下载源。目前只能使用官方下载源`)
	flag.Parse()
}

func arunlist() {
	l, err := down.Getversionlist(*atype)
	errr(err)
	w := bufio.NewScanner(os.Stdin)
	fmt.Print("输入想查看的类型，可选")
	m := make(map[string]bool)
	for _, v := range l.Versions {
		m[v.Type] = true
	}
	for k := range m {
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

func aonline() {
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
				} else if err.Error() == "proxy err" {
					log.Fatalln("设置的代理有误")
				} else {
					log.Fatalln("可能是网络问题，可再次尝试")
				}
			}
			aconfig.EmailAccessToken[*email] = a.AccessToken
		} else if err.Error() == "not ok" {
			log.Fatalln("账户名或密码错误")
		} else {
			log.Fatalln(err)
		}
	}
	g.Userproperties = aconfig.Userproperties
	g.AccessToken = aconfig.AccessToken
	g.Name = aconfig.Name
	g.UUID = aconfig.UUID

}

func arun() {
	g.Version = *run
	b, err := ioutil.ReadFile(g.Minecraftpath + "/versions/" + g.Version + "/" + g.Version + ".json")
	if err != nil {
		print(err)
	}
	g.Jsonbyte = b
	err = g.Run115()
	if err != nil {
		if err.Error() == "json not exist" {
			log.Fatalln("请先安装对应的原版")
		}
		if err.Error() == "json err" {
			log.Fatalln("json 错误，可尝试到 .minecraft/versions 中删除对应的 json 文件")
		}
	}
}

func errr(err error) {
	if err != nil {
		if err.Error() == "proxy err" {
			log.Fatalln("设置的代理有误")
		} else {
			log.Fatalln("可能是网络问题，可再次尝试")
		}
		log.Fatal(err)
	}
}
