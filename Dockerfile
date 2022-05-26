FROM golang:alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY="https://goproxy.io"

# 移动到工作目录：/build
WORKDIR /build

COPY . .

RUN go build -o app .

FROM scratch

COPY ./configs /configs

COPY --from=builder /build/app /

ENTRYPOINT ["/app"]
