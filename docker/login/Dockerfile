FROM golang:1.13 AS builder
COPY ../../cmd/felstorm_login.go .
WORKDIR /root/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build felstorm_login.go

FROM alpine:latest AS prod
WORKDIR /root/
COPY --from=0 /root/felstorm_login .
CMD ["sudo ./felstorm_login &"]
