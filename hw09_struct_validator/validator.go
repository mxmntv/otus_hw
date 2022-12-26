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

	var validationerrs validators.ValidationErrors
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		validate, ok := field.Tag.Lookup("validate")
		if ok && validate != "" {
			err := Converter(validate, s.FieldByName(field.Name), field.Name)
			if err != nil {
				var validationerr validators.ValidationErrors
				if errors.As(err, &validationerr) {
					validationerrs = append(validationerrs, validationerr...)
				} else {
					return err
				}
			}
		}
	}
	if len(validationerrs) > 0 {
		return validationerrs
	}
	return nil
}

func Converter(rules string, value reflect.Value, name string) error {
	switch value.Kind() { //nolint:exhaustive
	case reflect.Int:
		val := make([]interface{}, 1)
		val[0] = value.Int()
		rule := utils.RuleSlicer(rules)
		return validators.RuleSpreader(val, rule, name)
	case reflect.String:
		val := make([]interface{}, 1)
		val[0] = value.String()
		rule := utils.RuleSlicer(rules)
		return validators.RuleSpreader(val, rule, name)
	case reflect.Slice:
		val := make([]interface{}, value.Len())
		for i := 0; i < value.Len(); i++ {
			val[i] = value.Index(i).Interface()
		}
		rule := utils.RuleSlicer(rules)
		return validators.RuleSpreader(val, rule, name)
	}
	return nil
}
