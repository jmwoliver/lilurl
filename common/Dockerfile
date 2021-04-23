FROM golang:1.16-alpine3.12

WORKDIR /app/
ENV GO111MODULE=auto
RUN apk --no-cache add --virtual build-deps
RUN apk update && apk add --no-cache protobuf-dev~3.12.2-r0 protoc~=3.12.2-r0 bash~=5.0.17-r0 git~=2.26.3-r0 && rm /var/cache/apk/*

RUN go get -u -v google.golang.org/grpc
RUN go get -u -v github.com/golang/protobuf/proto

ENV PATH=$PATH:$GOPATH/bin:/opt/protoc/bin