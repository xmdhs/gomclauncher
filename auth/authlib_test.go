package auth

import "testing"

func TestGetauthlibapi(t *testing.T) {
	type args struct {
		api string
	}
	tests := []struct {
		name           string
		args           args
		wantApiaddress string
		wantErr        bool
	}{
		//{name: "1", args: args{api: "littleskin.cn"}, wantApiaddress: "https://mcskin.littleservice.cn/api/yggdrasil", wantErr: false},
		{name: "2", args: args{api: "skin.jingqingg.com"}, wantApiaddress: " https://skin.jingqingg.com/api/yggdrasil", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotApiaddress, err := Getauthlibapi(tt.args.api)
			if (err != nil) != tt.wantErr {
				t.Errorf("Getauthlibapi() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotApiaddress != tt.wantApiaddress {
				t.Errorf("Getauthlibapi() = %v, want %v", gotApiaddress, tt.wantApiaddress)
			}
		})
	}
}
