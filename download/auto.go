package download

import (
	"math/rand"
	"strings"
	"sync"
	"time"
)

func (r *randurls) fail(typee string) string {
	if v, ok := r.typeweight.Load(typee); ok {
		i := v.(int)
		i--
		if i <= 0 {
			r.typeweight.Delete(typee)
		} else {
			r.typeweight.Store(typee, i)
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

type arand struct {
	*rand.Rand
	*sync.Mutex
}

type randurls struct {
	typeweight sync.Map
	arand
	atype string
}

func newrandurls(typee string) *randurls {
	r := &randurls{}
	if typee == "" {
		r.typeweight.Store("vanilla", 10)
		r.typeweight.Store("bmclapi", 12)
		r.typeweight.Store("mcbbs", 18)
	} else {
		s := strings.Split(typee, "|")
		for _, v := range s {
			r.typeweight.Store(v, 5)
		}
	}
	r.Rand = rand.New(rand.NewSource(time.Now().Unix()))
	r.Mutex = &sync.Mutex{}
	r.atype = typee
	return r
}

func (r *randurls) auto() (int, string) {
	if r.atype != "" && !strings.Contains(r.atype, "|") {
		return 1, r.atype
	}
	i := 0
	lenmap := 0
	typemap := make(map[string]int)
	r.typeweight.Range(
		func(k, v interface{}) bool {
			if v.(int) > 0 {
				lenmap++
				typemap[k.(string)] = v.(int)
				i += v.(int)
			}
			return true
		})
	if lenmap == 0 {
		return 0, ""
	}
	r.Lock()
	a := r.Intn(i) + 1
	r.Unlock()
	for k, v := range typemap {
		a = a - v
		if a <= 0 {
			return lenmap, k
		}
	}
	panic(a)
}

func (r *randurls) add(typee string) {
	if v, ok := r.typeweight.Load(typee); ok {
		i := v.(int)
		i++
		r.typeweight.Store(typee, i)
	}
}
