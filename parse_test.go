package i18n

import (
	"net/http"
	"net/url"
	"testing"

	"golang.org/x/text/language"
)

func TestParseAcceptLanguage(t *testing.T) {
	cases := []struct {
		v    string
		want []language.Tag
	}{
		{"", []language.Tag{}},
		{"en", []language.Tag{language.English}},
		// Sorting
		{"en;q=0.1,de;q=0.2,fr;q=0.2", []language.Tag{language.German, language.French, language.English}},
		// Dropping
		{"en;q=0.1,de;q=0,fr;q=0.2,i-klingon;q=0.0", []language.Tag{language.French, language.English}},
	}
	for _, c := range cases {
		r := &http.Request{Header: http.Header{"Accept-Language": []string{c.v}}}
		got := ParseAcceptLanguage(r)
		if len(got) != len(c.want) {
			t.Errorf("ParseFormValue(%#v) == %v, want %v", c.v, got, c.want)
			continue
		}
		for i := range got {
			if got[i] != c.want[i] {
				t.Errorf("ParseAcceptLanguage(%#v) == %v, want %v", c.v, got, c.want)
				break
			}
		}
	}
}

func TestParseFormValue(t *testing.T) {
	cases := []struct {
		v    string
		want []language.Tag
	}{
		{"", []language.Tag{}},
		{"en", []language.Tag{language.English}},
	}
	for _, c := range cases {
		r := &http.Request{Form: url.Values{"locale": []string{c.v}}}
		got := ParseFormValue(r)
		if len(got) != len(c.want) {
			t.Errorf("ParseFormValue(%#v) == %v, want %v", c.v, got, c.want)
			continue
		}
		for i := range got {
			if got[i] != c.want[i] {
				t.Errorf("ParseFormValue(%#v) == %v, want %v", c.v, got, c.want)
				break
			}
		}
	}
}
