package download

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	"github.com/xmdhs/gomclauncher/launcher"
)

func Getversionlist(atype string) (*Version, error) {
	f := auto(atype)
	rep, err := Aget(source(`https://launchermeta.mojang.com/mc/game/version_manifest.json`, f))
	if rep != nil {
		defer rep.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadAll(rep.Body)
	if err != nil {
		return nil, err
	}
	v := Version{}
	json.Unmarshal(b, &v)
	v.atype = atype
	return &v, nil
}

type Version struct {
	Latest   VersionLatest    `json:"latest"`
	Versions []VersionVersion `json:"versions"`
	atype    string
}

type VersionLatest struct {
	Release  string `json:"release"`
	Snapshot string `json:"snapshot"`
}

type VersionVersion struct {
	ID          string `json:"id"`
	ReleaseTime string `json:"releaseTime"`
	Time        string `json:"time"`
	Type        string `json:"type"`
	URL         string `json:"url"`
}

func (v Version) Downjson(ver string) error {
	f := auto(v.atype)
	for _, vv := range v.Versions {
		if vv.ID == ver {
			err := get(source(vv.URL, f), launcher.Minecraft+`/versions/`+vv.ID+`/`+vv.ID+`.json`)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return errors.New("no such")
}
