FROM golang:alpine AS builder
RUN apk --no-cache -U upgrade && \
    apk --no-cache add --upgrade make build-base
WORKDIR /go/src/github.com/jdinabox/alpine-dockerfiles/wireguard
COPY . .
RUN go get -d -v ./...
RUN make build


FROM alpine:latest

RUN apk --no-cache -U upgrade
RUN apk --no-cache add --upgrade ca-certificates
RUN apk --no-cache add --upgrade wireguard-tools

COPY --from=builder /go/src/github.com/jdinabox/alpine-dockerfiles/wireguard/cmd/app.so /bin/wireguard-app.so
WORKDIR /data/wireguard/

EXPOSE 51820/udp

ENTRYPOINT ["/bin/wireguard-app.so"]