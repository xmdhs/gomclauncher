package flag

import (
	"bufio"
	"fmt"
	"gomclauncher/download"
	"os"
)

func (f Flag) Arunlist() {
	l, err := download.Getversionlist(f.Atype)
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
