package flag

import (
	"fmt"

	"github.com/xmdhs/gomclauncher/download"
)

func (f Flag) Arunlist() {
	l, err := download.Getversionlist(f.Atype)
	errr(err)
	m := make(map[string]bool)
	for _, v := range l.Versions {
		m[v.Type] = true
	}
	var ok bool
	for k := range m {
		if f.Verlist == k {
			ok = true
		}
	}
	if ok {
		for _, v := range l.Versions {
			if v.Type == f.Verlist {
				fmt.Println(v.ID)
			}
		}
	} else {
		fmt.Println("可选: ")
		for k := range m {
			fmt.Println(k)
		}
	}

}
