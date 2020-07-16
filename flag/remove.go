package flag

import (
	"fmt"
	"os"

	"github.com/xmdhs/gomclauncher/auth"
)

func (f Flag) Remove() {
	if f.Email == "" {
		fmt.Println("请设置 -email 参数来选择要删除的保存的账号")
		os.Exit(0)
	}
	if _, ok := gmlconfig[auth.ApiAddress][f.Email]; !ok {
		fmt.Println(auth.ApiAddress, f.Email, "不存在")
	} else {
		delete(gmlconfig[auth.ApiAddress], f.Email)
		fmt.Println("成功删除", auth.ApiAddress, f.Email)
		if len(gmlconfig[auth.ApiAddress]) == 0 {
			delete(gmlconfig, auth.ApiAddress)
		}
		saveconfig()
	}
	os.Exit(0)
}
