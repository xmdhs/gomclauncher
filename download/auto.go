package download

import (
	"math/rand"
	"sync"
	"time"
)

var (
	typeweight sync.Map
	one        sync.Once
	r          *rand.Rand
)

func auto(typee string) string {
	one.Do(func() {
		typeweight.Store("vanilla", 2)
		typeweight.Store("bmclapi", 2)
		typeweight.Store("mcbbs", 4)
		typeweight.Store("tss", 4)
		s := rand.NewSource(time.Now().Unix())
		r = rand.New(s)
	})
	if typee == "" {
		i := 0
		t := make([]string, 0, 4)
		b := make([]int, 0, 4)
		typeweight.Range(
			func(k, v interface{}) bool {
				t = append(t, k.(string))
				b = append(b, v.(int))
				i += v.(int)
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

func fail(typee string) string {
	if v, ok := typeweight.Load(typee); ok {
		i := v.(int)
		i--
		if i <= 0 {
			typeweight.Store(typee, 0)
		}
		for {
			t := auto("")
			if t != typee {
				return t
			}
		}
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
