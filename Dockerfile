FROM golang:1.21 AS build

# renovate: datasource=github-releases depName=codecov/uploader
ARG UPLOADER_VERSION=v0.6.2

WORKDIR /src
COPY . .
RUN go build -ldflags '-s -w -extldflags "-static"' -o plugin-codecov
RUN if [ $(arch) = "aarch64" ] ; then curl -Os https://github.com/codecov/uploader/releases/download/${UPLOADER_VERSION}/codecov-aarch64; fi
RUN if [ $(arch) = "x86_64" ] ; then curl -Os https://github.com/codecov/uploader/releases/download/${UPLOADER_VERSION}/codecov-alpine; fi
RUN chmod +x codecov

FROM alpine:3.18
RUN apk add -U --no-cache ca-certificates

COPY --from=build src/codecov /bin/
COPY --from=build src/plugin-codecov /bin/

ENTRYPOINT ["/bin/plugin-codecov"]
