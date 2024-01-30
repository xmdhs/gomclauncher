package flag

import (
	"encoding/json"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/xmdhs/gomclauncher/launcher"
)

func (f *Flag) Tidy() {
	l := []string{}
	s := Find(launcher.Minecraft + `/versions`)
	for _, v := range s {
		if Test(launcher.Minecraft + `/versions/` + v + `/` + v + ".json") {
			l = append(l, v)
		}
	}

	//librariesMap := map[string]struct{}{}

	assetsMap := map[string]struct{}{}

	for _, v := range l {
		b, err := os.ReadFile(f.Minecraftpath + "/versions/" + v + "/" + v + ".json")
		if err != nil {
			panic(err)
		}
		var l launcher.LauncherjsonX115
		err = json.Unmarshal(b, &l)
		if err != nil {
			panic(err)
		}
		assetsMap[l.AssetIndex.ID] = struct{}{}
	}

	/*for _, v := range l {
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
		j := ll.GetLauncherjsonX115()
		assetsMap[j.AssetIndex.ID] = struct{}{}

		for _, v := range j.Libraries {
			path, _, url := internal.Swichnatives(v)
			if url != "" {
				librariesMap[filepath.Join(f.Minecraftpath, `/libraries/`, path)] = struct{}{}
			}
		}
	}

	librariesMap[filepath.Join(f.Minecraftpath, `/libraries/`, `moe/yushi/authlibinjector/`, "authlib-injector/", auth.Authlibversion, "/authlib-injector-"+auth.Authlibversion+".jar")] = struct{}{}
	*/
	assetPathMap := map[string]struct{}{}

	for k := range assetsMap {
		if k == "" {
			continue
		}
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

	/*err := filepath.WalkDir(f.Minecraftpath+"/libraries", func(path string, d fs.DirEntry, err error) error {
		return removeFile(err, d, librariesMap, path)
	})
	if err != nil {
		panic(err)
	}*/
	err := filepath.WalkDir(f.Minecraftpath+"/assets/objects", func(path string, d fs.DirEntry, err error) error {
		return removeFile(err, d, assetPathMap, path)
	})
	if err != nil {
		panic(err)
	}
}

func removeFile(err error, d fs.DirEntry, m map[string]struct{}, path string) error {
	if err != nil {
		panic(err)
	}
	if d.IsDir() {
		if isEmptyDir(path) {
			err = os.RemoveAll(path)
			if err != nil {
				return err
			}
			return fs.SkipDir
		}
		return nil
	}
	_, ok := m[filepath.Join(path)]
	if !ok {
		err := os.Remove(path)
		if err != nil {
			return err
		}
	}
	return nil
}

func isEmptyDir(dir string) bool {
	d, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	if len(d) == 0 {
		return true
	}
	for _, v := range d {
		if v.IsDir() {
			rt := isEmptyDir(filepath.Join(dir, v.Name()))
			if !rt {
				return false
			}
			continue
		}
		return false
	}
	return true
}

type assets struct {
	Objects map[string]asset `json:"objects"`
}

type asset struct {
	Hash string `json:"hash"`
}
