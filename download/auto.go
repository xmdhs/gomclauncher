package download

import (
	"math/rand"
	"strings"
	"sync"
	"time"
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
	r          arand
)

type arand struct {
	*rand.Rand
	*sync.Mutex
}

func auto(typee string) string {
	if typee != "" && !strings.Contains(typee, "|") {
		return typee
	}
	one.Do(func() {
		if typee == "" {
			typeweight.Store("vanilla", 5)
			typeweight.Store("bmclapi", 6)
			typeweight.Store("mcbbs", 9)
		} else {
			s := strings.Split(typee, "|")
			for _, v := range s {
				typeweight.Store(v, 5)
			}
		}
		r.Rand = rand.New(rand.NewSource(time.Now().Unix()))
		r.Mutex = &sync.Mutex{}
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
	r.Lock()
	a := r.Intn(i) + 1
	r.Unlock()
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
