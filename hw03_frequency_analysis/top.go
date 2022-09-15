package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type Entity struct {
	Name  string
	Count int
}

func (e *Entity) counter() {
	e.Count++
}

func Finder(e []*Entity, t string) bool {
	for _, v := range e {
		if v.Name == t {
			v.counter()
			return true
		}
	}
	return false
}

func ResultSlice(e []*Entity) []string {
	max := e[0].Count
	var res []string
	for _, v := range e {
		if v.Count > max {
			max = v.Count
		}
	}
	for i := max; i >= 1; i-- {
		var s []string
		for _, v := range e {
			if v.Count == i {
				s = append(s, v.Name)
			}
		}
		sort.Slice(s, func(i, j int) (less bool) {
			return s[i] < s[j]
		})
		res = append(res, s...)
		if len(res) >= 10 {
			res = res[:10]
			break
		}
	}
	return res
}

func Top10(s string) []string {
	if len(s) == 0 {
		return nil
	}
	stringSlice := strings.Fields(s)
	var preresult []*Entity
	for _, v := range stringSlice {
		f := Finder(preresult, v)
		if !f {
			preresult = append(preresult, &Entity{
				Name:  v,
				Count: 1,
			})
		}
	}
	return ResultSlice(preresult)
}
