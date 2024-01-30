package internal

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSafePathJoin(t *testing.T) {
	p, err := SafePathJoin(".minecraft", "aaa")
	assert.Nil(t, err)
	assert.Equal(t, p, ".minecraft"+string(filepath.Separator)+"aaa")

	_, err = SafePathJoin(".minecraft", "../aaa")
	assert.ErrorIs(t, err, ErrPathInvalid)

	_, err = SafePathJoin(".minecraft", "aaa/../../a")
	assert.ErrorIs(t, err, ErrPathInvalid)

	_, err = SafePathJoin("Z:/game/.minecraft/assets/virtual/legacy/", "records/11.ogg")
	assert.Nil(t, err)

}
