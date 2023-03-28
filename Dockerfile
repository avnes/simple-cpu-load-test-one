FROM golang:latest AS build

ADD src/main.go src/

RUN cd src && \
    go mod init loadtest && \
    go mod tidy &&  \
    go build main.go

FROM alpine:latest

WORKDIR /

COPY --from=build /go/src/main /main

USER nobody:nobody

ENTRYPOINT ["/main"]
