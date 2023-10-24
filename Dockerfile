FROM golang:1.21-alpine as build

# renovate: datasource=github-releases depName=codecov/uploader
ARG UPLOADER_VERSION=v0.6.2

WORKDIR /src
COPY . .
#
RUN apk add -q --no-cache ca-certificates curl
RUN go build -ldflags '-s -w -extldflags "-static"' -o plugin-codecov
RUN if [ $(arch) = "aarch64" ] ; then curl -s https://github.com/codecov/uploader/releases/download/${UPLOADER_VERSION}/codecov-aarch64 -o codecov; fi
RUN if [ $(arch) = "x86_64" ] ; then curl -s https://github.com/codecov/uploader/releases/download/${UPLOADER_VERSION}/codecov-alpine -o codecov; fi
RUN chmod +x codecov plugin-codecov

FROM scratch

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build src/codecov /bin/
COPY --from=build src/plugin-codecov /bin/

ENTRYPOINT ["/bin/plugin-codecov"]

