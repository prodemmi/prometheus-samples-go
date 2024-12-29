# Stage 1: Build the Go application
FROM golang:1.23.4-alpine as build

WORKDIR /app

COPY go.mod go.sum ./

RUN go env -w GOPROXY=https://goproxy.cn

RUN go mod download && go mod verify

COPY ./ ./

RUN go build -o my-app

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/my-app .

EXPOSE 8085

ENTRYPOINT ["./my-app"]
