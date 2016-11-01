package i18n

import (
	"bytes"
	"strconv"

	"golang.org/x/text/language"
)

// Fmtn returns a formatted number with decimal and thousands marks.
func Fmtn(l language.Tag, n interface{}) (s string) {
	var b []byte

	switch n.(type) {
	case uint:
		b = []byte(strconv.FormatUint(uint64(n.(uint)), 10))
	case uint8:
		b = []byte(strconv.FormatUint(uint64(n.(uint8)), 10))
	case uint16:
		b = []byte(strconv.FormatUint(uint64(n.(uint16)), 10))
	case uint32:
		b = []byte(strconv.FormatUint(uint64(n.(uint32)), 10))
	case uint64:
		b = []byte(strconv.FormatUint(n.(uint64), 10))
	case int:
		b = []byte(strconv.Itoa(n.(int)))
	case int8:
		b = []byte(strconv.FormatInt(int64(n.(int8)), 10))
	case int16:
		b = []byte(strconv.FormatInt(int64(n.(int16)), 10))
	case int32:
		b = []byte(strconv.FormatInt(int64(n.(int32)), 10))
	case int64:
		b = []byte(strconv.FormatInt(n.(int64), 10))
	case float32:
		b = []byte(strconv.FormatFloat(float64(n.(float32)), 'f', -1, 32))
	case float64:
		b = []byte(strconv.FormatFloat(n.(float64), 'f', -1, 64))
	case string:
		b = []byte(n.(string))
	case []byte:
		b = n.([]byte)
	default:
		return
	}

	thousandsMark := ","
	decimalMark := "."
	if seps, ok := langNbrSeps[l]; ok {
		thousandsMark = seps[0]
		decimalMark = seps[1]
	}

	bb := bytes.Split(b, []byte("."))
	switch len(bb) {
	case 1:
		break
	case 2:
		s = decimalMark + string(bb[1])
		b = bb[0]
	default:
		return string(b) // Can't have 2 decimal marks in a number so return as provided.
	}

	j := 0
	for i := len(b) - 1; i >= 0; i-- {
		if j != 0 && j%3 == 0 {
			s = thousandsMark + s
		}
		s = string(b[i]) + s
		j++
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
