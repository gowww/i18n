// Package i18n provides internationalization utilities.
package i18n

import (
	"fmt"
	"html/template"
	"strings"

	"golang.org/x/text/language"
)

// Translations is a map of translations associated to keys.
type Translations map[string]string

// Locales is a map of locales and their translations.
type Locales map[language.Tag]Translations

// Has tells if locale lt exists in the locales map.
func (l Locales) Has(lt language.Tag) bool {
	_, ok := (l)[lt]
	return ok
}

// A Translator contains the current locale with its translations and provides functions to get them.
type Translator struct {
	Locale       language.Tag
	Translations Translations
}

// TnPlaceholder is the placeholder replaced by n in a translation, when using the Tn function.
const TnPlaceholder = "{{.n}}"

// Tn returns the translation associated to key, for the client locale.
// If the translation defines plural forms (zero, one, other), it uses the most appropriate.
// All TnPlaceholder in the translation are replaced with number n.
// When translation is not found, an empty string is returned.
func (t *Translator) Tn(key string, n interface{}, a ...interface{}) (s string) {
	if n == 0 {
		if v, ok := t.Translations[key+"Zero"]; ok {
			s = fmt.Sprintf(v, a...)
		}
	} else if n == 1 {
		if v, ok := t.Translations[key+"One"]; ok {
			s = fmt.Sprintf(v, a...)
		}
	} else if v, ok := t.Translations[key+"Other"]; ok {
		s = fmt.Sprintf(v, a...)
	} else if v, ok := t.Translations[key]; ok {
		s = fmt.Sprintf(v, a...)
	}
	s = strings.Replace(s, TnPlaceholder, Fmtn(t.Locale, n), -1)
	if s == "" {
		s = fmt.Sprintf("[%s]", key)
	}
	return
}

// T returns the translation associated to key, for the client locale.
func (t *Translator) T(key string, a ...interface{}) string {
	return t.Tn(key, -1, a...)
}

// TnHTML works like Tn but returns an HTML unescaped translation.
// An "nl2br" function is applied to the result.
func (t *Translator) TnHTML(key string, n interface{}, a ...interface{}) template.HTML {
	return template.HTML(strings.Replace(t.Tn(key, n, a...), "\n", "<br>", -1))
}

// THTML works like T but returns an HTML unescaped translation.
// An "nl2br" function is applied to the result.
func (t *Translator) THTML(key string, a ...interface{}) template.HTML {
	return t.TnHTML(key, -1, a...)
}

// CleanAcceptLanguage parses, cleans and returns the contents of a Accept-Language header.
// If an error is encountered, the returned string is the same as received.
func CleanAcceptLanguage(s string) (string, error) {
	tt, q, err := language.ParseAcceptLanguage(s)
	if err != nil {
		return s, err
	}

	s = ""
	for i := 0; i < len(tt); i++ {
		if i > 0 {
			s += ","
		}
		s += fmt.Sprintf("%s;q=%g", tt[i].String(), q[i])
	}
	return s, nil
}
