package lang

import (
	"errors"

	"golang.org/x/text/language"
)

var langmap = map[string]map[string]string{
	"zh": zh,
	"en": en,
}

var lang map[string]string

func Setlanguge(languge string) error {
	tag := language.Make(languge)
	l, _ := tag.Base()
	if langmap[l.String()] == nil {
		return ErrNotFind
	}
	lang = langmap[l.String()]
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
