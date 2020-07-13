package flag

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/xmdhs/gomclauncher/download"
	"github.com/xmdhs/gomclauncher/launcher"
)

type Flag struct {
	launcher.Gameinfo
	Atype       string
	Downint     int
	Username    string
	Password    string
	Email       string
	Download    string
	Verlist     string
	Run         string
	Runlist     bool
	Runram      string
	Runflag     string
	Proxy       string
	Independent bool
	Outmsg      bool
}

func (f Flag) D() {
	l, err := download.Getversionlist(f.Atype)
	errr(err)
	err = l.Downjson(f.Download)
	if !(f.Run != "" && err != nil && err.Error() == "no such") {
		errr(err)
	}
	var b []byte
	if f.Run != "" {
		b, err = ioutil.ReadFile(launcher.Minecraft + "/versions/" + f.Run + "/" + f.Run + ".json")
	} else {
		b, err = ioutil.ReadFile(launcher.Minecraft + "/versions/" + f.Download + "/" + f.Download + ".json")
	}
	if err != nil {
		panic(err)
	}
	dl, err := download.Newlibraries(b, f.Atype)
	errr(err)
	if f.Outmsg {
		fmt.Println("正在验证游戏核心")
	} else {
		fmt.Println("正在下载游戏核心")
	}
	err = dl.Downjar(f.Download)
	errr(err)
	fmt.Println("完成")
	if f.Outmsg {
		fmt.Println("正在验证库文件")
	} else {
		fmt.Println("正在下载库文件")
	}
	f.dd(dl, false)
	fmt.Println("完成")
	if f.Outmsg {
		fmt.Println("正在验证资源文件")
	} else {
		fmt.Println("正在下载资源文件")
	}
	f.dd(dl, true)
	fmt.Println("完成")
	if f.Outmsg {
		fmt.Println("正在验证解压 natives 库")
	} else {
		fmt.Println("正在下载解压 natives 库")
	}
	err = dl.Unzip(f.Downint)
	if err != nil {
		fmt.Println("下载失败")
		panic(err)
	}
	fmt.Println("完成")
}

func (f Flag) dd(l download.Libraries, a bool) {
	ch := make(chan int, 5)
	e := make(chan error)
	var err error
	go func() {
		if a {
			err = l.Downassets(f.Downint, ch)
		} else {
			err = l.Downlibrarie(f.Downint, ch)
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
			if !f.Outmsg {
				fmt.Println(i)
			}
		case err := <-e:
			errr(err)
		}
	}
}

func errr(err error) {
	if err != nil {
		switch err.Error() {
		case "no such":
			fmt.Println("没有这个版本")
		case "file download fail":
			fmt.Println("失败次数过多，尝试切换下载源或者重新尝试")
		default:
			panic(err)
		}
		os.Exit(0)
	}
}
