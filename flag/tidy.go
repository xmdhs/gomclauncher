package flag

import (
	"encoding/json"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/xmdhs/gomclauncher/launcher"
)

func (f Flag) Tidy() {
	l := []string{}
	s := Find(launcher.Minecraft + `/versions`)
	for _, v := range s {
		if Test(launcher.Minecraft + `/versions/` + v + `/` + v + ".json") {
			l = append(l, v)
		}
	}

	librariesMap := map[string]struct{}{}

	assetsMap := map[string]struct{}{}

	for _, v := range l {
		g := launcher.Gameinfo{}
		g.Version = v
		g.Gamedir = f.Minecraftpath
		g.Minecraftpath = f.Minecraftpath
		b, err := os.ReadFile(f.Minecraftpath + "/versions/" + g.Version + "/" + g.Version + ".json")
		if err != nil {
			panic(err)
		}
		g.Jsonbyte = b
		ll, _, err := g.GenLauncherCmdArgs()
		if err != nil {
			panic(err)
		}
		list := ll.CP()
		for _, v := range list {
			librariesMap[filepath.Join(v)] = struct{}{}
		}
		assetsMap[ll.GetLauncherjsonX115().AssetIndex.ID] = struct{}{}
	}

	assetPathMap := map[string]struct{}{}

	for k := range assetsMap {
		b, err := os.ReadFile(f.Minecraftpath + "/assets/indexes/" + k + ".json")
		if err != nil {
			panic(err)
		}
		a := assets{}
		err = json.Unmarshal(b, &a)
		if err != nil {
			panic(err)
		}
		for _, v := range a.Objects {
			assetPathMap[filepath.Join(f.Minecraftpath, "/assets/objects/", v.Hash[:2], "/", v.Hash)] = struct{}{}
		}
	}

	err := filepath.WalkDir(f.Minecraftpath+"/libraries", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			panic(err)
		}
		if d.IsDir() {
			return nil
		}
		_, ok := librariesMap[filepath.Join(path)]
		if !ok {
			err := os.Remove(path)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	err = filepath.WalkDir(f.Minecraftpath+"/assets/objects", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			panic(err)
		}
		if d.IsDir() {
			return nil
		}
		_, ok := assetPathMap[filepath.Join(path)]
		if !ok {
			err := os.Remove(path)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}

type assets struct {
	Objects map[string]asset `json:"objects"`
}

type asset struct {
	Hash string `json:"hash"`
}
