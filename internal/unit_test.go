package internal

import "testing"

func TestNeedFixlog4j(t *testing.T) {
	type args struct {
		ver string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				ver: "1.2.3",
			},
			want: false,
		},
		{
			name: "2",
			args: args{
				ver: "2.0-beta9",
			},
			want: true,
		},
		{
			name: "3",
			args: args{
				ver: "2.15.0",
			},
			want: false,
		},
		{
			name: "4",
			args: args{
				ver: "2.15.1",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NeedFixlog4j(tt.args.ver); got != tt.want {
				t.Errorf("NeedFixlog4j() = %v, want %v", got, tt.want)
			}
		})
	}
}
