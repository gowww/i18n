package i18n

import (
	"fmt"
	"net/http"

	"golang.org/x/text/language"
)

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
		panic("i18n: at least one matcher must be provided")
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
	var t *Translator
	for _, parser := range h.parsers {
		locales := parser(r)
		if locale, _, conf := h.matcher.Match(locales...); conf != language.No {
			t = &Translator{locale, h.locales[locale]}
			break
		}
	}
	if t == nil {
		t = &Translator{h.fallback, h.locales[h.fallback]}
	}
	h.next.ServeHTTP(w, RequestWithTranslator(r, t))
}
