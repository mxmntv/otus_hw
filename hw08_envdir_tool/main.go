package main

import (
	"os"
)

func main() {
	dir := os.Args[1]
	comand := os.Args[2:]
	v, e := ReadDir(dir)
	if e != nil {
		os.Exit(1)
	}
	os.Exit(RunCmd(comand, v))
}

// RUN go build && ./hw08_envdir_tool ./testdata/env /bin/bash ./testdata/echo.sh arg1=1 arg2=2
