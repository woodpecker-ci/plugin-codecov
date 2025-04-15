package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func (*Plugin) command(args []string) *exec.Cmd {
	fmt.Println("$ /bin/codecov", strings.Join(args, " "))

	return exec.Command("/bin/codecov",
		args...,
	)
}
