// 由res2go IDE插件自动生成。
package main

import (
	"github.com/ying32/govcl/vcl"
	"gomclauncherGUI/UIcode/form"
)

func main() {
	vcl.Application.SetScaled(true)
	vcl.Application.SetTitle("gomclauncherGUI")
	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)
	vcl.Application.CreateForm(form.MainWindow)
	vcl.Application.Run()
}
