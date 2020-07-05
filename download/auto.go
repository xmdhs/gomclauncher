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
		if typee == "" {
			typeweight.Store("vanilla", 5)
			typeweight.Store("bmclapi", 6)
			typeweight.Store("mcbbs", 9)
			typeweight.Store("tss", 12)
		} else {
			for _, v := range Types {
				typeweight.Store(v, 5)
			}
		}
		s := rand.NewSource(time.Now().Unix())
		r = rand.New(s)
	})
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
	if i == 0 {
		return typee
	}
	a := r.Intn(i) + 1
	for i, v := range b {
		a = a - v
		if a <= 0 {
			return t[i]
		}
	}
	panic(a)

}

func add(typee string) {
	if v, ok := typeweight.Load(typee); ok {
		i := v.(int)
		i++
		typeweight.Store(typee, i)
	}
}
