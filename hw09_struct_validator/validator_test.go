package hw09structvalidator

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

type UserRole string

// Test the function on different structures and other types.
type (
	User struct {
		ID     string `json:"id" validate:"len:10"`
		Name   string
		Age    int      `validate:"min:18|max:50"`
		Email  string   `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		Role   UserRole `validate:"in:admin,stuff"`
		Phones []string `validate:"len:11"`
		meta   json.RawMessage
	}

	App struct {
		Version string `validate:"len:5"`
	}

	Token struct {
		Header    []byte
		Payload   []byte
		Signature []byte
	}

	Response struct {
		Code int    `validate:"in:200,404,500"`
		Body string `json:"omitempty"`
	}
)

func TestValidate(t *testing.T) {
	tests := []struct {
		in          interface{}
		expectedErr error
	}{
		{
			in: Response{
				Code: 200,
				Body: "{}",
			},
			expectedErr: nil,
		},
		{
			in: App{
				Version: "0.0.1",
			},
			expectedErr: nil,
		},
		{
			in: Token{
				Header:    []byte{1, 2, 3},
				Payload:   []byte{3, 2, 1},
				Signature: []byte{2, 2, 2},
			},
			expectedErr: nil,
		},
		{
			in: User{
				ID:    "0123456789",
				Name:  "Vasya",
				Age:   22,
				Email: "example@example.com",
				Role:  "admin",
				Phones: []string{
					"11111111111",
					"22222222222",
				},
				meta: json.RawMessage("{}"),
			},
			expectedErr: nil,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			tt := tt
			t.Parallel()

			e := Validate(tt.in)

			require.True(t, errors.Is(e, tt.expectedErr))
			_ = tt
		})
	}
}
