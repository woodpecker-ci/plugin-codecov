package main

import (
	"context"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

var version = "next"

func main() {
	app := &cli.Command{}
	app.Name = "codecov plugin"
	app.Usage = "codecov plugin"
	app.Action = run
	app.Version = version
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "token",
			Usage:   "token for authentication",
			Sources: cli.EnvVars("PLUGIN_TOKEN", "CODECOV_TOKEN"),
		},
		&cli.StringFlag{
			Name:    "name",
			Usage:   "name for coverage upload",
			Sources: cli.EnvVars("PLUGIN_NAME"),
		},
		&cli.StringSliceFlag{
			Name:    "path",
			Usage:   "paths for searching for coverage files",
			Sources: cli.EnvVars("PLUGIN_PATHS"),
		},
		&cli.StringSliceFlag{
			Name:    "file",
			Usage:   "files for coverage upload",
			Sources: cli.EnvVars("PLUGIN_FILES"),
		},
		&cli.StringSliceFlag{
			Name:    "flag",
			Usage:   "flags for coverage upload",
			Sources: cli.EnvVars("PLUGIN_FLAGS"),
		},
		&cli.StringSliceFlag{
			Name:    "env",
			Usage:   "inject environment",
			Sources: cli.EnvVars("PLUGIN_ENV"),
		},
		&cli.BoolFlag{
			Name:    "verbose",
			Usage:   "print verbose output",
			Sources: cli.EnvVars("PLUGIN_VERBOSE"),
		},
		&cli.BoolFlag{
			Name:    "dry_run",
			Usage:   "dont upload files",
			Sources: cli.EnvVars("PLUGIN_DRY_RUN"),
		},
		&cli.BoolFlag{
			Name:    "required",
			Usage:   "errors on failed upload",
			Sources: cli.EnvVars("PLUGIN_REQUIRED"),
			Value:   true,
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(_ context.Context, cmd *cli.Command) error {
	plugin := Plugin{
		Token:    cmd.String("token"),
		Name:     cmd.String("name"),
		Paths:    cmd.StringSlice("path"),
		Files:    cmd.StringSlice("file"),
		Flags:    cmd.StringSlice("flag"),
		Env:      cmd.StringSlice("env"),
		Verbose:  cmd.Bool("verbose"),
		DryRun:   cmd.Bool("dry_run"),
		Required: cmd.Bool("required"),
	}

	return plugin.Exec()
}
