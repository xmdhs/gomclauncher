package flag

import (
	"strings"

	"github.com/google/uuid"
)

func UUIDgen(name string) string {
	b := []byte(name)
	if len(b) <= 16 {
		b = append(b, make([]byte, 16)...)
	}
	u, err := uuid.FromBytes(b[0:16])
	aerr(err)
	UUID := strings.ReplaceAll(u.String(), "-", "")
	return UUID
}

func aerr(err error) {
	if err != nil {
		panic(err)
	}
}
