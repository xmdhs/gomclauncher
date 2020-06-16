package launcher

import "fmt"

func (g *Gameinfo) argumentsjvm(l *launcher1155) {
	j := l.json.Patches[0].Arguments.Jvm
	for _, v := range j {
		switch v.(type) {
		case map[string]interface{}:

		case string:
			fmt.Println(v)
		}
	}
}
