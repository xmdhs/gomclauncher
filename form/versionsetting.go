// 由res2go IDE插件自动生成，不要编辑。
package form

import (
    "github.com/ying32/govcl/vcl"
)

type TVersionSetting struct {
    *vcl.TForm
    ListBox1     *vcl.TListBox
    Image1       *vcl.TImage
    LabeledEdit1 *vcl.TLabeledEdit
    ListBox2     *vcl.TListBox

    //::private::
    TVersionSettingFields
}

var VersionSetting *TVersionSetting




// vcl.Application.CreateForm(&VersionSetting)

func NewVersionSetting(owner vcl.IComponent) (root *TVersionSetting)  {
    vcl.CreateResForm(owner, &root)
    return
}

var versionSettingBytes = []byte("\x54\x50\x46\x30\x0F\x54\x56\x65\x72\x73\x69\x6F\x6E\x53\x65\x74\x74\x69\x6E\x67\x0E\x56\x65\x72\x73\x69\x6F\x6E\x53\x65\x74\x74\x69\x6E\x67\x04\x4C\x65\x66\x74\x03\x88\x01\x06\x48\x65\x69\x67\x68\x74\x03\x90\x01\x03\x54\x6F\x70\x03\x07\x01\x05\x57\x69\x64\x74\x68\x03\x58\x02\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0E\x56\x65\x72\x73\x69\x6F\x6E\x53\x65\x74\x74\x69\x6E\x67\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x03\x90\x01\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\x58\x02\x0A\x4C\x43\x4C\x56\x65\x72\x73\x69\x6F\x6E\x06\x08\x32\x2E\x30\x2E\x31\x30\x2E\x30\x00\x08\x54\x4C\x69\x73\x74\x42\x6F\x78\x08\x4C\x69\x73\x74\x42\x6F\x78\x31\x04\x4C\x65\x66\x74\x03\x90\x00\x06\x48\x65\x69\x67\x68\x74\x03\x50\x01\x03\x54\x6F\x70\x02\x00\x05\x57\x69\x64\x74\x68\x03\xC8\x01\x0A\x49\x74\x65\x6D\x48\x65\x69\x67\x68\x74\x02\x00\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x00\x06\x54\x49\x6D\x61\x67\x65\x06\x49\x6D\x61\x67\x65\x31\x04\x4C\x65\x66\x74\x03\x18\x02\x06\x48\x65\x69\x67\x68\x74\x02\x40\x03\x54\x6F\x70\x03\x50\x01\x05\x57\x69\x64\x74\x68\x02\x40\x0C\x50\x69\x63\x74\x75\x72\x65\x2E\x44\x61\x74\x61\x0A\x39\x13\x00\x00\x17\x54\x50\x6F\x72\x74\x61\x62\x6C\x65\x4E\x65\x74\x77\x6F\x72\x6B\x47\x72\x61\x70\x68\x69\x63\x1D\x13\x00\x00\x89\x50\x4E\x47\x0D\x0A\x1A\x0A\x00\x00\x00\x0D\x49\x48\x44\x52\x00\x00\x00\xC9\x00\x00\x00\xC8\x08\x06\x00\x00\x00\x42\x9A\xC5\xA0\x00\x00\x12\xE4\x49\x44\x41\x54\x78\x5E\xED\x9D\x07\xA8\x66\xC5\x15\xC7\x7F\x12\xB1\x65\x15\x8C\x5D\x89\xC6\x58\xA3\x62\xAC\x24\xB2\xF6\xC4\x5E\x88\xBA\x2A\xCA\xC6\x2E\x12\x6B\x82\xD8\x10\x89\x4A\xC0\x86\x26\xA8\x51\xB1\x2B\xB1\x6B\xC4\x96\xA8\x59\x35\x36\x5C\x62\x21\x6B\xEC\x44\x23\x2A\x82\x26\x16\xB0\xC6\x58\xC2\xFF\x39\xF7\xF9\xED\x7B\xDF\xF7\x6E\x9B\xB9\x77\xE6\xDE\x73\xE0\xE3\xBD\x7D\xDF\xCC\x99\x73\xFE\x33\xFF\x9D\x3B\x77\x66\xCE\x99\x07\x93\x90\x08\xAC\x06\xE8\xB3\x3A\xB0\x18\x30\x0D\x58\x78\xC2\xCF\x89\x7F\x93\x3D\x1F\x00\x1F\x0E\xFC\x1C\xFC\x3D\xFB\xEE\x1D\xE0\x25\xE0\x45\xF7\xF3\x8B\x90\x8E\xF4\x59\xF7\x3C\x7D\x76\xDE\x93\xEF\x8B\x3B\x22\x0C\x12\x22\xFB\xDD\x53\x13\x85\xD4\xBC\x3C\x81\x34\x22\x90\x3E\xAF\x17\xAA\x6D\x85\x46\x22\x60\x24\x29\x3F\x38\xBE\x07\x6C\x0E\x6C\xE6\x7E\xEA\xDF\x31\xCB\x7B\xC0\x7D\xEE\xF3\x10\xF0\x5C\xCC\xC6\xC6\x68\x9B\x91\x24\xBF\x57\x44\x02\x11\x62\xBA\xFB\xAC\x91\x5F\x25\xEA\x12\xAF\x39\xC2\x3C\x0A\x3C\x66\xA4\xC9\xEF\x2B\x23\xC9\x70\x8C\xB6\x02\xB6\x03\x36\x06\x36\xCC\x87\x31\xE9\x12\xCF\x02\xB3\x81\x59\xC0\xED\xC0\xC7\x49\x7B\x13\xC0\x78\x23\xC9\x37\xA0\xAE\x0F\xEC\xEC\x3E\xEB\x04\xC0\x3A\x05\x95\x5A\xBF\x88\x28\xFA\xDC\x9B\x82\xC1\x4D\xD8\xD8\x77\x92\xAC\x38\x40\x8C\x2D\x9B\x00\x3C\xA1\x36\xE6\x0C\x10\xE6\x89\x84\xEC\xF6\x6E\x6A\x1F\x49\x22\x9F\x67\x0E\x90\x63\x3E\xEF\xA8\x76\x4F\xE1\x03\x8E\x30\x37\x02\x6F\x76\xCF\xBD\xA9\x3D\xEA\x13\x49\xE6\x07\x0E\x72\x9F\xBE\x3E\x4E\xD5\x1D\xDF\x6F\x01\x97\xBA\xCF\xAB\x75\x95\xA5\x52\xBF\x0F\x24\xD1\xE6\x9D\xC8\x71\x20\xB0\x66\x2A\x1D\x13\xB9\x9D\xEF\x02\x97\x39\xB2\x68\x2F\xA6\xD3\xD2\x65\x92\x68\x87\x3B\x23\xC7\x2A\x9D\xEE\xC5\xF6\x9C\xD3\x49\x80\x8C\x2C\xCF\xB4\x67\x46\xD8\x96\xBB\x48\x92\x65\x06\xC8\xB1\x42\x58\xF8\x4C\xBB\x43\xE0\x33\x37\xAB\x88\x30\x4F\x75\x0D\x95\xAE\x91\xE4\x30\xE0\x58\x60\xF9\xAE\x75\x54\x42\xFE\x9C\x01\x9C\x09\xE8\x91\xAC\x13\xD2\x15\x92\x6C\x01\x1C\x07\x6C\xD3\x89\x5E\x49\xDF\x09\x1D\x7D\x11\x51\xAE\x4A\xDF\x15\x48\x9D\x24\x4B\xBB\x99\xE3\x57\x5D\xE8\x8C\x0E\xFA\x70\x0B\xA0\x99\xE5\xF1\x94\x7D\x4B\x99\x24\x07\x3B\x82\xAC\x9C\x72\x07\xF4\xC0\x76\xAD\x57\xB2\x47\x30\x2D\xF4\x93\x93\x14\x49\xA2\x83\x86\x7A\xB4\xDA\x29\x39\xB4\xFB\x6D\xB0\x76\xF0\x45\x96\xEB\x52\x83\x21\x35\x92\x1C\xE3\x9E\x75\x53\xC3\xD9\xEC\xFD\x06\x81\x8B\xDC\x13\x80\x2E\x8F\x25\x21\xA9\x90\x64\x39\x47\x8E\xBD\x93\x40\xD5\x8C\xCC\x43\x40\xA7\x8E\xF5\x16\xF2\xE1\xBC\x82\x31\x7C\x9F\x02\x49\x76\x70\x04\x49\xFD\x1E\x47\x0C\xFD\x1D\x93\x0D\x9F\x3A\xA2\x9C\x17\x93\x51\xC3\x6C\x89\x9D\x24\x27\x02\xBF\x89\x1D\x44\xB3\xAF\x16\x02\x97\x3B\xB2\xE8\xCE\x7E\x94\x12\x2B\x49\x74\x84\x5D\xEF\xD9\x67\x44\x89\x9A\x19\xE5\x1B\x01\xED\xD2\xEB\xF1\x4B\xD7\x8C\xA3\x93\x18\x49\xB2\x8B\x23\x88\xBD\xDA\x8D\x6E\xB8\x04\x35\xE8\x4B\x47\x94\xB3\x83\xB6\x52\x41\x79\x6C\x24\xD9\x0F\xB8\xA2\x82\x1F\x56\xA5\x3B\x08\x68\x8D\x72\x64\x4C\xEE\xC4\x44\x92\xC3\x81\xE8\x17\x71\x31\x75\x5E\x87\x6D\xB9\x19\xD8\x3D\x16\xFF\x62\x21\xC9\xF1\xC0\x69\xB1\x80\x62\x76\x44\x81\xC0\x23\xC0\x26\x31\x58\x12\x03\x49\xCE\x01\xEC\xEC\x55\x0C\xA3\x21\x3E\x1B\x14\x70\x6F\x03\xE0\xFD\x36\x4D\x6B\x9B\x24\x77\x00\x3B\xB6\x09\x80\xB5\x1D\x3D\x02\x1F\x01\x6B\x03\xAF\xB4\x65\x69\x9B\x24\x51\x0C\xDB\x55\xDB\x72\xDC\xDA\x4D\x0E\x01\xC5\x3F\x6B\x25\x6A\x4B\x5B\x24\xF9\x2A\xB9\x2E\x32\x83\x63\x40\xE0\x87\xC0\xD3\x4D\x1B\xD2\x06\x49\x14\x92\x46\x57\x6C\x4D\x0C\x81\x2A\x08\x7C\x07\x50\x7C\xE3\xC6\xA4\x69\x92\x68\x67\x75\xDD\xC6\xBC\xB3\x86\xBA\x88\x80\xFE\x93\xD5\x81\xD7\xC6\xA4\x49\x92\xDC\x05\x6C\xDF\x98\x67\xD6\x50\x97\x11\xD0\x58\x6A\xEC\x85\x4F\x53\x24\x51\x14\x8D\x03\xBA\xDC\x6B\xE6\x5B\xE3\x08\x9C\xE5\x8E\xB1\x04\x6F\xB8\x09\x92\xE8\x14\xAF\x4E\xF3\x9A\x18\x02\xBE\x11\x38\x01\x38\xDD\xB7\xD2\x89\xFA\x42\x93\xC4\x8E\x9A\x84\xEE\x41\xD3\x7F\x08\x70\x71\x48\x18\x42\x92\x44\xA7\x79\xFF\x18\xD2\x78\xD3\x6D\x08\x38\x04\x76\x05\x6E\x0D\x85\x46\x28\x92\x68\x93\xF0\x1E\x20\xF6\x54\x69\xA1\x70\x35\xBD\xCD\x22\xA0\xE0\xDD\x8A\xB9\x16\x24\x2E\x71\x28\x92\xDC\x09\xE8\xDA\xAD\x89\x21\xD0\x14\x02\xC1\xDE\x78\x85\x20\x89\xC2\xC6\xE8\x96\x99\x89\x21\xD0\x34\x02\xBA\xCD\xAA\x70\x53\x5E\xC5\x37\x49\xF6\xE9\x4A\x68\x4B\xAF\x28\x9B\xB2\x26\x11\xD8\x17\xB8\xDA\x67\x83\x3E\x49\xA2\x9C\x83\x77\x03\xCA\x6B\x6E\x52\x1E\x81\x4F\xDC\x49\xD7\x2F\x00\xA5\x8A\x58\xB0\xBC\x0A\xAB\x01\xFC\x07\xD8\x16\x78\xD2\x17\x1A\xBE\x48\xB2\x80\x23\x88\x52\x39\x9B\x94\x43\xE0\x08\xE0\xFE\x21\xA9\xA2\x57\x73\xBB\xCA\xDA\x07\x98\xB7\x9C\xCA\xDE\x97\x7E\xD0\x11\x45\x61\x8B\x6A\x8B\x2F\x92\xFC\x1E\x38\xB4\xB6\x35\xFD\x52\xA0\xD4\xD0\x4A\x83\xAD\x8C\xB7\x53\xC9\x92\x2E\x5F\xE1\x8F\xFA\x05\x4F\x6D\x6F\x2F\x00\x94\x8A\xA3\xB6\xF8\x20\xC9\x5E\xC0\xB5\xB5\x2D\xE9\x97\x02\x3D\x5A\x2D\x54\xD2\xE5\xEB\x81\x3D\x4B\xD6\xE9\x7B\xF1\x9F\x03\x7F\xA8\x0B\x42\x5D\x92\x4C\x03\x74\x17\x59\xE7\xFC\x4D\x8A\x23\xA0\x59\xF7\xC2\xE2\xC5\xC7\x4A\x2E\x02\xBC\x0D\x28\x41\xAA\x49\x31\x04\x5E\x00\xB4\x04\x10\x6E\x95\xA5\x2E\x49\x4E\x05\x4E\xAA\xDC\x7A\x3F\x2B\xEA\xCD\x8B\xDE\xC0\x54\x91\x5F\x00\x7A\x8C\x30\x29\x8E\xC0\xB9\xC0\x51\xC5\x8B\x4F\x2E\x59\x87\x24\xBA\x4E\xA9\x59\xC4\xF2\xA0\x97\xEB\x81\x3D\x80\x9B\xCA\x55\x19\x2F\xAD\x13\x0C\xFF\xAA\x58\xB7\xCF\xD5\x74\xAC\x5E\x9B\x8D\x95\xA4\x0E\x49\x14\x1B\x69\xB7\x4A\xAD\xF6\xBB\xD2\x77\x81\x37\x6A\x40\xA0\x54\x6B\x3F\xA8\x51\xBF\x8F\x55\x1F\x03\x36\x05\x3E\xAF\xE2\x7C\x55\x92\xE8\x71\xE1\xCA\x2A\x0D\xF6\xBC\xCE\x3F\xDD\x1E\x48\x1D\x18\x14\xE1\x52\x91\x2E\x4D\xCA\x21\x70\x0A\x70\x72\xB9\x2A\x5F\x97\xAE\x42\x12\xDD\x31\x56\x5E\x09\x4B\x85\x50\x1E\xF1\xBF\x02\x4A\x82\x5A\x47\xD4\xD1\xBF\xAE\xA3\xA0\xA7\x75\x35\x8B\x68\x36\xD1\xAC\x52\x4A\xAA\x90\x44\x9B\x5B\xDE\xCF\xC7\x94\xB2\x3A\xDD\xC2\x46\x92\x76\xFB\xAE\xD2\x21\xC8\xB2\x24\x59\xCF\xE7\x76\x7F\xBB\x78\xB5\xD2\xBA\x91\xA4\x15\xD8\xE7\x6A\xF4\x20\x40\xD7\xC9\x0B\x4B\x59\x92\xD8\xCE\x7A\x61\x68\x87\x16\x34\x92\xD4\xC3\xCF\x47\x6D\x2D\x15\xF4\xD8\x55\x58\xCA\x90\x64\x2D\x37\x8B\xD8\x2B\xDF\xC2\xF0\x4E\x2A\x68\x24\xA9\x8E\x9D\xCF\x9A\x7A\x2B\x5B\xF8\xD6\x6C\x19\x92\xFC\xAE\xEE\xA6\x8C\x4F\x2F\x13\xD5\x65\x24\x89\xA3\xE3\x6E\x03\x7E\x56\xD4\x94\xA2\x24\xD1\x75\x5C\x1D\x3D\xD6\x31\x14\x93\xEA\x08\x18\x49\xAA\x63\xE7\xBB\xA6\xDE\x32\xAA\x3F\x72\xA5\x28\x49\x74\xE3\x4B\x39\xD4\x4D\xEA\x21\x60\x24\xA9\x87\x9F\xCF\xDA\xDA\xE7\xDB\xBF\x88\xC2\x22\x24\xD1\x51\x08\xCD\x22\xDA\x1F\x31\xA9\x87\x80\x91\xA4\x1E\x7E\x3E\x6B\x2B\x68\xFB\x3A\x45\x02\x70\x17\x21\x89\x05\x97\xF3\xD7\x35\x46\x12\x7F\x58\xFA\xD0\xA4\x04\x52\x47\xE7\x29\xCA\x23\xC9\xB2\x80\x82\x5C\x2F\x95\xA7\xC8\xBE\x2F\x84\x80\x91\xA4\x10\x4C\x8D\x15\x52\xEE\x78\xCD\x26\x53\x9E\xA5\xCB\x23\x89\x8E\x18\xEB\xAD\x96\x89\x1F\x04\x8C\x24\x7E\x70\xF4\xA9\x45\x33\x89\x66\x94\x91\x92\x47\x92\x59\xC0\x4F\x7C\x5A\xD4\x73\x5D\x46\x92\xF8\x06\xC0\xA3\xC0\xC6\x55\x49\xA2\xDB\x86\x7F\x8F\xCF\xA7\xA4\x2D\x32\x92\xC4\xD9\x7D\xBA\xBD\xF8\xD0\x28\xD3\xA6\x9A\x49\x14\x09\x5E\x8B\x76\x13\x7F\x08\x18\x49\xFC\x61\xE9\x53\x93\x96\x14\x23\x33\x40\x4F\x45\x12\x1D\x29\xFE\xB1\x4F\x4B\x4C\xD7\xD8\xE6\x95\x1D\x95\x8F\x6F\x20\xBC\xE6\x2E\xB2\x7D\x3C\xCC\xB4\x51\x24\x99\xEE\xAE\xE6\xC6\xE7\x4E\xDA\x16\x19\x49\xE2\xED\xBF\x99\xC0\x35\x65\x48\x72\x1A\x70\x7C\xBC\xFE\x24\x6B\x99\x91\x24\xDE\xAE\xBB\x05\x98\x51\x86\x24\xFF\x00\x74\xEA\xD7\xC4\x2F\x02\x46\x12\xBF\x78\xFA\xD4\xA6\x1D\x78\xC5\x0E\x78\x71\xA2\xD2\x61\x8F\x5B\x5B\xBB\xDC\x22\x3E\x0D\x30\x5D\x5F\x23\x60\x24\x89\x7B\x24\x0C\x4D\x2F\x37\x8C\x24\xDA\x58\x19\xB9\xD2\x8F\xDB\xC7\xE8\xAD\x33\x92\xC4\xDD\x45\xB3\x81\x8D\x8A\xCC\x24\x2A\x68\x71\x67\xC3\x74\xA6\x0F\x92\x28\x08\x44\xA5\xA8\x1F\x61\x5C\xEA\x9C\xD6\x95\x5C\x74\xFF\x71\xC7\x26\xCE\x24\x16\xFC\x2C\x6C\x9F\x1B\x49\xC2\xE2\xEB\x43\xBB\x8E\xCF\xCF\x15\x2E\x6B\x22\x49\xF6\x1E\xF5\x1A\xCC\x47\xEB\xA6\xC3\xD6\x24\x09\x8C\x01\x05\x89\x50\xB0\x88\x91\x33\xC9\xF9\xBE\xC2\xD5\x27\x00\x46\x1B\x26\xFA\x98\x49\x2C\xEE\x56\xD8\x9E\x53\x72\x52\xE5\x86\x19\x49\x12\xBD\xFE\xD2\x55\x5D\x93\x30\x08\x18\x49\xC2\xE0\xEA\x5B\xAB\x02\x2F\x3E\x9F\x29\x1D\x7C\xDC\x12\x39\x26\xBD\x23\xF6\xDD\x7A\xCF\xF5\xF9\x20\x89\x2D\xDC\xC3\x0F\x22\x25\xFF\x19\x8F\xDE\x3F\x48\x12\x3D\x87\x5D\x12\xBE\xFD\x5E\xB7\xE0\x83\x24\xF6\xB8\x15\x7E\x08\xCD\xB5\xFB\x3E\x48\x92\x1B\x00\xA5\x05\x30\x09\x87\x80\x91\x24\x1C\xB6\x3E\x35\x2B\x39\xE9\x12\xC3\x1E\xB7\xDE\x05\x16\xF5\xD9\x92\xE9\x9A\x84\x80\x91\x24\x9D\x41\xA1\x6B\xBD\x73\x64\x6E\x36\x93\x28\x67\x86\x8E\x0B\x9B\x84\x45\xC0\x07\x49\x6C\x4D\x12\xB6\x8F\x32\xED\xE3\xF9\x16\x33\x92\xE8\x8A\xAE\xAE\xEA\x9A\x84\x45\xC0\x07\x49\x6C\x4D\x12\xB6\x8F\x32\xED\xBA\x70\x38\x96\xEA\x30\x23\x89\xE5\xE2\x6B\x06\x78\x23\x49\x33\x38\xFB\x68\xE5\xC6\x2C\xDB\x71\x46\x92\xDF\x02\xBF\xF4\xA1\xD9\x74\x4C\x89\x80\x0F\x92\xD8\xE3\x56\x33\x83\x4C\xF1\x1D\xD6\x1D\x9C\x49\xFE\x04\x6C\xD7\x4C\xDB\xBD\x6E\xC5\x07\x49\xEC\x71\xAB\x99\x21\xA4\xAB\xBC\xDF\x1E\x24\x89\x72\xF9\xE9\xF4\xA3\x49\x58\x04\x8C\x24\x61\xF1\xF5\xAD\x7D\x79\xE0\x75\x3D\x6E\x7D\xAB\x6A\x56\x52\xDF\x16\xF5\x40\x9F\x91\x24\xAD\x4E\xFE\x29\x70\x9F\x48\xA2\x2B\x8B\x4A\x7B\x6C\x12\x1E\x01\x1F\x24\xB1\x35\x49\xF8\x7E\xCA\x5A\x38\x14\xB8\x50\x24\x51\x32\x93\x5B\x9B\x6B\xB7\xD7\x2D\xF9\x20\x89\xAD\x49\x9A\x1B\x42\x63\xF1\xB8\x44\x92\x63\x81\x33\x9A\x6B\xB7\xD7\x2D\x19\x49\xD2\xEA\xFE\x3F\x03\xDB\x8B\x24\x22\x88\x88\x62\x12\x1E\x01\x23\x49\x78\x8C\x7D\xB6\xF0\x37\x5D\x65\x17\x49\x52\xCF\xA8\xFB\x89\xBB\x93\xFC\x6F\x9F\xE8\x04\xD4\x55\x37\x82\xA3\xD6\x24\x9B\x07\xB4\xCF\x97\xEA\x05\x81\x55\x12\x4F\xFE\xA4\xB5\xFA\x9A\x22\xC9\xD5\x80\xCE\xA9\xA4\x26\x47\x00\xF7\xDB\x4B\x87\xE8\xBB\x4D\xE7\x02\x37\x04\xCE\x05\x96\x8B\xDE\xDA\xB9\x0D\xD4\x79\xC6\x15\x44\x12\xA5\xEA\xDD\x25\x21\xE3\x9F\x75\x1B\x9F\xAF\x27\x64\xB3\x99\xFA\x35\x02\x0F\x24\x32\x0B\x66\xFD\xF5\x9E\x66\x42\x91\xE4\x2F\x80\xDE\x07\xA7\x20\x7A\xA4\x5A\x32\x05\x43\xCD\xC6\x91\x08\x5C\x9F\x9D\x89\x4A\x00\xA3\xFF\x01\xF3\x89\x24\x29\x45\x8F\x2F\x9C\x56\x38\x81\x0E\xE8\xAB\x89\x8B\x00\x6F\x03\xF3\x27\x02\xC0\x82\x22\x49\x2A\x71\x7F\x15\xC0\xDB\x5E\x55\x27\x32\xB2\x72\xCC\x4C\xE9\xD4\xF9\x12\x22\xC9\xAB\x5A\x9C\x24\x80\xFD\xEA\x16\xA8\x22\x81\x5E\x2A\x66\x62\x4A\x41\x10\xBF\x2F\x92\xE8\x3E\xEF\x62\xC5\x7C\x6B\xAD\x94\x5E\xF3\x2E\xD4\x5A\xEB\xD6\x70\x08\x04\xF4\x7A\x55\x47\xA2\x62\x97\xB5\x45\x92\xFF\x6A\x71\x12\xB9\xA5\x4F\x03\xCA\xE1\x68\xD2\x1D\x04\xAE\x00\xF6\x4B\xC0\x9D\xE9\x46\x92\x04\x7A\xA9\xA3\x26\x26\x45\x12\x7B\xDC\xEA\xE8\x28\x8C\xDC\xAD\xA4\x1E\xB7\x6C\xE1\x1E\xF9\x68\xEA\xA0\x79\xC9\x2D\xDC\xED\x15\x70\x07\x47\x61\xE4\x2E\x25\xF7\x0A\xD8\x36\x13\x23\x1F\x51\x1D\x33\x2F\xC9\xCD\x44\x3B\x96\xD2\xB1\x51\x18\xB9\x3B\x49\x1E\x4B\xB1\x03\x8E\x91\x8F\xAA\x0E\x99\x97\xEC\x01\x47\x3B\x2A\xDF\xA1\x51\x18\xA1\x2B\x9D\x38\x2A\x6F\x97\xAE\x9A\x1D\x59\x75\x2F\x5D\xE9\x8E\xFB\x66\xCD\x9A\x5C\xA9\xB5\x4E\x5D\xBA\xB2\xEB\xBB\x95\xC6\x40\xA5\x4A\x76\x7D\xB7\x12\x6C\xAD\x55\x1A\xBF\xBE\x6B\x81\x20\x9A\xEB\x03\x23\x49\x73\x58\xFB\x68\x69\x3C\x10\x84\x85\x14\xF2\x01\x67\x31\x1D\x46\x92\x62\x38\xC5\x52\x6A\x3C\xA4\x90\x05\xA7\x6B\xAE\x4B\x8C\x24\xCD\x61\xED\xA3\xA5\xF1\xE0\x74\x16\xE6\xD4\x07\x9C\xC5\x74\x18\x49\x8A\xE1\x14\x4B\xA9\xF1\x30\xA7\x32\xC8\x02\x66\x37\xD3\x2D\x46\x92\x66\x70\xF6\xD5\xCA\x78\xC0\x6C\x29\xB4\xD4\x0B\xBE\x60\x9D\x5A\x8F\x91\xA4\x19\x9C\x7D\xB4\x32\x29\xF5\x82\x25\xF1\xF1\x01\x6B\xBE\x0E\x23\x49\x3E\x46\xB1\x94\x98\x94\xC4\x27\xA5\x53\x99\xB1\x80\x58\xC5\x0E\x23\x49\x15\xD4\xDA\xA9\x33\x29\x1D\x9C\x25\x16\x6D\xA6\x23\x8C\x24\xCD\xE0\xEC\xA3\x95\x49\x89\x45\x2D\x45\xB5\x0F\x58\xF3\x75\x18\x49\xF2\x31\x8A\xA5\xC4\xA4\x14\xD5\x32\xEC\x5D\x60\xD1\x58\x2C\xEC\xA8\x1D\x46\x92\x74\x3A\x76\x1D\x60\x8E\xCC\xCD\xB2\xEF\xEA\xF7\x1B\x80\x3D\xD2\xF1\x21\x49\x4B\x8D\x24\x69\x74\x9B\xE2\x3E\x2C\x91\x99\x3A\x48\x92\x83\x80\x4B\xD2\xF0\x21\x59\x2B\x8D\x24\x69\x74\xDD\x2D\xC0\x8C\x61\x24\x59\xD5\x22\x24\x06\xEF\x41\x23\x49\x70\x88\xBD\x34\x70\x18\x70\xC1\x30\x92\xE8\x6F\x2F\x02\x22\x8B\x49\x18\x04\x8C\x24\x61\x70\xF5\xAD\x75\x0D\xE0\xF9\x51\x24\x39\x1F\x10\x8B\x4C\xC2\x20\x60\x24\x09\x83\xAB\x4F\xAD\x2F\x01\xAB\x0D\x2A\x1C\x5C\x93\xE8\xEF\x7B\x03\xD7\xF8\x6C\xD1\x74\xCD\x85\x80\x91\x24\xFE\x01\x71\x19\xA0\xF5\xF9\xB8\x4C\x24\x49\x4A\x41\xC3\xE2\x87\x7B\xB2\x85\x46\x92\xF8\x7B\x6D\x7F\xE0\xCA\xA9\x48\xA2\xEF\x66\x2B\xE3\x68\xFC\xBE\x24\x69\xA1\x91\x24\xFE\x6E\x5B\xC9\x25\xAA\x1D\x39\x93\xE8\x8B\x73\x94\xE0\x3D\x7E\x5F\x92\xB4\xD0\x48\x12\x77\xB7\x69\x82\xD8\x68\xA2\x89\x13\x1F\xB7\xF4\xFD\xD6\xC0\x3D\x71\xFB\x92\xAC\x75\x46\x92\xB8\xBB\xEE\x04\xE0\xF4\x22\x24\x51\x99\x54\xE2\x03\xC7\x0D\xB9\xAD\x49\x52\xEA\x9F\xAF\x5C\x52\x21\x6D\x83\xCC\x25\xC3\x66\x12\x15\x38\x0D\x50\x8E\x42\x13\xBF\x08\xD8\x4C\xE2\x17\x4F\x9F\xDA\xE6\xDA\x65\xCF\x5B\xB8\xEB\xFB\xE9\xC0\x23\x3E\x2D\x30\x5D\x63\x08\x18\x49\xE2\x1D\x08\x33\x47\x6D\x7F\x8C\x9A\x49\xE4\x4A\x4A\xD1\xE6\xE3\x85\x7E\x6E\xCB\x8C\x24\x71\xF6\xD4\x6B\xEE\x51\x4B\x57\x76\x27\xC9\x54\x24\x39\x11\xD0\xC5\x13\x13\x7F\x08\x18\x49\xFC\x61\xE9\x53\xD3\x58\x7C\xAD\x51\x0A\xA7\x22\x89\x12\x79\xEA\x9E\xAF\x89\x3F\x04\x8C\x24\xFE\xB0\xF4\xA9\x49\xB1\x95\x1F\xAA\x42\x12\xD5\x99\x05\xE8\x6A\xAF\x89\x1F\x04\x8C\x24\x7E\x70\xF4\xA9\xE5\x51\x60\xE3\xA9\x14\x4E\x35\x93\xA8\xDE\x51\x80\xA6\x22\x13\x3F\x08\x18\x49\xFC\xE0\xE8\x53\xCB\xD1\x6E\x03\x7D\xA4\xCE\x3C\x92\x2C\x0B\x3C\x05\x2C\xE5\xD3\xAA\x1E\xEB\x32\x92\xC4\xD5\xF9\xEF\x00\xBA\xA6\xFB\x46\x9D\x99\x44\x75\xB5\x78\xD7\x22\xDE\xA4\x3E\x02\x46\x92\xFA\x18\xFA\xD4\xA0\x23\x58\x9A\x49\xA6\x94\xBC\x99\x44\x95\x75\x32\xF8\x49\xE0\x3B\x79\xCA\xEC\xFB\x5C\x04\x8C\x24\xB9\x10\x35\x56\x40\x3B\xEC\x9A\x45\x9E\xCE\x6B\xB1\x08\x49\xA4\xE3\x4C\xE0\x98\x3C\x65\xF6\x7D\x2E\x02\x46\x92\x5C\x88\x1A\x2B\xA0\xE3\xF0\x3A\x16\x9F\x2B\x45\x49\xA2\x2B\xBD\x9A\x4D\xA6\xE5\x6A\xB4\x02\x53\x21\x60\x24\x89\x67\x7C\x28\x2D\x9F\xFA\x23\x57\x8A\x92\x44\x8A\xF4\x96\x4B\x6F\xBB\x4C\xAA\x23\x60\x24\xA9\x8E\x9D\xCF\x9A\xB7\x01\x4A\x5E\x55\x48\xCA\x90\x64\x2D\x37\x9B\xCC\x57\x48\xB3\x15\x1A\x86\x80\x91\x24\x8E\x71\xB1\x1B\xA0\xD4\xEC\x85\xA4\x0C\x49\xA4\x30\xF5\x4C\xBD\x85\x40\x09\x58\xC8\x48\x12\x10\xDC\x82\xAA\x1F\x06\x36\x2D\x58\x76\xAC\x58\x59\x92\xAC\xE7\x66\x93\x32\x6D\x58\xD9\x6F\x10\x30\x92\xB4\x3F\x1A\x14\xE4\x41\xC1\x1E\x0A\x4B\x59\x92\x48\xB1\x6E\x6E\x1D\x57\xB8\x05\x2B\x38\x88\x80\x91\xA4\xDD\xF1\x70\x17\xB0\x63\x59\x13\xAA\x90\x44\xFB\x25\x9A\xB2\x14\xC0\xCB\xA4\x1C\x02\x46\x92\x72\x78\xF9\x2C\xFD\xB9\x7B\xCC\xD2\x15\x90\x52\x52\x85\x24\x6A\x60\xDF\x89\x61\x57\x4A\xB5\xDA\xDF\xC2\xCA\x4D\xB9\x4A\x4D\xF7\xAF\x00\xF6\xAB\xA9\xA3\x8F\xD5\x4F\x06\x4E\xA9\xE2\x78\x55\x92\xA8\xAD\x9B\x01\xBD\x25\x30\x29\x87\x80\x72\xC1\x4C\x79\x56\x28\x47\xDD\x73\xEE\x82\x50\xB9\x56\xFB\x5D\x5A\xB3\x87\x16\xEB\x9A\x4D\x4A\x4B\x1D\x92\x6C\xE8\xAE\xF8\xDA\x2B\xE1\x72\xB0\x2B\xBD\xC5\x4D\xE5\xAA\x8C\x97\xB6\xE0\x81\xD5\x80\xD3\x3A\x44\xEB\x91\x4A\x52\x87\x24\x6A\xF0\x54\xE0\xA4\x4A\x2D\xF7\xB7\xD2\xD5\xEE\x71\xB5\x0A\x02\x96\xDB\xB2\x3C\x6A\xE7\xD6\xDD\x04\xAF\x4B\x12\x1D\x53\x51\xC0\x08\xDD\x62\x34\x29\x8E\xC0\xA1\xC0\x85\xC5\x8B\x8F\x95\x5C\x04\x78\x1B\x98\xBF\x64\xBD\x3E\x17\x7F\x01\xD0\xAD\x43\xE1\x56\x59\xEA\x92\x44\x0D\xEF\x05\x5C\x5B\xD9\x82\x7E\x56\xFC\x04\x58\xA8\xA4\xEB\xD7\x03\x7B\x96\xAC\xD3\xF7\xE2\xE3\x79\x0F\xEB\x00\xE1\x83\x24\x6A\xDF\x76\xE2\xCB\xF7\xC2\xB3\xC0\x76\xC0\xEB\x39\x55\x97\x04\x6E\xB7\xF8\xCC\xA5\x01\x56\x12\x1E\x2F\x69\x44\x7C\x91\x64\x01\xE0\x6E\x37\xB5\x95\xF6\xA6\xE7\x15\x8E\x00\xEE\x07\xF4\xD6\x6A\x50\x94\x23\x43\x0B\x4E\x6D\xDE\xCE\xDB\x73\x8C\xCA\xBA\xFF\x20\xB0\x2D\xF0\x69\xD9\x8A\xC3\xCA\xFB\x22\x89\x74\xAF\xEF\x88\xB2\xB8\x0F\xC3\x7A\xA8\x43\x8F\x60\xAF\x00\x5F\xB8\xBD\x94\x05\x7B\x88\x81\x0F\x97\x95\x14\x54\x04\xD1\xD5\x0E\x2F\xE2\x93\x24\x32\x68\x1F\xE0\x2A\x2F\x96\x99\x12\x43\xA0\x1A\x02\xDA\xE8\xD6\x1B\x44\x6F\xE2\x9B\x24\x32\xEC\x0C\xE0\x58\x6F\x16\x9A\x22\x43\xA0\x38\x02\xBA\x41\xEB\xFD\x5C\x61\x08\x92\xC8\xA5\x3B\x81\x1D\x8A\xFB\x66\x25\x0D\x81\xDA\x08\x54\x3A\xBC\x58\xA4\xD5\x50\x24\xD1\x75\x5F\xE5\x38\xD1\x0E\xB1\x89\x21\x10\x1A\x81\x57\x81\x6D\x00\x25\x05\xF5\x2E\xA1\x48\x22\x43\x77\x29\x73\xFB\xCB\xBB\x67\xA6\xB0\x4F\x08\xEC\x0A\xDC\x1A\xCA\xE1\x90\x24\x91\xCD\x87\x03\xE7\x85\x32\xDE\xF4\x1A\x02\xC0\x21\xC0\xC5\x21\x91\x08\x4D\x12\xD9\x6E\xC1\xED\x42\xF6\x60\xBF\x75\x0F\x4D\xDF\xE6\x1B\x92\x26\x48\x22\x9B\x75\x5D\xF2\x00\xDF\xC6\x9B\xBE\x5E\x23\x70\x56\x53\x6F\x51\x9B\x22\x89\x7A\x53\x6F\x1F\xB6\xEF\x75\xB7\x9A\xF3\xBE\x10\x08\xF6\x26\x6B\x98\x81\x4D\x92\x44\xED\x2B\xF8\xF6\xBA\xBE\x90\x32\x3D\xBD\x44\xE0\x4D\x60\xB9\x26\x3D\x6F\x9A\x24\xF2\x4D\x4E\x2E\xD3\xA4\x93\xD6\x56\xA7\x10\x50\x8C\x85\xF7\x9A\xF4\xA8\x0D\x92\xC8\x3F\x05\x2B\x36\x31\x04\xCA\x22\xA0\x7B\x4B\xB9\x01\xAE\xCB\x2A\xCD\x2B\xDF\x16\x49\x64\x97\xF2\x65\x6B\xD3\xD1\xC4\x10\x28\x82\x80\xAE\x8B\x3F\x51\xA4\xA0\xEF\x32\x6D\x92\x44\xBE\xDC\x51\x25\x0E\x92\x6F\x10\x4C\x5F\xD4\x08\x7C\x04\xAC\xED\x4E\x48\xB7\x62\x68\xDB\x24\x91\xD3\x4A\xA4\x32\x32\xF3\x69\x2B\xA8\x58\xA3\xB1\x20\xF0\x32\xB0\x01\xF0\x7E\x9B\x06\xC5\x40\x12\xF9\x7F\x3C\x70\x5A\x9B\x40\x58\xDB\xD1\x21\xA0\xD8\x09\x9B\xC4\x60\x55\x2C\x24\x11\x16\x76\x84\x25\x86\x11\x11\x87\x0D\x8A\xE9\xB6\x7B\x1C\xA6\x94\x0F\x98\x1D\xDA\x6E\x45\x26\x54\x84\x42\x93\xFE\x22\xA0\xB3\x7E\x47\xC6\xE4\x7E\x4C\x33\x49\x86\x8B\x4E\x0F\xEB\xF2\xCC\xCA\x31\x01\x65\xB6\x04\x47\xE0\x4B\x77\xCC\xE4\xEC\xE0\x2D\x95\x6C\x20\x46\x92\xC8\x85\x15\x1D\x51\x66\x94\xF4\xC7\x8A\xA7\x89\x80\x4E\x62\xE8\x36\xEB\x7D\x31\x9A\x1F\x2B\x49\x32\xAC\x94\x1A\x5B\xA7\x88\x4D\xBA\x8B\xC0\xE5\x8E\x20\xCA\xA9\x1E\xA5\xC4\x4E\x12\x81\xA6\x6B\xC0\x7A\xFC\xB2\x54\x0F\x51\x0E\xA1\xCA\x46\x29\xDC\x8F\x66\x8F\xE8\xEF\x1B\xA5\x40\x12\xF5\x82\x0E\xB4\x89\x28\x7B\x57\xEE\x12\xAB\x18\x13\x02\xB3\x1D\x41\x94\xE7\x26\x7A\x49\x85\x24\x19\x90\xCA\x25\x2F\xB2\x98\xA4\x8B\xC0\x45\x8E\x20\x1F\xA4\xE2\x42\x6A\x24\x11\xAE\xD3\x5D\xD8\x98\x9D\x52\x01\xD9\xEC\x1C\x43\x60\x8E\x0B\x37\x75\x5D\x6A\x78\xA4\x48\x92\x0C\xE3\x83\xDD\xFF\x48\xF6\xAA\x38\xEE\x51\xF7\x99\x23\x87\x9E\x00\x3E\x8C\xDB\xD4\xE1\xD6\xA5\x4C\x12\x79\xB4\xB4\x23\x8A\x9D\xFD\x8A\x73\xF4\xDD\xE2\x08\xF2\x78\x9C\xE6\x15\xB3\x2A\x75\x92\x64\x5E\x6E\xE1\x1E\xC1\x14\x7B\xC9\xA4\x7D\x04\x14\xFC\x5B\x33\x47\x27\x42\xDE\x76\x85\x24\xD9\xB0\x50\xA8\x7D\xBD\x56\x5C\xBE\xFD\x71\xD2\x5B\x0B\x14\xE6\x56\x04\x79\xB7\x2B\x08\x74\x8D\x24\xEA\x17\x5D\x0D\x56\x42\xFB\x03\x81\x15\xBA\xD2\x51\x91\xFB\xA1\x75\xC7\xA5\x2E\x2A\x8E\x76\xCF\x3B\x25\x5D\x24\x49\xD6\x41\x8B\x0D\x90\xA5\x6E\x5A\xE8\x4E\x75\xBA\x47\x67\xB4\x10\x57\xB8\x28\x11\xE4\x19\x8F\x7A\xA3\x52\xD5\x65\x92\x64\x40\x2F\x3C\x40\x96\x35\xA3\x42\x3F\x5D\x63\xF4\x28\x95\x91\x23\x48\xFC\xDD\x98\xA0\xE9\x03\x49\x32\xBC\x95\x90\x53\x8F\x61\xFA\xAC\x13\x53\x27\x24\x64\xCB\x5B\x6E\xD6\xD0\xCC\xA1\x20\xD5\xBD\x90\x3E\x91\x24\xEB\x50\xF9\x3C\x13\xD8\xD9\x7D\x2C\x0F\x7D\xFE\x50\x7F\xC0\xE5\x6D\xBC\xD1\x85\x84\xCA\xAF\xD1\xA1\x12\x7D\x24\xC9\x60\xF7\xE9\x48\x7E\x46\x96\x2D\x3B\xD4\xAF\x3E\x5C\xD1\x0E\xB9\x12\x9A\xEA\xD3\x4A\x94\x12\x1F\x4E\xF8\xD0\xD1\x77\x92\x0C\x62\xA8\x9C\x8F\x19\x61\xFA\xFA\x38\xA6\x4C\xC0\x19\x31\xEE\xF5\x31\xC0\xBA\xA0\xC3\x48\x32\xBC\x17\xB7\x72\xE9\xA3\x37\x06\x14\xEF\xA9\xCB\xA2\x54\xD9\x3A\x95\x3B\xCB\x11\xE4\xE3\x2E\x3B\x5B\xC5\x37\x23\x49\x3E\x6A\xCA\xD6\xB5\x99\x3B\x58\xA9\xC3\x95\xA9\xDF\x6B\x79\xCD\xDD\x00\x7C\x14\x78\x6C\x48\x6A\xEC\x7C\x44\x7A\x56\xC2\x48\x52\xBE\xC3\x45\x9A\xCD\x1D\x71\xF4\x33\xF6\x94\x77\x8A\x9B\xAB\x6B\xB1\xFA\x3C\x64\xA4\x28\xDF\xE1\x46\x92\xF2\x98\x4D\xAC\xA1\xBC\xF5\xAB\x0D\x7C\x56\x1F\xF8\xBD\xBE\xF6\xE2\x1A\x14\xC8\x4D\x7B\x16\x0A\x1F\xAB\x9F\xD9\x47\xEB\x0C\x93\x1A\x08\x18\x49\x6A\x80\x57\xA0\x6A\x46\x1E\x11\x47\x27\x00\xA6\x01\xDA\xDC\x1C\xFC\x39\xF1\x6F\x52\xAB\x0B\x49\xDA\xCD\xCE\x7E\x0E\xFE\x9E\xFD\x4D\x77\xC2\x07\x49\xF1\x45\x01\x7B\xAC\x48\x05\x04\xFE\x0F\xC3\x8D\x02\x3E\x6C\x8E\xC9\x05\x00\x00\x00\x00\x49\x45\x4E\x44\xAE\x42\x60\x82\x0C\x50\x72\x6F\x70\x6F\x72\x74\x69\x6F\x6E\x61\x6C\x09\x00\x00\x0C\x54\x4C\x61\x62\x65\x6C\x65\x64\x45\x64\x69\x74\x0C\x4C\x61\x62\x65\x6C\x65\x64\x45\x64\x69\x74\x31\x04\x4C\x65\x66\x74\x03\x90\x00\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x03\x68\x01\x05\x57\x69\x64\x74\x68\x03\x87\x01\x10\x45\x64\x69\x74\x4C\x61\x62\x65\x6C\x2E\x48\x65\x69\x67\x68\x74\x02\x11\x0F\x45\x64\x69\x74\x4C\x61\x62\x65\x6C\x2E\x57\x69\x64\x74\x68\x03\x87\x01\x11\x45\x64\x69\x74\x4C\x61\x62\x65\x6C\x2E\x43\x61\x70\x74\x69\x6F\x6E\x06\x0C\x4C\x61\x62\x65\x6C\x65\x64\x45\x64\x69\x74\x31\x15\x45\x64\x69\x74\x4C\x61\x62\x65\x6C\x2E\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x01\x04\x54\x65\x78\x74\x06\x0C\x4C\x61\x62\x65\x6C\x65\x64\x45\x64\x69\x74\x31\x00\x00\x08\x54\x4C\x69\x73\x74\x42\x6F\x78\x08\x4C\x69\x73\x74\x42\x6F\x78\x32\x04\x4C\x65\x66\x74\x02\x00\x06\x48\x65\x69\x67\x68\x74\x03\x95\x01\x03\x54\x6F\x70\x02\x00\x05\x57\x69\x64\x74\x68\x03\x90\x00\x0A\x49\x74\x65\x6D\x48\x65\x69\x67\x68\x74\x02\x00\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x02\x00\x00\x00")

// 注册Form资源  
var _ = vcl.RegisterFormResource(VersionSetting, &versionSettingBytes)
