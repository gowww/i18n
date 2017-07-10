package i18n

import "golang.org/x/text/language"

// Translatable defines a translatable content that can vary from a language to another.
type Translatable interface {
	T(language.Tag) string
}

// TransInt is a translatable argument of type int.
type TransInt int

// T implements the Translatable interface.
func (v TransInt) T(l language.Tag) string {
	return FmtNumber(l, v)
}

// TransFloat64 is a translatable argument of type float64.
type TransFloat64 float64

// T implements the Translatable interface.
func (v TransFloat64) T(l language.Tag) string {
	return FmtNumber(l, v)
}

// TransFileSize is a translatable argument of representing a file size.
type TransFileSize uint64

// T implements the Translatable interface.
func (v TransFileSize) T(l language.Tag) string {
	var u string
	switch l {
	case language.French:
		u = "octets"
	default:
		u = "bytes"
	}
	// TODO: bytes, kB, MB, GB, TB, PB, EB.
	return FmtNumber(l, v) + " " + u
}
