package utils

import "strings"

func RuleSlicer(rules string) map[string]interface{} {
	ruleslist := make(map[string]interface{})
	parts := strings.Split(rules, "|")
	for _, p := range parts {
		r := strings.Split(p, ":")
		ruleslist[r[0]] = r[1]
	}
	return ruleslist
}
