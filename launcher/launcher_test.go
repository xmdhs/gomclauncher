package launcher

import (
	"encoding/json"
	"gomclauncher/launcher/launcherjson"
	"io/ioutil"
	"testing"
)

func TestLauncher115(t *testing.T) {
	b, err := ioutil.ReadFile("./launcherjson/1.15.2.json")
	if err != nil {
		t.Error(err)
	}
	j := launcherjson.LauncherjsonX115{}
	json.Unmarshal(b, &j)
	//Launcher115(j)
}
