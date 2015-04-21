FROM golang:onbuild

ENV CGO_ENABLED 0

RUN go build -ldflags "-linkmode external -extldflags -static" -o dumbdock

ENTRYPOINT ["/bin/sh"] 
