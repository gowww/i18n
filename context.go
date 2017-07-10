package i18n

import (
	"context"

	"net/http"
)

type contextKey int

// Context keys
const (
	ContextKeyTranslator contextKey = iota
)

// ContextTranslator returns the translator used for the context.
func ContextTranslator(c context.Context) *Translator {
	v := c.Value(ContextKeyTranslator)
	if t, ok := v.(*Translator); ok {
		return t
	}
	return nil
}

// RequestTranslator returns the translator used for the HTTP request.
func RequestTranslator(r *http.Request) *Translator {
	return ContextTranslator(r.Context())
}

// ContextWithTranslator returns the context with the translator t set.
func ContextWithTranslator(c context.Context, t *Translator) context.Context {
	return context.WithValue(c, ContextKeyTranslator, t)
}

// RequestWithTranslator returns the HTTP request with the translator t set.
func RequestWithTranslator(r *http.Request, t *Translator) *http.Request {
	return r.WithContext(ContextWithTranslator(r.Context(), t))
}
