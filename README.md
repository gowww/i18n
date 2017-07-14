# [![gowww](https://avatars.githubusercontent.com/u/18078923?s=20)](https://github.com/gowww) i18n [![GoDoc](https://godoc.org/github.com/gowww/i18n?status.svg)](https://godoc.org/github.com/gowww/i18n) [![Build](https://travis-ci.org/gowww/i18n.svg?branch=master)](https://travis-ci.org/gowww/i18n) [![Coverage](https://coveralls.io/repos/github/gowww/i18n/badge.svg?branch=master)](https://coveralls.io/github/gowww/i18n?branch=master) [![Go Report](https://goreportcard.com/badge/github.com/gowww/i18n)](https://goreportcard.com/report/github.com/gowww/i18n) ![Status Testing](https://img.shields.io/badge/status-testing-orange.svg)

Package [i18n](https://godoc.org/github.com/gowww/i18n) provides internationalization utilities.

## Installing

1. Get package:

	```Shell
	go get -u github.com/gowww/i18n
	```

2. Import it in your code with dependencies:

	```Go
	import (
		"github.com/gowww/i18n"
		"golang.org/x/text/language"
	)
	```

## Usage

Make the [Locales](https://godoc.org/github.com/gowww/i18n#Locales) (string to string, for each language):

```Go
locales := i18n.Locales{
	language.English: {
		"hello": "Hello!",
	},
	language.French: {
		"hello": "Bonjour !",
	},
}
```

You're ready to make a handler with these locales, the default locale and the request parsers (matching the client language) you want to use.

Inside a handler, use [RequestTranslator](https://godoc.org/github.com/gowww/i18n#RequestTranslator) to get the translator containing the best locale for client.  
Use [Translator.T](https://godoc.org/github.com/gowww/i18n#Translator.T), [Translator.THTML](https://godoc.org/github.com/gowww/i18n#Translator.THTML), [Translator.Tn](https://godoc.org/github.com/gowww/i18n#Translator.Tn) or [Translator.TnHTML](https://godoc.org/github.com/gowww/i18n#Translator.TnHTML) to retreive the translation from a key.


```Go
i18n.RequestTranslator(r).T("hello")
```

So, to wrap an [http.Handler](https://golang.org/pkg/net/http/#Handler), use [Handle](https://godoc.org/github.com/gowww/i18n#Handle):

```Go
mux := http.NewServeMux()

mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	rt := i18n.RequestTranslator(r)
	fmt.Fprint(w, rt.T("hello"))
})

http.ListenAndServe(":8080", i18n.Handle(mux, locales, language.English, i18n.ParseAcceptLanguage))
```

To wrap an [http.HandlerFunc](https://golang.org/pkg/net/http/#HandlerFunc), use [HandleFunc](https://godoc.org/github.com/gowww/i18n#HandleFunc):

```Go
http.Handle("/", i18n.HandleFunc(func(w http.ResponseWriter, r *http.Request) {
	rt := i18n.RequestTranslator(r)
	fmt.Fprint(w, rt.T("hello"))
}, locales, language.English, i18n.ParseAcceptLanguage))

http.ListenAndServe(":8080", nil)
```

