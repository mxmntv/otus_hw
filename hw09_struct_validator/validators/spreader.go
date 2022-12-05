package validators

import (
	"errors"
	"strconv"
)

func RuleSpreader(values []interface{}, r map[string]interface{}, n string) error {
	var ves ValidationErrors
	var ve ValidationError
	for _, val := range values {
		for k, v := range r {
			switch k {
			case "len":
				v, _ := v.(string)
				i, _ := strconv.ParseInt(v, 10, 64)
				err := lenValidator(val, i, n)
				if err != nil {
					if errors.As(err, &ve) {
						ves = append(ves, err.(ValidationError))
					} else {
						return err
					}
				}
			case "regexp":
				v, _ := v.(string)
				err := regexpValidator(val, v, n)
				if err != nil {
					if errors.As(err, &ve) {
						ves = append(ves, err.(ValidationError))
					} else {
						return err
					}
				}
			case "in":
				v, _ := v.(string)
				err := inValidator(val, v, n)
				if err != nil {
					if errors.As(err, &ve) {
						ves = append(ves, err.(ValidationError))
					} else {
						return err
					}
				}
			case "max":
				i, _ := strconv.ParseInt(v.(string), 10, 64)
				err := maxValidator(val, i, n)
				if err != nil {
					if errors.As(err, &ve) {
						ves = append(ves, err.(ValidationError))
					} else {
						return err
					}
				}
			case "min":
				i, _ := strconv.ParseInt(v.(string), 10, 64)
				err := minValidator(val, i, n)
				if err != nil {
					if errors.As(err, &ve) {
						ves = append(ves, err.(ValidationError))
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
