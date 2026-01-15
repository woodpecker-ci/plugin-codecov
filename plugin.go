package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Plugin struct {
	Token    string
	Name     string
	Files    []string
	Paths    []string
	Flags    []string
	Env      []string
	Verbose  bool
	DryRun   bool
	Required bool
}

func (p *Plugin) Exec() error {
	if !p.DryRun && p.Token != "" {
		return errors.New("you must provide a token")
	}

	args := p.generateArgs()
	cmd := p.command(args)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Env = append(
		os.Environ(),
		fmt.Sprintf("CODECOV_TOKEN=%s", p.Token),
	)

	return cmd.Run()
}

func (*Plugin) command(args []string) *exec.Cmd {
	fmt.Println("$ /bin/codecov", strings.Join(args, " "))

	return exec.Command("/bin/codecov",
		args...,
	)
}

func (p *Plugin) generateArgs() []string {
	args := []string{"-Q", "woodpecker"}

	if path, err := os.Getwd(); err == nil {
		args = append(args, "--rootDir", path)
	}

	if p.Name != "" {
		args = append(args, "--name", p.Name)
	}

	if len(p.Flags) != 0 {
		args = append(args, "--flags", strings.Join(p.Flags, ","))
	}

	if len(p.Env) != 0 {
		args = append(args, "--env", strings.Join(p.Env, ","))
	}

	if p.DryRun {
		args = append(args, "--dryRun")
	}

	if p.Required {
		args = append(args, "--nonZero")
	}

	for _, file := range p.Files {
		args = append(args, "--file", file)
	}

	for _, path := range p.Paths {
		args = append(args, "--dir", path)
	}

	return args
}
