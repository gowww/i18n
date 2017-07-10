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
	locale       language.Tag
	Translations Translations
}

// Locale returns the locale used be the translator.
func (t *Translator) Locale() language.Tag {
	return t.locale
}

// TnPlaceholder is the placeholder replaced by n in a translation, when using the Tn function.
const TnPlaceholder = "{{.n}}"

// Tn returns the translation associated to key, for the client locale.
// If the translation defines plural forms (keys with a "Zero", "One" or "Other" suffix), it uses the most appropriate.
// All TnPlaceholder in the translation are replaced with number n.
// When translation is not found, an empty string is returned.
func (t *Translator) Tn(key string, n int, args ...interface{}) (s string) {
	for i, arg := range args {
		if ta, ok := arg.(Translatable); ok {
			args[i] = ta.T(t.Locale()) // Translate translatable arguments.
		}
	}
	if n == 0 {
		if v, ok := t.Translations[key+"Zero"]; ok {
			s = fmt.Sprintf(v, args...)
		}
	} else if n == 1 {
		if v, ok := t.Translations[key+"One"]; ok {
			s = fmt.Sprintf(v, args...)
		}
	} else if v, ok := t.Translations[key+"Other"]; ok {
		s = fmt.Sprintf(v, args...)
	} else if v, ok := t.Translations[key]; ok {
		s = fmt.Sprintf(v, args...)
	}
	s = strings.Replace(s, TnPlaceholder, FmtNumber(t.Locale(), n), -1)
	return
}

// T returns the translation associated to key, for the client locale.
func (t *Translator) T(key string, args ...interface{}) string {
	return t.Tn(key, -1, args...)
}

// TnHTML works like Tn but returns an HTML unescaped translation.
// An "nl2br" function is applied to the result.
func (t *Translator) TnHTML(key string, n int, args ...interface{}) template.HTML {
	s := t.Tn(key, n, args...)
	return template.HTML(strings.Replace(s, "\n", "<br>", -1))
}

// THTML works like T but returns an HTML unescaped translation.
// An "nl2br" function is applied to the result.
func (t *Translator) THTML(key string, args ...interface{}) template.HTML {
	return t.TnHTML(key, -1, args...)
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
