FROM golang:1.21 AS build

# renovate: datasource=github-releases depName=codecov/uploader
ARG UPLOADER_VERSION=0.1.17

WORKDIR /src
COPY . .
RUN go build -ldflags '-s -w -extldflags "-static"' -o plugin-codecov
RUN curl -Os https://uploader.codecov.io/v${CODECOV_UPLOADER_VERSION}/alpine/codecov
RUN chmod +x codecov

FROM alpine:3.18
RUN apk add -U --no-cache ca-certificates

COPY --from=build src/codecov /bin/
COPY --from=build src/plugin-codecov /bin/

ENTRYPOINT ["/bin/plugin-codecov"]
