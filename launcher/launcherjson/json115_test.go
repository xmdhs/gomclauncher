package launcherjson

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestLauncherjson115X115(t *testing.T) {
	b, err := ioutil.ReadFile("1.15.2.json")
	if err != nil {
		t.Error(err)
	}
	j := LauncherjsonX115{}
	json.Unmarshal(b, &j)
	if j.ID != "1.15.2" {
		t.Error(j.ID)
	}
	if j.Patches[0].Libraries[0].Downloads.Artifact.Path != "com/mojang/patchy/1.1/patchy-1.1.jar" {
		t.Error(j.Patches[0].Libraries[0].Downloads.Artifact.Path)
	}
}
