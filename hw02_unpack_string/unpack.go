package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func checkNum(s string) error {
	var prev rune
	for _, v := range s {
		if prev != 0 && unicode.IsDigit(prev) && unicode.IsDigit(v) || prev == 0 && unicode.IsDigit(v) {
			return ErrInvalidString
		}
		prev = v
	}
	return nil
}

func repeater(s rune, n int) string {
	if s == 10 {
		return strings.Repeat(`\n`, n)
	}
	return strings.Repeat(string(s), n-1)
}

func Unpack(s string) (string, error) {
	if err := checkNum(s); err != nil {
		return "", err
	}
	var prev rune
	var sb strings.Builder
	for _, v := range s {
		if prev != 0 && unicode.IsDigit(v) {
			i, err := strconv.Atoi(string(v))
			if err != nil {
				return "", err
			}
			if i > 0 {
				p := repeater(prev, i)
				sb.WriteString(p)
			} else {
				r := sb.String()
				r = r[0 : len(r)-1]
				sb.Reset()
				sb.WriteString(r)
			}
		}
		if !unicode.IsDigit(v) {
			prev = v
			if v != 10 {
				sb.WriteString(string(v))
			}
		}
	}
	return sb.String(), nil
}
