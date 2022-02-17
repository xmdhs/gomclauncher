package flag

import (
	"crypto/md5"
	"encoding/hex"
)

func UUIDgen(t string) string {
	data := []byte("OfflinePlayer:" + t)
	h := md5.New()
	h.Write(data)
	uuid := h.Sum(nil)
	uuid[6] = (uuid[6] & 0x0f) | uint8((3&0xf)<<4)
	uuid[8] = (uuid[8] & 0x3f) | 0x80
	return hex.EncodeToString(uuid)
}

func aerr(err error) {
	if err != nil {
		panic(err)
	}
}
