FROM golang:alpine AS builder
RUN apk --no-cache -U upgrade && \
    apk --no-cache add --upgrade make build-base
# WORKDIR /go/src/{{codeModule}}
WORKDIR /go/src/{{codeModule}}
COPY go.* ./
RUN go mod download
COPY ./ ./
# Go build cache
RUN --mount=type=cache,target=/root/.cache/go-build make build

# Docker build
FROM alpine:latest

RUN apk --no-cache -U upgrade \
    && apk --no-cache add --upgrade ca-certificates \
    && wget -O /bin/dumb-init https://github.com/Yelp/dumb-init/releases/download/v1.2.5/dumb-init_1.2.5_x86_64 \
    && chmod +x /bin/dumb-init
RUN apk --no-cache add --upgrade {{apk-add}}

# COPY --from=builder /go/src/{{codeModule}}/cmd/app.so /bin/app.so
COPY --from=builder /go/src/{{codeModule}}/cmd/app.so /bin/app.so
# WORKDIR /{{workDir}}/
WORKDIR /{{workDir}}/

EXPOSE 51820/udp

# Use dumb-init to prevent gofiber prefork from failing as PID 1
ENTRYPOINT ["/bin/dumb-init", "--"]
CMD ["/bin/app.so"]