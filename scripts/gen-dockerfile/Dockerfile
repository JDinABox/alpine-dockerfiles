FROM golang:alpine AS builder
RUN apk --no-cache -U upgrade && \
    apk --no-cache add --upgrade make build-base
# WORKDIR /go/src/{{codeModule}}
WORKDIR /go/src/{{codeModule}}
COPY go.* ./
RUN go mod download
COPY ./ ./
RUN make build

# Wireguard server/client
FROM alpine:latest

RUN apk --no-cache -U upgrade
RUN apk --no-cache add --upgrade ca-certificates
RUN apk --no-cache add --upgrade {{apk-add}}

# COPY --from=builder /go/src/{{codeModule}}/cmd/app.so /bin/app.so
COPY --from=builder /go/src/{{codeModule}}/cmd/app.so /bin/app.so
# WORKDIR /{{workDir}}/
WORKDIR /{{workDir}}/

EXPOSE 51820/udp

ENTRYPOINT ["/bin/app.so"]