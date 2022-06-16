FROM golang:1.18-alpine3.16 AS builder
RUN apk add --no-cache --update git && apk add build-base
WORKDIR /app
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY ./ .
RUN go build -o .bin/main main.go
EXPOSE 1323
CMD []

FROM alpine:3.16
RUN apk --no-cache add ca-certificates
RUN apk add --no-cache tzdata
WORKDIR /app
COPY --from=builder /app/.bin/main .
EXPOSE 1323
CMD ["./main"]