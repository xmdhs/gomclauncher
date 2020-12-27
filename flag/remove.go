package flag

import (
	"fmt"
)

func (f Flag) Remove(ms bool) {
	ApiAddress := f.ApiAddress
	if ms {
		ApiAddress = "ms"
	}
	if f.Email == "" {
		fmt.Println("请设置 -email 参数来选择要删除的保存的账号")
	}
	if _, ok := f.Gmlconfig[ApiAddress][f.Email]; !ok {
		fmt.Println(ApiAddress, f.Email, "不存在")
	} else {
		delete(f.Gmlconfig[ApiAddress], f.Email)
		fmt.Println("成功删除", ApiAddress, f.Email)
		if len(f.Gmlconfig[ApiAddress]) == 0 {
			delete(f.Gmlconfig, ApiAddress)
		}
		saveconfig(f.Gmlconfig)
	}
}
