package i18n

import (
	"bytes"
	"fmt"
	"regexp"

	"golang.org/x/text/language"
)

var reNumber = regexp.MustCompile(`^-?\d+(\.\d+)?$`)

// FmtNumber returns a formatted number with decimal and thousands marks.
func FmtNumber(l language.Tag, n interface{}) (s string) {
	var b []byte

	switch n.(type) {
	case []byte:
		b = n.([]byte)
	default:
		b = []byte(fmt.Sprint(n))
	}

	if !reNumber.Match(b) {
		return string(b)
	}

	thousandsMark := ","
	decimalMark := "."
	if seps, ok := langNbrSeps[l]; ok {
		thousandsMark = seps[0]
		decimalMark = seps[1]
	}

	var isNeg bool
	if b[0] == '-' {
		isNeg = true
		b = b[1:]
	}

	bb := bytes.Split(b, []byte("."))
	switch len(bb) {
	case 1:
		break
	case 2:
		s = decimalMark + string(bb[1])
		b = bb[0]
	}

	j := 0
	for i := len(b) - 1; i >= 0; i-- {
		if j != 0 && j%3 == 0 {
			s = thousandsMark + s
		}
		s = string(b[i]) + s
		j++
	}
	if isNeg {
		s = "-" + s
	}
	return
}

// TODO: Set decimalMark and thousandsMark for each language.
var langNbrSeps = map[language.Tag][2]string{
	language.Afrikaans:            {",", "."},
	language.Albanian:             {",", "."},
	language.AmericanEnglish:      {",", "."},
	language.Amharic:              {",", "."},
	language.Arabic:               {",", "."},
	language.Armenian:             {",", "."},
	language.Azerbaijani:          {",", "."},
	language.Bengali:              {",", "."},
	language.BrazilianPortuguese:  {",", "."},
	language.BritishEnglish:       {",", "."},
	language.Bulgarian:            {",", "."},
	language.Burmese:              {",", "."},
	language.CanadianFrench:       {",", "."},
	language.Catalan:              {",", "."},
	language.Chinese:              {",", "."},
	language.Croatian:             {",", "."},
	language.Czech:                {",", "."},
	language.Danish:               {",", "."},
	language.Dutch:                {",", "."},
	language.English:              {",", "."},
	language.Estonian:             {",", "."},
	language.EuropeanPortuguese:   {",", "."},
	language.EuropeanSpanish:      {",", "."},
	language.Filipino:             {",", "."},
	language.Finnish:              {",", "."},
	language.French:               {" ", ","},
	language.Georgian:             {",", "."},
	language.German:               {",", "."},
	language.Greek:                {",", "."},
	language.Gujarati:             {",", "."},
	language.Hebrew:               {",", "."},
	language.Hindi:                {",", "."},
	language.Hungarian:            {",", "."},
	language.Icelandic:            {",", "."},
	language.Indonesian:           {",", "."},
	language.Italian:              {",", "."},
	language.Japanese:             {",", "."},
	language.Kannada:              {",", "."},
	language.Kazakh:               {",", "."},
	language.Khmer:                {",", "."},
	language.Kirghiz:              {",", "."},
	language.Korean:               {",", "."},
	language.Lao:                  {",", "."},
	language.LatinAmericanSpanish: {",", "."},
	language.Latvian:              {",", "."},
	language.Lithuanian:           {",", "."},
	language.Macedonian:           {",", "."},
	language.Malay:                {",", "."},
	language.Malayalam:            {",", "."},
	language.Marathi:              {",", "."},
	language.ModernStandardArabic: {",", "."},
	language.Mongolian:            {",", "."},
	language.Nepali:               {",", "."},
	language.Norwegian:            {",", "."},
	language.Persian:              {",", "."},
	language.Polish:               {",", "."},
	language.Portuguese:           {",", "."},
	language.Punjabi:              {",", "."},
	language.Romanian:             {",", "."},
	language.Russian:              {",", "."},
	language.Serbian:              {",", "."},
	language.SerbianLatin:         {",", "."},
	language.SimplifiedChinese:    {",", "."},
	language.Sinhala:              {",", "."},
	language.Slovak:               {",", "."},
	language.Slovenian:            {",", "."},
	language.Spanish:              {",", "."},
	language.Swahili:              {",", "."},
	language.Swedish:              {",", "."},
	language.Tamil:                {",", "."},
	language.Telugu:               {",", "."},
	language.Thai:                 {",", "."},
	language.TraditionalChinese:   {",", "."},
	language.Turkish:              {",", "."},
	language.Ukrainian:            {",", "."},
	language.Und:                  {",", "."},
	language.Urdu:                 {",", "."},
	language.Uzbek:                {",", "."},
	language.Vietnamese:           {",", "."},
	language.Zulu:                 {",", "."},
}
