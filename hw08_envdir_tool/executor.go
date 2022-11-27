package main

import (
	"fmt"
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	if len(cmd) < 1 {
		return -1
	}
	cmds := exec.Command(cmd[0], cmd[1:]...) //nolint:gosec
	envs := updateEnv(env)
	cmds.Env = append(os.Environ(), envs...)
	cmds.Stdin = os.Stdin
	cmds.Stdout = os.Stdout
	cmds.Stderr = os.Stderr
	if err := cmds.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return cmds.ProcessState.ExitCode()
}

func updateEnv(e Environment) []string {
	env := []string{}
	for k, v := range e {
		if _, ok := os.LookupEnv(k); ok {
			os.Unsetenv(k)
		}
		if !v.NeedRemove {
			s := fmt.Sprintf("%s=%s", k, v.Value)
			env = append(env, s)
		}
	}
	return env
}
