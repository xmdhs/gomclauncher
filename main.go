package main

import (
	"flag"
	"fmt"
	"gomclauncher/auth"
	aflag "gomclauncher/flag"
	"gomclauncher/launcher"
	"os"
	"runtime"
	"strings"
)

func main() {
	auth.Proxyaddr = f.Proxy
	if credit {
		credits()
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

var credit bool

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
	flag.IntVar(&f.Downint, "int", 32, "下载文件时使用的协程数。默认为 32")
	flag.StringVar(&f.Run, "run", "", `尝试启动的版本`)
	flag.BoolVar(&f.Runlist, "runlist", false, "显示所有可启动的版本")
	flag.StringVar(&f.RAM, "ram", "2048", `分配启动游戏的内存大小(mb)`)
	flag.StringVar(&f.Aflag, "flag", "", "自定的启动参数，比如 -XX:+AggressiveOpts -XX:+UseCompressedOops")
	flag.StringVar(&f.Proxy, `proxy`, "", `设置下载用的代理(http)`)
	flag.StringVar(&f.Atype, "type", "", `设置下载源。可选 bmclapi 和 mcbbs ，不设置此项则使用官方下载源`)
	flag.BoolVar(&f.Independent, "independent", false, "是否开启版本隔离")
	flag.BoolVar(&f.Outmsg, "test", true, "启动游戏前是否效验文件的完整和正确性")
	flag.BoolVar(&credit, "credits", false, "")
	flag.Parse()
}

func credits() {
	fmt.Println(`使用了 bmclapi 作为下载源，地址 https://bmclapidoc.bangbang93.com/`)
	fmt.Println(`使用了 github.com/google/uuid 用于生成 uuid ，开源协议`)
	fmt.Println(`
	Copyright (c) 2009,2014 Google Inc. All rights reserved.

	Redistribution and use in source and binary forms, with or without
	modification, are permitted provided that the following conditions are
	met:
	
	   * Redistributions of source code must retain the above copyright
	notice, this list of conditions and the following disclaimer.
	   * Redistributions in binary form must reproduce the above
	copyright notice, this list of conditions and the following disclaimer
	in the documentation and/or other materials provided with the
	distribution.
	   * Neither the name of Google Inc. nor the names of its
	contributors may be used to endorse or promote products derived from
	this software without specific prior written permission.
	
	THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
	"AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
	LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
	A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
	OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
	SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
	LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
	DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
	THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
	(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
	OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
	`)
}
