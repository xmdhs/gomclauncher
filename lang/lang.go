package lang

import (
	"errors"
	"sync"

	"golang.org/x/text/language"
)

var langmap = map[string]map[string]string{
	"zh": zh,
	"en": en,
}

var lang map[string]string

var lock = sync.RWMutex{}

func Setlanguge(languge string) error {
	lock.Lock()
	defer lock.Unlock()
	tag := language.Make(languge)
	l, _ := tag.Base()
	if langmap[l.String()] == nil {
		return ErrNotFind
	}
	lang = langmap[l.String()]
	return nil
}

func Lang(key string) string {
	lock.RLock()
	defer lock.RUnlock()
	word, ok := lang[key]
	if !ok {
		return langmap["en"][key]
	}
	return word
}

var ErrNotFind = errors.New("ErrNotFind")
