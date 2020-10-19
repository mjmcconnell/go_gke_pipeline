FROM golang:1.15.3-alpine3.12 as builder

RUN apk add --no-cache \
    gcc \
    git \
    make

WORKDIR /proxy

COPY cmd/ cmd/
COPY pkg/ pkg/
COPY go.mod go.mod
COPY go.sum go.sum
COPY Makefile Makefile

RUN make go.build

FROM alpine:3.12

COPY --from=builder /proxy/build/bin/ /usr/local/bin/

ENTRYPOINT ["proxy"]
