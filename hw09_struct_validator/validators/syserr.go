package validators

import "errors"

var (
	ErrUnsupportedType = errors.New("type is not supported")
	ErrUnsupportedRule = errors.New("this rule not supported")
	ErrParseRegexp     = errors.New("regexp compile failed")
)
