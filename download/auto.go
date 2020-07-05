package download

import (
	"math/rand"
	"sync"
	"time"
)

var (
	Types []string
	Fail  bool
)

func fail(typee string) string {
	if v, ok := typeweight.Load(typee); ok {
		i := v.(int)
		i--
		if i <= 0 {
			typeweight.Store(typee, 0)
		} else {
			typeweight.Store(typee, i)
		}
		for {
			t := auto("")
			if t != typee {
				return t
			}
		}
	} else if Fail {
		s := rand.NewSource(time.Now().UnixNano())
		r := rand.New(s)
		for {
			i := r.Intn(len(Types))
			t := Types[i]
			if t != typee {
				return t
			}
		}
	}
	return typee
}

var (
	typeweight sync.Map
	one        sync.Once
	r          *rand.Rand
)

func auto(typee string) string {
	one.Do(func() {
		typeweight.Store("vanilla", 5)
		typeweight.Store("bmclapi", 8)
		typeweight.Store("mcbbs", 5)
		typeweight.Store("tss", 8)
		s := rand.NewSource(time.Now().Unix())
		r = rand.New(s)
	})
	if typee == "" {
		i := 0
		t := make([]string, 0, 4)
		b := make([]int, 0, 4)
		typeweight.Range(
			func(k, v interface{}) bool {
				if v.(int) != 0 {
					t = append(t, k.(string))
					b = append(b, v.(int))
					i += v.(int)
				}
				return true
			})
		a := r.Intn(i) + 1
		for i, v := range b {
			a = a - v
			if a <= 0 {
				return t[i]
			}
		}
		panic(a)
	}
	return typee
}

func add(typee string) {
	if v, ok := typeweight.Load(typee); ok {
		i := v.(int)
		i++
		typeweight.Store(typee, i)
	}

}
