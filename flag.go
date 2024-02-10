package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	aflag "github.com/xmdhs/gomclauncher/flag"
	"github.com/xmdhs/gomclauncher/lang"
	"github.com/xmdhs/gomclauncher/launcher"
)

var f *aflag.Flag = aflag.NewFlag()

var (
	credit    bool
	update    bool
	list      bool
	remove    bool
	ms        bool
	v         bool
	tidy      bool
	buildDate string
	buildOn   string
	uselang   string
)

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
	flag.StringVar(&f.Runflag, "flag", "", lang.Lang("flagusage"))
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
