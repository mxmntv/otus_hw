package hw02unpackstring

import (
	"errors"
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

func Unpack(s string) (string, error) {
	if err := checkNum(s); err != nil {
		return "", err
	}
	var sb strings.Builder
	par := make([]rune, 0)
	for _, v := range s {
		if unicode.IsDigit(v) {
			if v == '0' {
				sb.Reset()
				par = par[:len(par)-1]
				sb.WriteString(string(par))
			} else {
				repeat := strings.Repeat(string(par[len(par)-1]), int(v)-'0'-1)
				par = append(par, []rune(repeat)...)
				sb.WriteString(repeat)
			}
		} else {
			par = append(par, v)
			sb.WriteString(string(v))
		}
	}
	return sb.String(), nil
}
