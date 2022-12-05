package utils

import "strings"

func RuleSlicer(r string) map[string]interface{} {
	res := make(map[string]interface{})
	s := strings.Split(r, "|")
	for _, v := range s {
		m := strings.Split(v, ":")
		res[m[0]] = m[1]
	}
	return res
}
