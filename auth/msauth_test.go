package auth

import (
	"fmt"
	"testing"
)

func Test_getToken(t *testing.T) {
	token, err := getToken(`M.R3_BAY.443f8227-3d2-f575-8f47-a0d08ac394a`)
	if err == nil {
		t.Fatal()
		return
	}
	fmt.Println(token)
}
