package launcher

func (g *Gameinfo) argumentsjvm(l *launcher1155) {
	j := l.json.Patches[0].Arguments.Jvm
	for _, v := range j {
		switch v.(type) {
		case jvmRules:

		case string:

		}
	}
}

type jvmRules struct {
	Rules []Rules  `json:"rules"`
	Value []string `json:"value"`
}
type Os struct {
	Name string `json:"name"`
}
type Rules struct {
	Action string `json:"action"`
	Os     Os     `json:"os"`
}
