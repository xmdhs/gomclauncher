package auth

import (
	"reflect"
	"testing"
)

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

func TestGetProfile(t *testing.T) {
	type args struct {
		Authorization string
	}
	tests := []struct {
		name    string
		args    args
		want    *Profile
		wantErr bool
	}{
		{name: "1", args: args{
			Authorization: "",
		},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetProfile(tt.args.Authorization)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetProfile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProfile() = %v, want %v", got, tt.want)
			}
		})
	}
}
