# plugin-codecov

[![Build status](https://ci.woodpecker-ci.org/api/badges/woodpecker-ci/plugin-codecov/status.svg)](https://ci.woodpecker-ci.org/woodpecker-ci/plugin-codecov)
[![Docker Image Version (latest by date)](https://img.shields.io/docker/v/woodpeckerci/plugin-codecov?label=DockerHub%20latest%20version&sort=semver)](https://hub.docker.com/r/woodpeckerci/plugin-codecov/tags)

Woodpecker plugin to send coverage reports to [Codecov](https://codecov.io/).

## Build

Build the Docker image with the following command:

```sh
docker build -f Dockerfile -t woodpeckerci/plugin-codecov:next .
```

## Test

```bash
go test -race -coverprofile=coverage.out ./...

docker run --rm -it \
  -e CI=woodpecker \
  -e PLUGIN_TOKEN="dummy" \
  -e PLUGIN_DRY_RUN="true" \
  -v $(pwd):/woodpecker/src \
  -w /woodpecker/src \
  woodpeckerci/plugin-codecov:next
```
