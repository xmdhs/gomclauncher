package flag

import (
	"fmt"

	"github.com/xmdhs/gomclauncher/lang"
)

func (f *Flag) Remove(ms bool) {
	ApiAddress := f.ApiAddress
	if ms {
		ApiAddress = "ms"
	}
	if f.Email == "" {
		fmt.Println(lang.Lang("emailnil"))
	}
	if _, ok := f.Gmlconfig[ApiAddress][f.Email]; !ok {
		fmt.Println(ApiAddress, f.Email, lang.Lang("nofind"))
	} else {
		delete(f.Gmlconfig[ApiAddress], f.Email)
		fmt.Println(lang.Lang("removeok"), ApiAddress, f.Email)
		if len(f.Gmlconfig[ApiAddress]) == 0 {
			delete(f.Gmlconfig, ApiAddress)
		}
		saveconfig(f.Gmlconfig)
	}
}
