#開発用 docker-compose用
FROM golang:alpine as builder

RUN apk update \
  && apk add --no-cache git \
  && go get github.com/oxequa/realize

RUN mkdir /app
WORKDIR /app
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN go build -o /main

#本番用
FROM alpine:3.9

COPY --from=builder /main .

ENTRYPOINT ["/main"]
EXPOSE 7777