FROM golang:1.21.7-alpine3.18 AS builder
LABEL authors="anycat"

WORKDIR /build

ENV GOPROXY=https://goproxy.cn,direct
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o anycat .

FROM alpine:3.18

WORKDIR /app
COPY --from=builder /build/anycat /app/
COPY --from=builder /build/config.yaml /app/
RUN echo "https://mirrors.aliyun.com/alpine/v3.8/main/" > /etc/apk/repositories \
    && echo "https://mirrors.aliyun.com/alpine/v3.8/community/" >> /etc/apk/repositories \
    && apk add --no-cache tzdata \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime  \
    && echo Asia/Shanghai > /etc/timezone \
    && apk del tzdata


EXPOSE 20100

ENTRYPOINT ["./anycat"]