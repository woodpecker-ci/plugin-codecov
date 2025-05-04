package main

import (
	"os"
	"reflect"
	"testing"
)

func TestPlugin_generateArgs(t *testing.T) {
	tests := []struct {
		name     string
		plugin   Plugin
		expected []string
	}{
		{
			name: "basic configuration",
			plugin: Plugin{
				Token: "test-token",
			},
			expected: []string{"-Q", "woodpecker", "--rootDir", mustGetwd(t)},
		},
		{
			name: "with name",
			plugin: Plugin{
				Token: "test-token",
				Name:  "test-name",
			},
			expected: []string{"-Q", "woodpecker", "--rootDir", mustGetwd(t), "--name", "test-name"},
		},
		{
			name: "with flags",
			plugin: Plugin{
				Token: "test-token",
				Flags: []string{"flag1", "flag2"},
			},
			expected: []string{"-Q", "woodpecker", "--rootDir", mustGetwd(t), "--flags", "flag1,flag2"},
		},
		{
			name: "with env",
			plugin: Plugin{
				Token: "test-token",
				Env:   []string{"env1", "env2"},
			},
			expected: []string{"-Q", "woodpecker", "--rootDir", mustGetwd(t), "--env", "env1,env2"},
		},
		{
			name: "with dry run",
			plugin: Plugin{
				Token:  "test-token",
				DryRun: true,
			},
			expected: []string{"-Q", "woodpecker", "--rootDir", mustGetwd(t), "--dryRun"},
		},
		{
			name: "with required",
			plugin: Plugin{
				Token:    "test-token",
				Required: true,
			},
			expected: []string{"-Q", "woodpecker", "--rootDir", mustGetwd(t), "--nonZero"},
		},
		{
			name: "with files",
			plugin: Plugin{
				Token: "test-token",
				Files: []string{"file1.out", "file2.json"},
			},
			expected: []string{"-Q", "woodpecker", "--rootDir", mustGetwd(t), "--file", "file1.out", "--file", "file2.json"},
		},
		{
			name: "with paths",
			plugin: Plugin{
				Token: "test-token",
				Paths: []string{"path1", "path2"},
			},
			expected: []string{"-Q", "woodpecker", "--rootDir", mustGetwd(t), "--dir", "path1", "--dir", "path2"},
		},
		{
			name: "with all options",
			plugin: Plugin{
				Token:    "test-token",
				Name:     "test-name",
				Flags:    []string{"flag1", "flag2"},
				Env:      []string{"env1", "env2"},
				DryRun:   true,
				Required: true,
				Files:    []string{"file1.out", "file2.json"},
				Paths:    []string{"path1", "path2"},
				Verbose:  true,
			},
			expected: []string{
				"-Q", "woodpecker",
				"--rootDir", mustGetwd(t),
				"--name", "test-name",
				"--flags", "flag1,flag2",
				"--env", "env1,env2",
				"--dryRun",
				"--nonZero",
				"--file", "file1.out",
				"--file", "file2.json",
				"--dir", "path1",
				"--dir", "path2",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := tt.plugin.generateArgs()
			if !reflect.DeepEqual(args, tt.expected) {
				t.Errorf("generateArgs() = %v, want %v", args, tt.expected)
			}
		})
	}
}

func TestPlugin_command(t *testing.T) {
	p := &Plugin{}
	args := []string{"--arg1", "--arg2"}
	cmd := p.command(args)

	if cmd.Path != "/bin/codecov" {
		t.Errorf("command() path = %v, want %v", cmd.Path, "/bin/codecov")
	}

	if !reflect.DeepEqual(cmd.Args, append([]string{"/bin/codecov"}, args...)) {
		t.Errorf("command() args = %v, want %v", cmd.Args, append([]string{"/bin/codecov"}, args...))
	}
}

// Helper function to get current working directory for tests
func mustGetwd(t *testing.T) string {
	path, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get working directory: %v", err)
	}
	return path
}
