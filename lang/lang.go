package lang

import "errors"

var langmap = map[string]map[string]string{
	"zh": zh,
}

var lang map[string]string

func Setlanguge(languge string) error {
	if langmap[languge] == nil {
		return ErrNotFind
	}
	lang = langmap[languge]
	return nil
}

func Lang(key string) string {
	word, ok := lang[key]
	if !ok {
		return langmap["en"][key]
	}
	return word
}

var ErrNotFind = errors.New("ErrNotFind")
