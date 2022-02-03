---
name: Codecov plugin
authors: Woodpecker Authors
icon: https://woodpecker-ci.org/img/logo.svg
description: Plugin to upload coverage reports to Codecov.io.
tags: [coverage, testing]
image: woodpeckerci/plugin-codecov
---

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
