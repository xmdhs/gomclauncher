package launcher

import (
	"testing"
)

func Test_needFixlog4j(t *testing.T) {
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
				ver: "2.16.0",
			},
			want: false,
		},
		{
			name: "4",
			args: args{
				ver: "2.16.1",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := needFixlog4j(tt.args.ver); got != tt.want {
				t.Errorf("needFixlog4j() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_launcher1155_cp(t *testing.T) {
	newl := func(name string) *launcher1155 {
		return &launcher1155{
			json: LauncherjsonX115{
				patchX115: patchX115{
					Libraries: []LibraryX115{
						{
							Name: name,
						},
					},
				},
			},
			Gameinfo: &Gameinfo{},
			flag:     []string{},
			fixlog4j: false,
		}
	}
	tests := []struct {
		name string
		l    *launcher1155
		want bool
	}{
		{
			name: "1",
			l:    newl("org.apache.logging.log4j:log4j-core:2.16.0"),
			want: false,
		},
		{
			name: "2",
			l:    newl("org.apache.logging.log4j:log4j-core:2.16.1"),
			want: false,
		},
		{
			name: "3",
			l:    newl("org.apache.logging.log4j:log4j-core:2.14.0"),
			want: true,
		},
		{
			name: "4",
			l:    newl("org.apache.logging.log4j:log4j-core:1.14.0"),
			want: false,
		},
		{
			name: "5",
			l:    newl("com.google.code.gson:gson:2.8.8"),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.cp(); tt.l.fixlog4j != tt.want {
				t.Errorf("launcher1155.cp() = %v, want %v", got, tt.want)
			}
		})
	}
}
