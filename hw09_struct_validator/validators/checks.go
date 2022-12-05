package validators

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func lenValidator(v interface{}, rd int64, name string) error {
	value, ok := v.(string)
	if !ok {
		return errors.Wrap(ErrUnsupportedRule, "this rule does not apply to this type")
	}
	if int64(len(value)) != rd {
		return ValidationError{
			Field: name,
			Err:   fmt.Errorf("string length %s must equal to %d actual %d", name, len([]rune(value)), rd),
		}
	}
	return nil
}

func regexpValidator(v interface{}, rd string, name string) error {
	value, ok := v.(string)
	if !ok {
		return errors.Wrap(ErrUnsupportedRule, "this rule does not apply to this type")
	}
	rd = strings.ReplaceAll(rd, "\\\\", "\\")
	rgex, err := regexp.Compile(rd)
	if err != nil {
		return errors.Wrap(ErrParseRegexp, "this rule does not apply to this type")
	}
	if !rgex.MatchString(value) {
		return ValidationError{
			Field: name,
			Err:   fmt.Errorf("field %s doesn't match to regexp %s", value, rd),
		}
	}
	return nil
}

func inValidator(v interface{}, rd string, name string) error {
	rds := strings.Split(rd, ",")
	for _, r := range rds {
		rule, err := strconv.Atoi(r)
		if err == nil {
			val, ok := v.(int64)
			if ok {
				if val == int64(rule) {
					return nil
				}
			}
		} else {
			val, ok := v.(string)
			if ok {
				if val == r {
					return nil
				}
			}
		}
	}
	return ValidationError{
		Field: name,
		Err:   fmt.Errorf("field value %s must be in range %s", v.(string), rd),
	}
}

func maxValidator(v interface{}, rd int64, name string) error {
	value, ok := v.(int64)
	if !ok {
		return errors.Wrap(ErrUnsupportedRule, "this rule does not apply to this type")
	}
	if int64(value) > rd {
		return ValidationError{
			Field: name,
			Err:   fmt.Errorf("field value %d is less than %d", value, rd),
		}
	}
	return nil
}

func minValidator(v interface{}, rd int64, name string) error {
	value, ok := v.(int64)
	if !ok {
		return errors.Wrap(ErrUnsupportedRule, "this rule does not apply to this type")
	}
	if int64(value) < rd {
		return ValidationError{
			Field: name,
			Err:   fmt.Errorf("field value %d is less than %d", value, rd),
		}
	}
	return nil
}
