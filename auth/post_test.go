package auth

import (
	"testing"
)

func TestPost(t *testing.T) {
	j := `
	{
		"agent": {                            
			"name": "Minecraft",               
			"version": 1                       
												
		},
		"username": "mojang帐号名",             
												
		"password": "mojang帐号密码",
		"clientToken": "客户端标识符",          
		"requestUser": true                   
	}`
	b, err, i := post("", []byte(j))
	if err != nil {
		t.Fatal(err)
	}
	a := string(b)
	t.Log(i)
	t.Log(a)
}
