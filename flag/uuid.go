package flag

import (
	"encoding/hex"
)

func UUIDgen(t string) string {
	bb := []byte(t)
	bb = append(bb, make([]byte, 16)...)
	bb = bb[0:16]
	return hex.EncodeToString(bb)
}

func aerr(err error) {
	if err != nil {
		panic(err)
	}
}
