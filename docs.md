---
name: Codecov
authors: Woodpecker Authors
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
