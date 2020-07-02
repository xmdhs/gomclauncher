package flag

import (
	"encoding/hex"
)

func UUIDgen(t string) string {
	bb := []byte(t)
	bb = append(bb, make([]byte, 16)...)
	bb = bb[0:16]
	s := make([]byte, hex.EncodedLen(len(bb)))
	hex.Encode(s, bb)
	return string(s)
}

func aerr(err error) {
	if err != nil {
		panic(err)
	}
}
