package download

import (
	"math/rand"
	"time"
)

var (
	Types []string
	Fail  bool
)

func fail(typee string) string {
	if Fail {
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
