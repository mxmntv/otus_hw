package validators

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func lenValidator(val interface{}, rulevalue int64, name string) error {
	value, ok := val.(string)
	if !ok {
		return errors.Wrap(ErrUnsupportedRule, "this rule does not apply to this type")
	}
	if int64(len(value)) != rulevalue {
		return ValidationError{
			Field: name,
			Err:   fmt.Errorf("string length %s must equal to %d actual %d", name, len([]rune(value)), rulevalue),
		}
	}
	return nil
}

func regexpValidator(val interface{}, rulevalue string, name string) error {
	value, ok := val.(string)
	if !ok {
		return errors.Wrap(ErrUnsupportedRule, "this rule does not apply to this type")
	}
	rulevalue = strings.ReplaceAll(rulevalue, "\\\\", "\\")
	rgex, err := regexp.Compile(rulevalue)
	if err != nil {
		return errors.Wrap(ErrParseRegexp, "this rule does not apply to this type")
	}
	if !rgex.MatchString(value) {
		return ValidationError{
			Field: name,
			Err:   fmt.Errorf("field %s doesn't match to regexp %s", value, rulevalue),
		}
	}
	return nil
}

func inValidator(val interface{}, rulevalue string, name string) error {
	rules := strings.Split(rulevalue, ",")
	for _, r := range rules {
		rule, err := strconv.Atoi(r)
		if err != nil {
			val, ok := val.(string)
			if ok && val == r {
				return nil
			}
		} else {
			val, ok := val.(int64)
			if ok && val == int64(rule) {
				return nil
			}
		}
	}
	return ValidationError{
		Field: name,
		Err:   fmt.Errorf("field value %s must be in range %s", val.(string), rulevalue),
	}
}

func maxValidator(val interface{}, rulevalue int64, name string) error {
	value, ok := val.(int64)
	if !ok {
		return errors.Wrap(ErrUnsupportedRule, "this rule does not apply to this type")
	}
	if value > rulevalue {
		return ValidationError{
			Field: name,
			Err:   fmt.Errorf("field value %d is less than %d", value, rulevalue),
		}
	}
	return nil
}

func minValidator(val interface{}, rulevalue int64, name string) error {
	value, ok := val.(int64)
	if !ok {
		return errors.Wrap(ErrUnsupportedRule, "this rule does not apply to this type")
	}
	if value < rulevalue {
		return ValidationError{
			Field: name,
			Err:   fmt.Errorf("field value %d is less than %d", value, rulevalue),
		}
	}
	return nil
}
