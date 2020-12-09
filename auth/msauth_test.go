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

func Test_jsonEscape(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{s: `tes"t`}, want: `tes\"t`},
		{name: "2", args: args{s: `test"`}, want: `test\"`},
		{name: "3", args: args{s: `"test"`}, want: `\"test\"`},
		{name: "4", args: args{s: `tes\"t`}, want: `tes\\\"t`},
		{name: "5", args: args{s: `tes\t`}, want: `tes\\t`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := jsonEscape(tt.args.s); got != tt.want {
				t.Errorf("jsonEscape() = %v, want %v", got, tt.want)
			}
		})
	}
}
