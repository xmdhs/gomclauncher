package download

import (
	"encoding/json"
	"io/ioutil"
)

func Getversionlist(atype string) (*Version, error) {
	rep, err := aget(source(`https://launchermeta.mojang.com/mc/game/version_manifest.json`, atype))
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadAll(rep.Body)
	defer rep.Body.Close()
	if err != nil {
		return nil, err
	}
	v := Version{}
	json.Unmarshal(b, &v)
	return &v, nil
}

type Version struct {
	Latest   VersionLatest    `json:"latest"`
	Versions []VersionVersion `json:"versions"`
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
	for _, v := range v.Versions {
		if v.ID == ver {
			err := get(v.URL, `.minecraft/versions/`+v.ID+`/`+v.ID+`.json`)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
