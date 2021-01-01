package lang

import (
	"github.com/Xuanwo/go-locale"
)

func init() {
	tag, err := locale.Detect()
	if err != nil {
		Setlanguge("en")
		return
	}
	l, _ := tag.Base()
	err = Setlanguge(l.String())
	if err != nil {
		Setlanguge("en")
		return
	}
}
