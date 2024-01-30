package flag

import (
	"context"
	"fmt"

	"github.com/xmdhs/gomclauncher/download"
	"github.com/xmdhs/gomclauncher/lang"
)

func (f *Flag) Arunlist() {
	l, err := download.Getversionlist(context.Background(), f.Atype, func(s string) { fmt.Println(s) })
	errr(err)
	m := make(map[string]struct{})
	for _, v := range l.Versions {
		m[v.Type] = struct{}{}
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
		fmt.Println(lang.Lang("runlist"))
		for k := range m {
			fmt.Println(k)
		}
	}

}
