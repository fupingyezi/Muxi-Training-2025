FROM golang:1.24 AS builder

WORKDIR /homework

RUN go env -w GOPROXY=https://goproxy.cn,direct

COPY go.mod go.sum ./
RUN go mod download

COPY main.go .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o work .

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /homework/work .
EXPOSE 8080
CMD ["./work"]
