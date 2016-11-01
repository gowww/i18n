package i18n

import "testing"

func TestFmtn(t *testing.T) {
	type foobarString string
	type foobar map[interface{}]interface{}

	for l, seps := range langNbrSeps {
		thousandsMark := seps[0]
		decimalMark := seps[1]

		cases := []struct {
			v    interface{}
			want string
		}{
			{123, "123"},
			{-123, "-123"},
			{1234, "1" + thousandsMark + "234"},
			{-1234, "-1" + thousandsMark + "234"},
			{uint(1234), "1" + thousandsMark + "234"},
			{uint8(123), "123"},
			{uint16(1234), "1" + thousandsMark + "234"},
			{uint32(1234), "1" + thousandsMark + "234"},
			{uint64(1234), "1" + thousandsMark + "234"},
			{int(1234), "1" + thousandsMark + "234"},
			{int8(123), "123"},
			{int16(1234), "1" + thousandsMark + "234"},
			{int32(1234), "1" + thousandsMark + "234"},
			{int64(1234), "1" + thousandsMark + "234"},
			{float32(1234.567), "1" + thousandsMark + "234" + decimalMark + "567"},
			{float64(1234.5678), "1" + thousandsMark + "234" + decimalMark + "5678"},
			{[]byte("12345678.9101"), "12" + thousandsMark + "345" + thousandsMark + "678" + decimalMark + "9101"},
			{foobar{"foo": 1234}, "map[foo:1234]"},              // Misc. type
			{foobarString("1234"), "1" + thousandsMark + "234"}, // Misc. type (string in fact)
			{"1234.5678.9", "1234.5678.9"},                      // Invalid number
		}

		for _, c := range cases {
			got := Fmtn(l, c.v)
			if got != c.want {
				t.Errorf("Fmtn(%v, %#v) == %#v, want %#v", l, c.v, got, c.want)
			}
		}

	}
}
