FROM       alpine:latest
MAINTAINER Brian Glogower <bglogower@docker.com>
EXPOSE     9104

ENV  GOPATH /go
ENV APPPATH $GOPATH/src/github.com/docker-infra/container-exporter
COPY . $APPPATH
RUN apk add --update -t build-deps go git mercurial libc-dev gcc libgcc \
    && cd $APPPATH && go get -d \
    && rm -rf ../../docker/docker/vendor/github.com/opencontainers/runc/ \
    && go build -o /bin/container-exporter \
    && apk del --purge build-deps && rm -rf $GOPATH
RUN chmod a+x /bin/container-exporter
ENTRYPOINT [ "/bin/container-exporter" ]
