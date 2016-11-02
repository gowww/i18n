package i18n

import (
	"net/http"

	"golang.org/x/text/language"
)

// A Parser is a funcion that returns a list of accepted languages, most preferred first.
type Parser func(*http.Request) []language.Tag

// ParseAcceptLanguage parses the Accept-Language header.
func ParseAcceptLanguage(r *http.Request) []language.Tag {
	t, _, err := language.ParseAcceptLanguage(r.Header.Get("Accept-Language"))
	if err != nil {
		return nil
	}
	return t
}

// ParseCookie parses the "locale" cookie.
func ParseCookie(r *http.Request) []language.Tag {
	c, err := r.Cookie("locale")
	if err != nil {
		return nil
	}
	t, err := language.Parse(c.Value)
	if err != nil {
		return nil
	}
	return []language.Tag{t}
}

// ParseFormValue parses the "locale" form value.
func ParseFormValue(r *http.Request) []language.Tag {
	t, err := language.Parse(r.FormValue("locale"))
	if err != nil {
		return nil
	}
	return []language.Tag{t}
}

// TODO: Parsers for location, GeoIP, subdomain, TLD, request path.
