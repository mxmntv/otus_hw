package validators

import (
	"errors"
	"strconv"
)

func RuleSpreader(value []interface{}, rules map[string]interface{}, name string) error { //nolint:gocognit
	var validationerrs ValidationErrors
	for _, val := range value {
		for specname, spec := range rules {
			switch specname {
			case "len":
				rulevalue, _ := strconv.ParseInt(spec.(string), 10, 64)
				if err := lenValidator(val, rulevalue, name); err != nil {
					var validationerr ValidationError
					if errors.As(err, &validationerr) {
						validationerrs = append(validationerrs, validationerr)
					} else {
						return err
					}
				}
			case "regexp":
				rulevalue := spec.(string)
				if err := regexpValidator(val, rulevalue, name); err != nil {
					var validationerr ValidationError
					if errors.As(err, &validationerr) {
						validationerrs = append(validationerrs, validationerr)
					} else {
						return err
					}
				}
			case "in":
				rulevalue := spec.(string)
				if err := inValidator(val, rulevalue, name); err != nil {
					var validationerr ValidationError
					if errors.As(err, &validationerr) {
						validationerrs = append(validationerrs, validationerr)
					} else {
						return err
					}
				}
			case "max":
				rulevalue, _ := strconv.ParseInt(spec.(string), 10, 64)
				if err := maxValidator(val, rulevalue, name); err != nil {
					var validationerr ValidationError
					if errors.As(err, &validationerr) {
						validationerrs = append(validationerrs, validationerr)
					} else {
						return err
					}
				}
			case "min":
				rulevalue, _ := strconv.ParseInt(spec.(string), 10, 64)
				if err := minValidator(val, rulevalue, name); err != nil {
					var validationerr ValidationError
					if errors.As(err, &validationerr) {
						validationerrs = append(validationerrs, validationerr)
					} else {
						return err
					}
				}
			}
		}
	}
	if len(validationerrs) > 0 {
		return validationerrs
	}
	return nil
}
