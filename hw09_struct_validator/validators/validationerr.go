package validators

import (
	"fmt"
	"strings"
)

type ValidationError struct {
	Field string
	Err   error
}

func (ve ValidationError) Error() string {
	s := strings.Builder{}
	s.WriteString(fmt.Sprintf("error: %s, from field: %s", ve.Err, ve.Field))
	return s.String()
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	s := strings.Builder{}
	for i, err := range v {
		s.WriteString(fmt.Sprintf("line: %d err: %s\n", i, err.Error()))
	}
	return s.String()
}
