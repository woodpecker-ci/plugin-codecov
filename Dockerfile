FROM golang:1.21

# renovate: datasource=github-releases depName=codecov/uploader
ARG UPLOADER_VERSION=v0.6.2

WORKDIR /src
COPY . .
RUN go build -ldflags '-s -w -extldflags "-static"' -o plugin-codecov
RUN if [ $(arch) = "aarch64" ] ; then curl -s https://github.com/codecov/uploader/releases/download/${UPLOADER_VERSION}/codecov-aarch64 -o codecov; fi
RUN if [ $(arch) = "x86_64" ] ; then curl -s https://github.com/codecov/uploader/releases/download/${UPLOADER_VERSION}/codecov-alpine -o codecov; fi
RUN chmod +x codecov && mv codecov /bin/ && mv plugin-codecov /bin/

ENTRYPOINT ["/bin/plugin-codecov"]

