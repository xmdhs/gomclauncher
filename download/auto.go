package download

import (
	"math/rand"
	"time"
)

var types = []string{"bmclapi", "mcbbs", "tss"}

var Fail bool

func fail(typee string) string {
	if Fail {
		s := rand.NewSource(time.Now().UnixNano())
		r := rand.New(s)
		for {
			i := r.Intn(len(types))
			t := types[i]
			if t != typee {
				return t
			}
		}
	}
	return typee
}
