package download

import (
	"math/rand"
	"strings"
	"sync"
)

func (r *randurls) fail(typee string) string {
	r.typeweightL.Lock()
	defer r.typeweightL.Unlock()

	if i, ok := r.typeweight[typee]; ok {
		i--
		if i <= 0 {
			r.typeweight[typee] = 0
		} else {
			r.typeweight[typee] = i
		}
		for {
			lenmap, t := r.auto()
			if lenmap <= 1 {
				break
			}
			if t != typee {
				return t
			}
		}
	}
	return typee
}

type randurls struct {
	typeweight  map[string]int
	typeweightL *sync.RWMutex
	atype       string
}

func newrandurls(typee string) *randurls {
	r := &randurls{}
	if typee == "" {
		r.typeweight["vanilla"] = 10
		r.typeweight["bmclapi"] = 20
		r.typeweight["mcbbs"] = 20
	} else {
		s := strings.Split(typee, "|")
		for _, v := range s {
			r.typeweight[v] = 20
		}
	}
	r.typeweightL = &sync.RWMutex{}
	r.atype = typee
	return r
}

func (r *randurls) auto() (int, string) {
	r.typeweightL.RLock()
	defer r.typeweightL.RUnlock()

	if r.atype != "" && !strings.Contains(r.atype, "|") {
		return 1, r.atype
	}

	i := 0
	for _, v := range r.typeweight {
		i += v
	}

	a := rand.Intn(i) + 1
	for k, v := range r.typeweight {
		a = a - v
		if a <= 0 {
			return len(r.typeweight), k
		}
	}
	panic(a)
}

func (r *randurls) add(typee string) {
	r.typeweightL.Lock()
	defer r.typeweightL.Unlock()

	if i, ok := r.typeweight[typee]; ok {
		i++
		r.typeweight[typee] = i
	}
}
