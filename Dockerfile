FROM golang

ADD . /go/src/github.com/l-vitaly/mailapi

RUN go get -u github.com/golang/dep/...
RUN cd /go/src/github.com/l-vitaly/mailapi; dep ensure;
RUN go install github.com/l-vitaly/mailapi/cmd/mailapi

ENTRYPOINT /go/bin/mailapi

EXPOSE 9000
