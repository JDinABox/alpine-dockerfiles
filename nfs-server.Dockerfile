# Auto generated by github.com/jdinabox/alpine-dockerfiles

FROM golang:alpine AS builder
RUN apk --no-cache -U upgrade && \
    apk --no-cache add --upgrade make build-base
# WORKDIR /go/src/github.com/jdinabox/alpine-dockerfiles
WORKDIR /go/src/github.com/jdinabox/alpine-dockerfiles
COPY go.* ./
RUN go mod download
COPY ./ ./
# Go build cache
RUN --mount=type=cache,target=/root/.cache/go-build make -C nfs-server build

# Docker build
FROM alpine:latest

RUN apk --no-cache -U upgrade \
    && apk --no-cache add --upgrade ca-certificates \
    && wget -O /bin/dumb-init https://github.com/Yelp/dumb-init/releases/download/v1.2.5/dumb-init_1.2.5_x86_64 \
    && chmod +x /bin/dumb-init
RUN apk --no-cache add --upgrade wireguard-tools nfs-utils

# COPY --from=builder /go/src/github.com/jdinabox/alpine-dockerfiles/nfs-server/cmd/app.so /bin/app.so
COPY --from=builder /go/src/github.com/jdinabox/alpine-dockerfiles/nfs-server/cmd/app.so /bin/app.so
# WORKDIR /data/nfs/
WORKDIR /data/nfs/

EXPOSE 51820/udp

# Use dumb-init to prevent gofiber prefork from failing as PID 1
ENTRYPOINT ["/bin/dumb-init", "--"]
CMD ["/bin/app.so"]