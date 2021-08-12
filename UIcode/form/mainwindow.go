// 由res2go IDE插件自动生成，不要编辑。
package form

import (
    "github.com/ying32/govcl/vcl"
    _ "embed"
)

type TMainWindow struct {
    *vcl.TForm

    //::private::
    TMainWindowFields
}

var MainWindow *TMainWindow




// vcl.Application.CreateForm(&MainWindow)

func NewMainWindow(owner vcl.IComponent) (root *TMainWindow)  {
    vcl.CreateResForm(owner, &root)
    return
}

//go:embed resources/mainwindow.gfm
var mainWindowBytes []byte

// 注册Form资源  
var _ = vcl.RegisterFormResource(MainWindow, &mainWindowBytes)
