---
name: Codecov
author: Woodpecker Authors
icon: https://raw.githubusercontent.com/woodpecker-ci/plugin-codecov/main/codecov.svg
description: Plugin to upload coverage reports to Codecov.io.
tags: [coverage, testing]
containerImage: woodpeckerci/plugin-codecov
containerImageUrl: https://hub.docker.com/r/woodpeckerci/plugin-codecov
url: https://github.com/woodpecker-ci/plugin-codecov
---

# Codecov

The Codecov plugin uploads coverage reports in one of the [supported formats](https://docs.codecov.com/docs/supported-report-formats) to [Codecov.io](https://about.codecov.io/).

## Usage

To use the plugin add a step similar to the following one to your Woodpecker pipeline config:

```yml
pipeline:
  codecov:
    image: woodpeckerci/plugin-codecov
    settings:
      files:
        - my-coverage-report-output.out
        - another-coverage-report.json
      token:
        from_secret: codecov_token
```

## Settings

| Settings Name | Default | Description                                                                                                          |
| ------------- | ------- | -------------------------------------------------------------------------------------------------------------------- |
| `token`       | _none_  | Codecov upload token, can also be set via the `CODECOV_TOKEN` environment variable. Required unless `dry_run` is set |
| `name`        | _none_  | Custom name for the coverage upload                                                                                  |
| `files`       | _none_  | Coverage report files to upload (list)                                                                               |
| `paths`       | _none_  | Directories to search for coverage reports (list)                                                                    |
| `flags`       | _none_  | Codecov flags to attach to the upload (list)                                                                         |
| `env`         | _none_  | Environment variable names to send along with the upload (list)                                                      |
| `dry_run`     | `false` | Do not upload the reports                                                                                            |
| `required`    | `true`  | Exit with an error when the upload fails                                                                             |
