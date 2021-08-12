// 由res2go IDE插件自动生成，不要编辑。
package form

import (
    "github.com/ying32/govcl/vcl"
)

type TMainWindow struct {
    *vcl.TForm
    StaticText1 *vcl.TStaticText
    GroupBox1   *vcl.TGroupBox
    Image1      *vcl.TImage
    StaticText2 *vcl.TStaticText
    StaticText3 *vcl.TStaticText
    Image2      *vcl.TImage
    Image3      *vcl.TImage
    Image4      *vcl.TImage

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
