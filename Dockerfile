FROM golang

ADD . $GOPATH/src/github.com/l-vitaly/mailapi

RUN go get -u github.com/l-vitaly/mailapi/cmd/...

ENTRYPOINT $GOPATH/bin/mailapi

EXPOSE 9000
