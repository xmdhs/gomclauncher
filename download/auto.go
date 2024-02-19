package download

import (
	"math/rand"
	"strings"
	"sync"
)

func (r *randurls) fail(typee string) string {
	r.typeweightL.Lock()

	if i, ok := r.typeweight[typee]; ok {
		i--
		if i <= 0 {
			r.typeweight[typee] = 1
		} else {
			r.typeweight[typee] = i
		}
		r.typeweightL.Unlock()
		for {
			lenmap, t := r.auto()
			if t != typee {
				return t
			}
			if lenmap <= 1 {
				break
			}
		}
	} else {
		r.typeweightL.Unlock()
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
	r.typeweight = make(map[string]int)
	if typee == "" {
		r.typeweight["vanilla"] = 40
		r.typeweight["bmclapi"] = 50
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
	lenMap := 0
	for _, v := range r.typeweight {
		i += v
		if v > 0 {
			lenMap++
		}
	}

	a := rand.Intn(i) + 1
	for k, v := range r.typeweight {
		a = a - v
		if a <= 0 {
			return lenMap, k
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
