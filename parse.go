package i18n

import (
	"net/http"

	"golang.org/x/text/language"
)

// A Parser is a funcion that returns a list of accepted languages, most preferred first.
type Parser func(*http.Request) []language.Tag

// ParseAcceptLanguage parses the Accept-Language header.
func ParseAcceptLanguage(r *http.Request) []language.Tag {
	t, _, _ := language.ParseAcceptLanguage(r.Header.Get("Accept-Language"))
	return t
}

// ParseFormValue parses the "locale" form value.
func ParseFormValue(r *http.Request) []language.Tag {
	t, _ := language.Parse(r.FormValue("locale"))
	return []language.Tag{t}
}

// TODO: Parsers for location, GeoIP, subdomain, TLD, request path.
