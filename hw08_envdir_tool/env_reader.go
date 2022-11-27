package main

import (
	"bytes"
	"os"
	"path"
	"strings"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	env := make(Environment)
	for _, f := range files {
		d, err := os.ReadFile(path.Join(dir, f.Name()))
		if err != nil {
			return nil, err
		}
		fi, err := f.Info()
		if err != nil {
			return nil, err
		}
		env[strings.TrimRight(f.Name(), "=")] = EnvValue{
			clearStr(string(d)),
			fi.Size() == 0,
		}
	}
	return env, nil
}

func clearStr(s string) string {
	ms := strings.Split(s, "\n")[0]
	ms = string(bytes.ReplaceAll([]byte(ms), []byte{0x00}, []byte("\n")))
	return strings.TrimRight(ms, " \t")
}
