FROM golang:1.21-alpine@sha256:926f7f7e1ab8509b4e91d5ec6d5916ebb45155b0c8920291ba9f361d65385806 as build

# renovate: datasource=github-releases depName=codecov/uploader
ARG UPLOADER_VERSION=v0.7.1

WORKDIR /src
COPY . .
#
RUN apk add -q --no-cache ca-certificates curl
RUN go build -ldflags '-s -w -extldflags "-static"' -o plugin-codecov
RUN if [ $(arch) = "aarch64" ] ; then curl -sLf https://github.com/codecov/uploader/releases/download/${UPLOADER_VERSION}/codecov-aarch64 -o codecov; fi
RUN if [ $(arch) = "x86_64" ] ; then curl -sLf https://github.com/codecov/uploader/releases/download/${UPLOADER_VERSION}/codecov-linux -o codecov; fi
RUN chmod +x codecov plugin-codecov

FROM scratch

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build src/codecov /bin/
COPY --from=build src/plugin-codecov /bin/

ENTRYPOINT ["/bin/plugin-codecov"]

