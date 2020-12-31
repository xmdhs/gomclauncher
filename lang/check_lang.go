package lang

import "github.com/cloudfoundry/jibber_jabber"

func init() {
	l, err := jibber_jabber.DetectLanguage()
	if _, ok := langmap[l]; !ok || err != nil {
		l = "en"
	}
	lang = langmap[l]
}
