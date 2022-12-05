package hw09structvalidator

import (
	"fmt"
	"reflect"

	"github.com/mxmntv/otus_hw/hw09_struct_validator/utils"
	"github.com/mxmntv/otus_hw/hw09_struct_validator/validators"
	"github.com/pkg/errors"
)

func Validate(i interface{}) error {
	v := reflect.TypeOf(i)
	s := reflect.ValueOf(i)

	if s.Kind() != reflect.Struct {
		return errors.Wrap(validators.ErrUnsupportedType,
			fmt.Sprintf("validation of this type (%s) is not possible", s.Kind()))
	}

	var valerrs validators.ValidationErrors
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		validate, ok := field.Tag.Lookup("validate")
		if ok && validate != "" {
			err := Converter(validate, s.FieldByName(field.Name), field.Name)
			if err != nil {
				var vaes validators.ValidationErrors
				if errors.As(err, &vaes) {
					valerrs = append(valerrs, vaes...)
				} else {
					return err
				}
			}
		}
	}
	if len(valerrs) > 0 {
		return valerrs
	}
	return nil
}

func Converter(rule string, v reflect.Value, name string) error {
	switch v.Kind() { //nolint:exhaustive
	case reflect.Int:
		i := make([]interface{}, 1)
		i[0] = v.Int()
		r := utils.RuleSlicer(rule)
		return validators.RuleSpreader(i, r, name)
	case reflect.String:
		s := make([]interface{}, 1)
		s[0] = v.String()
		r := utils.RuleSlicer(rule)
		return validators.RuleSpreader(s, r, name)
	case reflect.Slice:
		sl := make([]interface{}, v.Len())
		for i := 0; i < v.Len(); i++ {
			sl[i] = v.Index(i).Interface()
		}
		r := utils.RuleSlicer(rule)
		return validators.RuleSpreader(sl, r, name)
	}
	return nil
}
