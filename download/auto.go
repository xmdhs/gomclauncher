package download

import (
	"math/rand"
	"time"
)

var types = []string{"bmclapi", "mcbbs", "tss"}

func fail(typee string) string {
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
