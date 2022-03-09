package main

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/xmdhs/gomclauncher/auth"
	aflag "github.com/xmdhs/gomclauncher/flag"
	"github.com/xmdhs/gomclauncher/lang"
	"github.com/xmdhs/gomclauncher/launcher"
)

func main() {
	if v {
		version()
	}
	if tidy {
		f.Tidy()
		return
	}
	if f.Proxy != "" {
		proxy, err := url.Parse(f.Proxy)
		if err != nil {
			panic(err)
		}
		auth.Transport.Proxy = http.ProxyURL(proxy)
	}
	if credit {
		credits()
	}
	if f.Verlist != "" {
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
	if list {
		f.Listname()
	}
	if f.ApiAddress != "" {
		f.Authlib()
	} else {
		f.ApiAddress = "https://authserver.mojang.com"
	}
	if remove {
		f.Remove(ms)
		return
	}
	if ms {
		f.MsLogin()
	} else {
		if f.Email != "" {
			f.Aonline()
		} else {
			f.UUID = aflag.UUIDgen(f.Name)
			f.AccessToken = f.UUID
		}
	}
	if f.Runflag != "" {
		s := strings.Split(f.Runflag, " ")
		f.Flag = s
	}
	f.Gameinfo.RAM = f.RAM
	if f.Run != "" {
		if f.Name == "" && f.Email == "" {
			fmt.Println(lang.Lang("nousername"))
		} else {
			f.Arun()
		}
	}
	if update {
		check()
	}
}

var f *aflag.Flag = aflag.NewFlag()

func init() {
	f.Gmlconfig = make(aflag.Gmlconfig)
	b, err := ioutil.ReadFile("gml.json")
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		panic(err)
	}
	err = json.Unmarshal(b, &f.Gmlconfig)
	if err != nil {
		fmt.Println(lang.Lang("jsonBreak"))
		panic(err)
	}
}

var (
	credit    bool
	update    bool
	list      bool
	remove    bool
	ms        bool
	v         bool
	tidy      bool
	gitHash   string
	buildDate string
	buildOn   string
	uselang   string
)

func init() {
	str, err := os.Getwd()
	str = strings.ReplaceAll(str, `\`, `/`)
	if err != nil {
		panic(err)
	}
	f.Minecraftpath = str + "/" + launcher.Minecraft
	flag.StringVar(&f.Name, "username", "", lang.Lang("username"))
	flag.StringVar(&f.Email, "email", "", lang.Lang("emailusage"))
	flag.StringVar(&f.Password, "password", "", lang.Lang("emailusage"))
	flag.StringVar(&f.Download, "downver", "", lang.Lang("Downloadusage"))
	flag.StringVar(&f.Verlist, "verlist", "", lang.Lang("verlistusage"))
	flag.IntVar(&f.Downint, "int", 64, lang.Lang("intusage"))
	flag.StringVar(&f.Run, "run", "", lang.Lang("runusage"))
	flag.BoolVar(&f.Runlist, "runlist", false, lang.Lang("runlistusage"))
	flag.StringVar(&f.RAM, "ram", "2048", lang.Lang("ramusage"))
	flag.StringVar(&f.Runflag, "flag", "-XX:+UseG1GC", lang.Lang("flagusage"))
	flag.StringVar(&f.Proxy, `proxy`, "", lang.Lang("proxyusage"))
	flag.StringVar(&f.Atype, "type", "", lang.Lang("typeusage"))
	flag.BoolVar(&f.Independent, "independent", true, lang.Lang("Independentusage"))
	flag.BoolVar(&f.Outmsg, "test", true, lang.Lang("testusage"))
	flag.BoolVar(&credit, "credits", false, lang.Lang("creditsusage"))
	flag.BoolVar(&update, "update", true, lang.Lang("updateusage"))
	flag.BoolVar(&f.Log, "log", false, lang.Lang("logusage"))
	flag.StringVar(&f.ApiAddress, "yggdrasil", "", lang.Lang("yggdrasilusage"))
	flag.BoolVar(&list, "list", false, lang.Lang("listusage"))
	flag.BoolVar(&remove, "remove", false, lang.Lang("removeusage"))
	flag.StringVar(&f.JavePath, "javapath", "java", lang.Lang("javapathusage"))
	flag.BoolVar(&ms, "ms", false, lang.Lang("msusage"))
	//flag.StringVar(&uselang, "lang", "", lang.Lang("langusage"))
	flag.BoolVar(&v, "v", false, lang.Lang("vusage"))
	flag.BoolVar(&tidy, "tidy", false, lang.Lang("tidy"))
	flag.Parse()
	if uselang != "" {
		err := lang.Setlanguge(uselang)
		if err != nil {
			fmt.Println(lang.Lang("nofindLanguage"))
			os.Exit(0)
		}
	}
}

func credits() {
	fmt.Println(lang.Lang("bmclapiinfo"))
	fmt.Println(lang.Lang("authlib-injectorinfo"))
	fmt.Println(lang.Lang("useproject"))
}

type up struct {
	Version string `json:"version"`
	Msg     string `json:"msg"`
}

func check() {
	version, err := checkByDns()
	if err != nil {
		log.Println(err)
		return
	}
	b, err := base64.StdEncoding.DecodeString(version)
	if err != nil {
		log.Println(err)
		return
	}
	u := up{}
	err = json.Unmarshal(b, &u)
	if err != nil {
		fmt.Println(lang.Lang("checkupdateerr"))
		fmt.Println(err)
		return
	}
	s, err := semver.NewVersion(u.Version)
	if err != nil {
		fmt.Println(lang.Lang("checkupdateerr"))
		fmt.Println(err)
		return
	}
	if s.GreaterThan(semver.MustParse(launcher.Launcherversion)) {
		fmt.Println(lang.Lang("checkupdate"), u.Version)
		fmt.Println(lang.Lang("nowversion"), launcher.Launcherversion)
		fmt.Println(lang.Lang("updateinfo"))
		fmt.Println(u.Msg)
	}
}

var Errtxt = errors.New("LookupTXT err")

func checkByDns() (string, error) {
	l, err := net.LookupTXT("gml.xmdhs.com")
	if err != nil {
		return "", fmt.Errorf("checkByDns: %w", err)
	}
	if len(l) != 1 {
		return "", fmt.Errorf("checkByDns: %w", Errtxt)
	}
	return l[0], nil
}

func version() {
	fmt.Println("gomclauncher-" + launcher.Launcherversion + "-" + gitHash)
	fmt.Println("Build date: " + buildDate)
	fmt.Println("Build on: " + buildOn)
	fmt.Println("Repository: https://github.com/xmdhs/gomclauncher")
}
