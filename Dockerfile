# golang alpine 1.15.6-alpine as base image
FROM golang:1.17-alpine AS builder

ARG BUILD_VERSION

# create appuser.
RUN adduser -D -g '' elf
# create workspace
WORKDIR /opt/app/
COPY go.mod go.sum ./
# fetch dependancies
RUN go mod download && \
    go mod verify
# copy the source code as the last step
COPY . .
# build binary
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s -X main.BuildVersion=${BUILD_VERSION}" -a -installsuffix cgo -o /go/bin/wee_web_test .


# build a small image
FROM alpine:latest
LABEL language="golang"
LABEL org.opencontainers.image.source https://github.com/weedev/wee-web-test
# import the user and group files from the builder
COPY --from=builder /etc/passwd /etc/passwd
# copy the static executable
COPY --from=builder --chown=elf:1000 /go/bin/wee_web_test /wee_web_test
# use a non-root user
USER elf
# run app
ENTRYPOINT ["./wee_web_test"]