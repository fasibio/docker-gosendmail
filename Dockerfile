FROM golang:alpine as build_go_env
RUN  mkdir -p /src/bin &&  apk update && apk add git
ADD . /go/src/github.fasibio.de/docker-gosendmail
RUN cd /go/src/github.fasibio.de/docker-gosendmail && go get && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /src/bin/gosendmail main.go

FROM alpine:3.5
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/* 

COPY --from=build_go_env /src/bin/gosendmail /bin