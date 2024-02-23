FROM golang:1.21.1-alpine3.18 AS builder
LABEL authors="anycat"

WORKDIR /build

ENV GOPROXY=https://goproxy.cn,direct
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o anycat .

FROM alpine:3.18

WORKDIR /app
COPY --from=builder /build/anycat /app/
COPY --from=builder /build/config/ /app/

EXPOSE 0

ENTRYPOINT ["./anycat"]