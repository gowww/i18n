package i18n

import (
	"testing"

	"golang.org/x/text/language"
)

func TestTransFileSize(t *testing.T) {
	cases := []struct {
		v  TransFileSize
		ll map[language.Tag]string
	}{
		{0, map[language.Tag]string{
			language.English: "0 bytes",
			language.French:  "0 octets",
		}},
		{1, map[language.Tag]string{
			language.English: "1 byte",
			language.French:  "1 octet",
		}},
		{2, map[language.Tag]string{
			language.English: "2 bytes",
			language.French:  "2 octets",
		}},
		{1000, map[language.Tag]string{
			language.English: "1 kB",
			language.French:  "1 ko",
		}},
		{1200000, map[language.Tag]string{
			language.English: "1.2 MB",
			language.French:  "1,2 Mo",
		}},
	}

	for _, c := range cases {
		for l, want := range c.ll {
			got := c.v.T(l)
			if want != got {
				t.Errorf("TransFileSize(%v).T(%v): want %v, got %v", c.v, l, want, got)
			}
		}
	}
}
