package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	t.Run("unset env", func(t *testing.T) {
		os.Setenv("CAT", "dummy")
		env := make(Environment)
		env["CAT"] = EnvValue{
			"clever",
			false,
		}
		RunCmd([]string{"ls"}, env)
		cat, ok := os.LookupEnv("CAT")
		if !ok {
			fmt.Println("env not found")
		}
		require.Equal(t, "", cat)
	})

	t.Run("empty cmd & env", func(t *testing.T) {
		s := []string{}
		r := RunCmd(s, Environment{})
		require.Equal(t, -1, r)
	})

	t.Run("empty env", func(t *testing.T) {
		r := RunCmd([]string{"ls"}, Environment{})
		require.Equal(t, 0, r)
	})
}
