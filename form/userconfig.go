// 由res2go IDE插件自动生成，不要编辑。
package form

import (
    "github.com/ying32/govcl/vcl"
)

type TUserConfig struct {
    *vcl.TForm

    //::private::
    TUserConfigFields
}

var UserConfig *TUserConfig




// vcl.Application.CreateForm(&UserConfig)

func NewUserConfig(owner vcl.IComponent) (root *TUserConfig)  {
    vcl.CreateResForm(owner, &root)
    return
}

var userConfigBytes = []byte("\x54\x50\x46\x30\x0B\x54\x55\x73\x65\x72\x43\x6F\x6E\x66\x69\x67\x0A\x55\x73\x65\x72\x43\x6F\x6E\x66\x69\x67\x04\x4C\x65\x66\x74\x03\x6C\x02\x06\x48\x65\x69\x67\x68\x74\x03\xF0\x00\x03\x54\x6F\x70\x03\xD2\x00\x05\x57\x69\x64\x74\x68\x03\x40\x01\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0A\x55\x73\x65\x72\x43\x6F\x6E\x66\x69\x67\x0A\x4C\x43\x4C\x56\x65\x72\x73\x69\x6F\x6E\x06\x08\x32\x2E\x30\x2E\x31\x30\x2E\x30\x00\x00")

// 注册Form资源  
var _ = vcl.RegisterFormResource(UserConfig, &userConfigBytes)
