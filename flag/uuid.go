package flag

import (
	"crypto/md5"
	"encoding/hex"
)

func UUIDgen(t string) string {
	data := []byte(t)
	has := md5.Sum(data)
	return hex.EncodeToString(has[:])
}

func aerr(err error) {
	if err != nil {
		panic(err)
	}
}
