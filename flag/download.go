package flag

import (
	"fmt"
	"gomclauncher/download"
	"gomclauncher/launcher"
	"io/ioutil"
	"log"
	"os"
)

type Flag struct {
	launcher.Gameinfo
	Atype       string
	Downint     int
	Online      bool
	Username    string
	Passworld   string
	Email       string
	Download    string
	Verlist     bool
	Run         string
	Runlist     bool
	Runram      string
	Runflag     string
	Proxy       string
	Aflag       string
	Independent bool
	Outmsg      bool
}

func (f Flag) D() {
	l, err := download.Getversionlist(f.Atype)
	errr(err)
	err = l.Downjson(f.Download)
	errr(err)
	b, err := ioutil.ReadFile(".minecraft/versions/" + f.Download + "/" + f.Download + ".json")
	if err != nil {
		panic(err)
	}
	dl, err := download.Newlibraries(b)
	errr(err)
	fmt.Println("正在下载游戏核心")
	err = dl.Downjar(f.Atype)
	errr(err)
	fmt.Println("完成")
	fmt.Println("正在下载库文件")
	f.dd(dl, false)
	fmt.Println("完成")
	fmt.Println("正在下载资源文件")
	f.dd(dl, true)
	fmt.Println("完成")
	fmt.Println("正在下载解压 natives 库")
	err = dl.Unzip(f.Atype, f.Downint)
	if err != nil {
		fmt.Println(err)
		fmt.Println("下载失败")
		os.Exit(0)
	}
	fmt.Println("完成")
}

func (f Flag) dd(l download.Libraries, a bool) {
	ch := make(chan int, 5)
	e := make(chan error)
	var err error
	go func() {
		if a {
			err = l.Downassets(f.Atype, f.Downint, ch)
		} else {
			err = l.Downlibrarie(f.Atype, f.Downint, ch)
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
			if f.Outmsg {
				fmt.Println(i)
			}
		case err := <-e:
			log.Fatal(err)
			break b
		}
	}
}

func errr(err error) {
	if err != nil {
		if err.Error() == "proxy err" {
			fmt.Println(err)
			log.Fatalln("设置的代理有误")
		} else {
			fmt.Println(err)
			log.Fatalln("可能是网络问题，可再次尝试")
		}
		log.Fatal(err)
	}
}
