package i18n

import (
	"fmt"
	"net/http"

	"golang.org/x/text/language"
)

// LocaleFieldName defines the name used for the locale cookie or any parsable field name.
const LocaleFieldName = "locale"

// An handler contains all the necessary to parse a request and match the best language for the request.
type handler struct {
	next     http.Handler
	locales  Locales
	fallback language.Tag
	parsers  []Parser // Request parsers
	matcher  language.Matcher
}

// Handle returns a handler that will parse the language from the request thanks to the parsers, in the same order they are provided.
func Handle(h http.Handler, l Locales, fallback language.Tag, parsers ...Parser) http.Handler {
	if !l.Has(fallback) {
		panic(fmt.Errorf("i18n: fallback locale %q doesn't exist in locales map", fallback))
	}
	if len(parsers) == 0 {
		panic("i18n: at least one parser must be provided")
	}
	t := []language.Tag{fallback} // In a language.Matcher, the first element is used as the default value in case no match is found.
	for lt := range l {
		if lt != fallback {
			t = append(t, lt)
		}
	}
	return &handler{h, l, fallback, parsers, language.NewMatcher(t)}
}

// HandleFunc returns a handler wrapping an http.HandlerFunc and works the same as the Handle function.
func HandleFunc(f http.HandlerFunc, l Locales, fallback language.Tag, parsers ...Parser) http.Handler {
	return Handle(f, l, fallback, parsers...)
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var pref []language.Tag
	for _, parser := range h.parsers {
		pref = append(pref, parser(r)...)
	}
	locale, _, _ := h.matcher.Match(pref...)
	t := &Translator{locale, h.locales[locale]}
	h.next.ServeHTTP(w, RequestWithTranslator(r, t))
}

// SetCookie writes the response cookie when the parsed request locale differs from the one set inside the LocaleFieldName cookie.
// If the parsing has not been done yet (request translator is nil), this function can't help you so no cookie will be set.
func SetCookie(w http.ResponseWriter, r *http.Request) {
	rt := RequestTranslator(r)
	if rt == nil {
		panic("i18n: no locale without i18n.Handle")
	}
	if c, err := r.Cookie(LocaleFieldName); err == nil {
		if cv, err := language.Parse(c.Value); err == nil && cv == rt.Locale() {
			return // Cookie value and locale parsed from request are the same: nothing to do.
		}
	}
	http.SetCookie(w, &http.Cookie{
		Name:   LocaleFieldName,
		Value:  rt.Locale().String(),
		Path:   "/",
		MaxAge: 315360000,
	})
}
