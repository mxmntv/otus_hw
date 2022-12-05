package validators

import (
	"errors"
	"strconv"
)

func RuleSpreader(v []interface{}, r map[string]interface{}, n string) error { //nolint:gocognit
	var ves ValidationErrors
	for _, val := range v {
		for k, v := range r {
			switch k {
			case "len":
				v, _ := v.(string)
				i, _ := strconv.ParseInt(v, 10, 64)
				err := lenValidator(val, i, n)
				if err != nil {
					var _ve ValidationError
					if errors.As(err, &_ve) {
						ves = append(ves, _ve)
					} else {
						return err
					}
				}
			case "regexp":
				v, _ := v.(string)
				err := regexpValidator(val, v, n)
				if err != nil {
					var _ve ValidationError
					if errors.As(err, &_ve) {
						ves = append(ves, _ve)
					} else {
						return err
					}
				}
			case "in":
				v, _ := v.(string)
				err := inValidator(val, v, n)
				if err != nil {
					var _ve ValidationError
					if errors.As(err, &_ve) {
						ves = append(ves, _ve)
					} else {
						return err
					}
				}
			case "max":
				i, _ := strconv.ParseInt(v.(string), 10, 64)
				err := maxValidator(val, i, n)
				if err != nil {
					var _ve ValidationError
					if errors.As(err, &_ve) {
						ves = append(ves, _ve)
					} else {
						return err
					}
				}
			case "min":
				i, _ := strconv.ParseInt(v.(string), 10, 64)
				err := minValidator(val, i, n)
				if err != nil {
					var _ve ValidationError
					if errors.As(err, &_ve) {
						ves = append(ves, _ve)
					} else {
						return err
					}
				}
			}
		}
	}
	if len(ves) > 0 {
		return ves
	}
	return nil
}
